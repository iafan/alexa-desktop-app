package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/net/websocket"
)

// AlexaState defines Alexa state
type AlexaState struct {
	State      string `json:"state,omitempty"`
	AudioState string `json:"audiostate,omitempty"`
	Mute       bool   `json:"mute,omitempty"`
	Card       string `json:"card,omitempty"`
	Title      string `json:"title,omitempty"`
	RawPayload string `json:"raw_payload,omitempty"`
}

// AlexaStateChanger is a function that updates Alexa state
type AlexaStateChanger func(a *AlexaState)

type Server struct {
	clients map[*Client]bool
	ch      chan []byte
}

func (s *Server) AddClient(ws *websocket.Conn) *Client {
	c := &Client{
		ch:     make(chan []byte, 10),
		ws:     ws,
		server: s,
	}
	s.clients[c] = true
	return c
}

func (s *Server) RemoveClient(c *Client) {
	log.Println("[Server.RemoveClient]")
	delete(s.clients, c)
	c.Close()
}

func (s *Server) Serve() {
	// wait for new state change events and send them to the client
	log.Println("[Server.Serve] waiting for events...")
	for bytes := range events {
		n := len(s.clients)
		if n == 0 {
			log.Printf("[Server.Serve] no clients connected")
		} else {
			log.Printf("[Server.Serve] sending event to %d client(s)", n)

			for c := range s.clients {
				c.SendData(bytes)
			}
		}
	}
	log.Println("[Server.Serve] finished")
}

func NewServer(ch chan []byte) *Server {
	return &Server{
		clients: make(map[*Client]bool),
		ch:      ch,
	}
}

type Client struct {
	ch     chan []byte
	ws     *websocket.Conn
	server *Server
}

func (c *Client) Serve() {
	log.Printf("[Client.Serve] started")
	defer c.server.RemoveClient(c)
	for data := range c.ch {
		_, err := c.ws.Write(data)
		if err != nil {
			return
		}
	}
	log.Printf("[Client.Serve] finished")
}

func (c *Client) SendData(data []byte) {
	c.ch <- data
}

func (c *Client) Close() {
	c.ws.Close()
}

var state = &AlexaState{}
var events = make(chan []byte, 10)
var server = NewServer(events)

var triggerLines = map[string]AlexaStateChanger{
	"#       Connecting...       #": func(a *AlexaState) {
		a.State = "connecting"
	},

	"#       Authorized!       #": func(a *AlexaState) {
		a.State = "authorized"
	},
	"#       Alexa is currently idle!       #": func(a *AlexaState) {
		a.State = "idle"
	},
	"#       Listening...       #": func(a *AlexaState) {
		a.State = "listening"
	},
	"#       Thinking...       #": func(a *AlexaState) {
		a.State = "thinking"
	},
	"#       Speaking...       #": func(a *AlexaState) {
		a.State = "speaking"
	},
	"# Audio state         : PLAYING": func(a *AlexaState) {
		a.AudioState = "playing"
	},
	"# Audio state         : PAUSED": func(a *AlexaState) {
		a.AudioState = "paused"
	},
	"# Audio state         : STOPPED": func(a *AlexaState) {
		a.AudioState = ""
	},
	"#     RenderTemplateCard": func(a *AlexaState) {
		a.Card = "template"
		a.Title = ""
		a.RawPayload = ""
	},
	"#     RenderTemplateCard - Cleared": func(a *AlexaState) {
		a.Card = ""
		a.Title = ""
		a.RawPayload = ""
	},
	"#     RenderPlayerInfoCard": func(a *AlexaState) {
		a.Card = "player"
		a.Title = ""
		a.RawPayload = ""
	},
	"#     RenderPlayerInfoCard - Cleared": func(a *AlexaState) {
		a.Card = ""
		a.Title = ""
		a.RawPayload = ""
	},
}

var reMainTitle = regexp.MustCompile("^# Main Title          : (.*)$")
var reContent = regexp.MustCompile("^\\{\"content\":")
var reMute = regexp.MustCompile("^#       SOURCE:DIRECTIVE .+ MUTE:([01])")

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func processLine(s string) (stateChanged bool) {
	// try exact line matches
	f := triggerLines[s]
	if f != nil {
		f(state)
		return true
	}

	// try regex matches

	// title
	matches := reMainTitle.FindStringSubmatch(s)
	if len(matches) > 0 {
		state.Title = matches[1]
		return true
	}

	// content
	matches = reContent.FindStringSubmatch(s)
	if len(matches) > 0 {
		state.RawPayload = s
		return true
	}

	// mute
	matches = reMute.FindStringSubmatch(s)
	if len(matches) > 0 {
		state.Mute = matches[1] == "1"
		return true
	}

	return false
}

func getStateBytes() []byte {
	bytes, err := json.Marshal(state)
	if err != nil {
		log.Printf("json.Marshal Error: %s", err)
		bytes = []byte("{error: \"internal\"}")
	}
	return bytes
}

func monitorStdout(stdout io.ReadCloser) {
	reader := bufio.NewReader(stdout)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error: %s", err)
			break
		}

		log.Print(s)

		isChanged := processLine(strings.TrimSpace(s))
		if !isChanged {
			continue
		}

		bytes := getStateBytes()
		fmt.Println(string(bytes))
		events <- bytes // write event to a channel
	}
}

func wsServer(ws *websocket.Conn) {
	// send current state
	log.Println("[wsServer] new client connected")
	c := server.AddClient(ws)
	c.SendData(getStateBytes())
	c.Serve()
}

func startAlexa() {
	// Run the proxy runner app and attach to its STDIN/STDOUT

	cmd := exec.Command("./bin/sample-app-runner")

	stdout, err := cmd.StdoutPipe()
	check(err)
	defer stdout.Close()

	stdin, err := cmd.StdinPipe()
	check(err)
	defer stdin.Close()

	err = cmd.Start()

	// Start monitoring and parsing STDOUT

	go monitorStdout(stdout)

	// Start the web server

	go server.Serve()

	// run for as long as proxy command runs

	check(cmd.Wait())
}

func main() {
	go startAlexa()

	addr := ":8080"

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	http.Handle("/ws", websocket.Handler(wsServer))
	fmt.Printf("Starting server on %s\n", addr)
	check(http.ListenAndServe(addr, nil))
}

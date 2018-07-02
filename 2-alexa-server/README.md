This is a web server that wraps the functionality of the sample app
from AVS Device SDK and provides a web-based UI for Alexa.

# Prerequisites

Before you run this app, make sure your sample app from AVS Device SDK is configured properly. Check out `bin` subfolder and [bin/README.md](bin/README.md).

# Build

    go get -u golang.org/x/net/websocket

    go build alexa-server.go

# Run

    ./alexa-server

Note that this server will run the AVS sample app under the hood, so you shouldn't run the sample app manually.

To test the web UI, open http://localhost:8080/. Say "Alexa" to start the conversation.
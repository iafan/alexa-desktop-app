const electron = require('electron');
const { app, BrowserWindow } = electron;
const path = require('path');
const url = require('url');

let win;

let winSize = 280;

function createWindow() {
    win = new BrowserWindow({ width: winSize, height: winSize, transparent: true, frame: false });

    win.setSize(winSize, winSize); // actually set to required dimensions, now that frame is gone

    // from https://discuss.atom.io/t/set-browserwindow-always-on-top-even-other-app-is-in-fullscreen/34215/3
    // hides the dock icon for our app which allows our windows to join other
    // apps' spaces. without this our windows open on the nearest "desktop" space
    app.dock.hide();
    // "floating" + 1 is higher than all regular windows, but still behind things
    // like spotlight or the screen saver
    win.setAlwaysOnTop(true, "floating", 1);
    // allows the window to show over a fullscreen window
    win.setVisibleOnAllWorkspaces(true);

    // load the server app
    win.loadURL('http://localhost:8080');

    //win.webContents.openDevTools(); // DEBUG

    win.on('closed', () => {
        console.log('window: closed');
        win = null;
    });

    win.on('move', () => {
        console.log('window: move');
        // when dragged onto 640x480 display, snap to screen
        let displayBounds = electron.screen.getDisplayMatching(win.getBounds()).bounds;
        const { width, height } = displayBounds;
        console.log('screen size:', width, height);
        if (width === 640 && height === 480) {
            win.setBounds(displayBounds);
            win.setFullScreen(true);
        }
    });
}

app.on('ready', createWindow);

app.on('window-all-closed', () => {
    console.log('window: window-all-closed');
    app.quit();
});

app.on('activate', () => {
    console.log('window: activate');
});

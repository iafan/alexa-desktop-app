This Electron app is a wrapper for the server (see `../2-alexa-server`) that runs Alexa UI as a chromeless round always-on-top window on your desktop.

This application also detects if it's being dragged onto a 640x480 display, in which case it forces the full-screen mode and a slighly larger size. This functionality was added so that Alexa could run on a device like Gakken WorldEye (a spherical projector with 640x480 logical and 480x480 effective resolution).

# Prerequisites

You need to install Electron either via npm (`npm install electron --save-dev`) or by downloading the Electron app and installing it into your Applications folder.

Make sure `alexa-server` is running (see `../2-alexa-server`).

# Run

If have Electron installed via npm, run:

    electron .

If have Electron in your Applications folder, use this helper script instead:

    ./run

# Usage

You can drag and drop the Alexa window on your desktop. There's no way to minimize the app (though technically it should be possible to hide the app by default and make it appear when a wake word is detected; I just didn't need this for my practical purposes).

To close the app, make sure it is focused (by e.g. clicking or dragging it) and then press `Cmd+W`.
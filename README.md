# Amazon Alexa on your desktop

This was started more as an art project rather than an attempt to create a robust production-ready application. I wanted to run Alexa on my Gakken WorldEye projector (a spherical display), because I like all retro-futuristic things.

First of all, check out the final result:

[![](http://img.youtube.com/vi/JrK_ADgYo8g/0.jpg)](https://www.youtube.com/watch?v=JrK_ADgYo8g)

And here's the same app running on a Mac desktop:

[![](http://img.youtube.com/vi/rEFqu6UMKb4/0.jpg)](https://www.youtube.com/watch?v=rEFqu6UMKb4)

# Approach

To write less code, I decided to do the following:

1. Use the console sample app from AVS Device SDK with no modifications. I just needed to compile it with wake word detection from KITT.AI / Snowboy as per official instructions.
2. Write a wrapper server app that would in turn run that sample app, parse its console output and send events via websocket to an HTML page that would render the UI.
3. To also make the app work on a desktop, write a helper Electron app to initialize the chromeless window with a custom round shape and always-on-top behavior.

# Setup

This repository contains a bunch of components, each residing in its own directory. Check out READMEs in each directory for further setup instructions:

1. [1-sample-app](1-sample-app/README.md) — build the AVS Device SDK and sample app
2. [2-alexa-server/bin](2-alexa-server/bin/README.md) — configure a proxy script to run the sample app
3. [2-alexa-server](2-alexa-server/README.md) — build and run the server
4. [3-electron-app](3-electron-app/README.md) — run the Electron app
5. [4-bluetooth-fix](4-bluetooth-fix/README.md) — [optionally] fix audio issues with Bluetooth speakers

# Happy hacking!

If you do something fun with this code, share with the world! Feel free to reach out to me [on Gitter](https://gitter.im/iafan).
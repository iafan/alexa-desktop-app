This script builds [AVS Device SDK](https://github.com/alexa/avs-device-sdk) with the support for wake word from [KITT.AI / Snowboy](https://snowboy.kitt.ai/),
and then builds the sample console app that comes as a part of AVS Device SDK.

The build script assumes that AVS SDK is installed under ~/sdk-folder
(as per original installation instructions from AVS Device SDK).

After finishing this step you will get a console app

# Prerequisites

The instructions below and the provided helper scripts are for Mac OS. You should be able to run this app to some extent on other platforms supported by AVS Device SDK, but this will require adjusting things here and there.

### _Q: Will it work on Raspberry PI?_

I believe so (though I didn't try that). AVS Device SDK supports Raspberry PI by itself, and the server app (written in Go) can be compile on Raspberry PI as well. After running the server, you will be able to open the web UI in Raspberry's browser, which means you will have your Pi-powered Alexa device with a screen.

## Install AVS Device SDK

Go to https://github.com/alexa/avs-device-sdk#get-started and follow the installation instructions. Use the default `$HOME/sdk-folder` directory as your SDK root (otherwise you will need to fix paths in provided scripts/configs).

I recommend compiling the SDK and running the sample app without wake word first, to check that this part of the setup is working properly.

Use [../2-alexa-server/bin/AlexaClientSDKConfig.json.sample](../2-alexa-server/bin/AlexaClientSDKConfig.json.sample) file as a template for your configuration file.

## Install Snowboy

Check out https://github.com/alexa/avs-device-sdk/wiki/Build-Options#kittai page on prerequisites for KITT.AI.

Create `$HOME/sdk-folder/third-party/snowboy` folder and check out its repo there:

    cd $HOME/sdk-folder/third-party/snowboy
    git clone https://github.com/Kitt-AI/snowboy.git .

You will need to register at https://snowboy.kitt.ai/ and download the "Alexa" wake word model file (`alexa.umdl`). In the end you will need it located at `$HOME/sdk-folder/third-party/snowboy/resources/alexa.umdl`.

Use the provided script to rebuild the sample app with the wake word support:

    ./build.sh


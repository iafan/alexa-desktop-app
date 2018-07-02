You need to register your product at Amazon dev portal and get your own Alexa device client ID and product ID.

Then copy `AlexaClientSDKConfig.json.sample` to `AlexaClientSDKConfig.json`
and replace the following placehodlers:

1. `{MY_CLIENT_ID}` with the client ID (it looks like `amzn1.application-oa2-client.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`)

2. `{MY_PRODUCT_ID}` with the product ID that you defined when registering it. It can be pretty much any string, e.g. `my_desktop_app`. This ID will be visible in the Alexa app.

3. `{MY_HOME_PATH}` with the absolute path to your home directory (where you have `sdk-folder` folder). It will be something like `/Users/johndoe`. Note that `~` won't work.

To test if your config works, run:

    ./sample-app-runner

This will start your console app, and you can say "Alexa" (or whatever wake word you enabled) and start the conversation.
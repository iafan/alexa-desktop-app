If you're playing back your audio through a Bluetooth speaker, you may experience an issue where the first second of Alexa's response is being cut off (because it takes some time for Bluetooth connection to wake up after idling). If that's the case with your setup, run the provided script for as long as your Alexa app is running:

    ./playsilence.sh

This will continuously play silence and thus keep Bluetooth connection from sleeping. Not the best possible solution, but it does the job.
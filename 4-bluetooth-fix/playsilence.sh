#!/bin/sh
echo "Infinite playback of silence started (to keep Bluetooth connection open)"
echo "Press [CTRL+C] to stop."

gst-launch-1.0 audiotestsrc wave=4 ! autoaudiosink >/dev/null 2>&1

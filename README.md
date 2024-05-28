# go-mqttclient

I had a need for an mqtt client on an odd armv7 device, so I put this together to have an easy way to download it. Maybe this exists out there already, or there is some other solution I shold have used, but this is what I came up with.

## Usage
```./go-mqttclient -message 'sample message' -server tcp://mqtt-server:1883 -topic this/is/my/test/topic -username mqtt_user -password password123```

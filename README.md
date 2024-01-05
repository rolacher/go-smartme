# go-smartme
GoLang library to connect to smart-me API written in go language. This is not a complete impementation of the API. https://github.com/rolacher/go-smartme-tools is an example of how to use this library.

Concepts for this lib have been taken from https://github.com/guckykv/freeathome-go-fahapi, thank you for the inspiration. 

## smartmeapi
This package reads devices and values from the [Smart-me API](https://api.smart-me.com/swagger/index.html).

Supported are GETs to:
* /Devices
* /Devices/{id}
* /Values/{id}
* /ValuesInPast/{id}
* /ValuesInPastMultiple/{id}

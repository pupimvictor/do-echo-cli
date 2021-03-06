# Echo Cli

This is a CLI that yells to a service and expects to get an Echo back!

#### OpenAPI
The CLI (and [service](https://github.com/pupimvictor/do-echo-srv)) were created using [go-swagger](https://github.com/go-swagger/go-swagger/), a golang implementation of Swagger 2.0 (aka OpenAPI 2.0), to create a contract between the client and the server.

#### Structure
The go-swagger code generator provides a skeleton of a client for the api, based on the swagger.yml spec. This tool uses [github.com/urfave/cli] to help with the CLI params and control the app execution.

All the code for the app is in main.go. All the other files are auto generated by go-swagger to create an robust client api.

#### Authentication
You will need a x-api-token to be authorized to request for an echo. make sure you pass it using `--token <the-token>`

#### Usage
This app supports single message yelling and batch yelling using a list of messages from a file. (you can find a sample file in `resources` folder)

To use the app you can start by building it using `go build .`
To yell at the service do:
 - single message: `./do-echo-cli --host <server-host>:<port> --token <x-api-token> yell --msg <message to be echoed>`

 - batch messages: `./do-echo-cli --host <server-host>:<port> --token <x-api-token> batch --file ./resources/msgs.txt`

#### Help
Type ./do-echo-cli --help



 ;)---------- playing with gitflow releases -----------
---------- feature 1 -----------
---------- change 1 -----------
---------- change 2 -----------
---------- rebase develop. breaks build -----------
---------- change 3 -----------

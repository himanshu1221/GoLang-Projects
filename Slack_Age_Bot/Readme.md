# Slack Age Bot

## What is this project about ?

In this project we will be building a slack bot to calculate age using slack API

## Dependencies

- `github.com/shomali11/slacker` - Built on top of the Slack API github.com/slack-go/slack, Slacker is a low-friction framework for creating Slack Bots.
- `github.com/joho/godotenv` - gotenv package in golang can be used to load .env or environment file in golang

## Topics I have Learnt
- `context.WithCancel` - it is used to do cancellation and request scoped values (req.body kind of stuff from javascript)
- `context.Background()` - Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
- `slacker.NewClient` - NewClient creates a new client using the Slack API
- `CommandEvents` - CommandEvents returns read only command events channel
- `godotenv.Load` - Used to load env file
- `Slack API` - learnt using slack API
- `Socket Mode` - Socket Mode allows your app to communicate with Slack via a WebSocket URL. WebSockets use a bidirectional stateful protocol with low latency to communicate between two parties—in this case, Slack and your app.
- `WebSocket Protocol` - The WebSocket Protocol enables two-way communication between a client
  running untrusted code in a controlled environment to a remote host
  that has opted-in to communications from that code.
- `.env` - Learnt how to setup .env
- `slacker.NewClient` - It creates a new client using the Slack API
- `CommandEvents` - CommandEvents returns read only command events channel
- `Defer` - used to delay the execution of a function until the surrounding function completes

## Resources
- (Dontenv Setup)[https://golangbyexample.com/load-env-fiie-golang/]
- (Slacker Documentation)[https://pkg.go.dev/github.com/shomali11/slacker#section-readme]
- (Slack API)[https://api.slack.com/]
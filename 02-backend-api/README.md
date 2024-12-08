# Project 02 - Backend API

The goal of this project is to create an http+json API for a calculator service.

## Requirements

The API should conform to the given OpenAPI spec found in this directory, which can also be viewed at this URL.

### Production Ready

In order to make this API more production ready, there's a few other requiements you'll need to consider

#### Input Validation

You should never trust input from the user. This means you'll need to be sure to validate and sanitize any inputs. Division by zero is a common cause
panics when it comes to applications, so you'll want to make sure you're handling it.

#### Error feedback

Additionally, it's a good idea to let the user know when they've made a mistake with input, so they can fix it if the mistake was innocent.

#### Logging

In order to be able to debug issues that occur, you're going to want to log out each request as it comes in, as well as any associated data such as the status code, ip address, and what the request path was.

For logging in Go, I recommend using the `log/slog` package, which provides structured logging in either JSON, or Text format. You can create a logger using the following snippet:

##### Text

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
```

##### JSON

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
```

## Recommended Packages

Some of the packages that I used for my implementation include:

- `net/http` - This is the http package of the standard library which I use for routing and setting up an http server. If you want to know how to use this package for advanced routing, [I have a video you can check out](https://youtu.be/H7tbjKFSg58).
- `encoding/json` - This package is used for encoding and decoding JSON from the request and to the response bodies.
- `log/slog` - For structured logging
- `github.com/rs/cors` - For cors, if you need it.

## Additional Tasks

- Add in a middleware to print the executed path.

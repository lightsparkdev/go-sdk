# Lightspark Go SDK - v0.16.2

The Lightspark Go SDK provides a convenient way to interact with the Lightspark services from applications written in
the Go.

**_WARNING: This SDK is in version 0.16.2 (active development). It means that its APIs may not be fully stable. Please
expect that changes to the APIs may happen until we move to v1.0.0._**

## Documentation

The documentation for this SDK (installation, usage, etc.) is available at https://docs.lightspark.com/lightspark-sdk/getting-started?language=Go

## Sample code

For your convenience, we included an example that shows you how to use the SDK. Open the file
`examples/example.go` and make sure to update the variables at the top of the page with your
information, then run it using `go run`:

```
go run examples/example.go
```

There is also a sample LNURL server in `examples/lnurl-server`. This can be run by building with
`go build` and running the resulting binary:

```
cd examples/lnurl-server
go build
./lnurl-server
```

See `examples/lnurl-server/config.go` for configuring the server and
`examples/lnurl-server/server.go` for more information about the API it provides.

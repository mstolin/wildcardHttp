# WildcardHttp

This is a simple HTTP server, that serves any routes and prints the HTTP request information.
The purpose of this project is to have a simple and lightweight HTTP serves for debugging in
single, simple, and zero-dependency go module.

## Usage

First, build the binary using `make build`.
Then, use the CLI tool as the following:

```
whttp <host> [<port>]
```

Either provide the port as part of a valid hostname or separately.
For example, the following are equal and serve at port 5000 on localhost:

```
whttp :5000
whttp localhost:5000
whttp localhost 5000
```

## Docker

Pull the image from [https://hub.docker.com/r/marcelstolin/wildcardhttp](https://hub.docker.com/r/marcelstolin/wildcardhttp).

The provided [Dockerfile](./Dockerfile) exposes the HTTP server at port 80.
The [Makefile](./Makefile) supports Docker and Podman, and provides the following commands:

```
make build-docker # Build the OCI image.
make run-docker   # Start the container in background. Set port with PORT=8000, otherwise 5000 is used.
make stop-docker  # Stop and remove the container.
```

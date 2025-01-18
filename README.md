# WildcardHttp

This ia asimple HTTP server, that serves any routes and prints the HTTP request information.
The purpose of this project is to have a simple and lightweight HTTP serves for debugging.
Additionally, the development goal is to have a single, simple, and zero-dependency go module.

## Usage

Use the CLI tool as the following:

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

The provided [Dockerfile](./Dockerfile) exposes the HTTP server at port 80. You can use as follows:

```
docker build -t wildcardhttp .
docker run --rm -p 127.0.0.1:5000:80 localhost/whttp:latest
```

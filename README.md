# quic-echo

Echo server, client over the QUIC protocol.

## build

Run `make` in the root directory. Execute the server and client as two different processes in order.

## files

### server.go

Creates a QUIC echo server.

```bash
$ ./server -h
Usage of ./server:
  -hostname string
        hostname/ip of the server (default "localhost")
  -port string
        port number of the server (default "4242")
```

### client.go

Creates a QUIC client to talk to the echo server.

```bash
$ ./client -h
Usage of ./client:
  -hostname string
        hostname/ip of the server (default "localhost")
  -necho int
        number of echos (default 100)
  -port string
        port number of the server (default "4242")
```

## author

Sayan Goswami

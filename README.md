# quic-echo

Echo client, server over the QUIC protocol.

## server.go

```bash
$ ./server -h
Usage of ./server:
  -hostname string
        hostname/ip of the server (default "localhost")
  -port string
        port number of the server (default "4242")
```

## client.go

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

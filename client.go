package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"strconv"

	quic "github.com/lucas-clemente/quic-go"
)

func main() {
	const addr = "localhost:4242"

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo"},
	}

	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		panic(err)
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}

	counter := 0

	for {
		message := strconv.Itoa(counter)
		counter++

		log.Printf("Client: Sending '%s'\n", message)
		_, err = stream.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		log.Println("Done. Waiting for echo.")

		buf := make([]byte, 1024)
		_, err = io.ReadFull(stream, buf)
		if err != nil {
			panic(err)
		}
		log.Printf("Client: Got '%s'\n", buf)
	}
}

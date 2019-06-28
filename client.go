package main

import (
	"context"
	"crypto/tls"
	"flag"
	"io"
	"log"
	"strconv"

	quic "github.com/lucas-clemente/quic-go"
)

func main() {
	hostName := flag.String("hostname", "localhost", "hostname/ip of the server")
	portNum := flag.String("port", "4242", "port number of the server")
	numEcho := flag.Int("necho", 100, "number of echos")

	flag.Parse()

	addr := *hostName + ":" + *portNum

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

		buff := make([]byte, len(message))
		_, err = io.ReadFull(stream, buff)
		if err != nil {
			panic(err)
		}

		log.Printf("Client: Got '%s'\n", string(buff))

		if counter == *numEcho {
			break
		}
	}
}

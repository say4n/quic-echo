package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"math/big"

	quic "github.com/lucas-clemente/quic-go"
)

func main() {
	hostName := "localhost"
	portNum := "4242"
	addr := hostName + ":" + portNum
	log.Println("Server running @", addr)

	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		panic(err)
	}
	sess, err := listener.Accept(context.Background())
	if err != nil {
		panic(err)
	}

	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 1024)
		_, err = io.ReadFull(stream, buf)
		if err != nil {
			panic(err)
		}

		msg := string(buf)

		log.Printf("Server: Got '%s'", msg)
		log.Printf("Client: Sending '%s", msg)

		_, err = stream.Write([]byte(msg))
	}
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo"},
	}
}

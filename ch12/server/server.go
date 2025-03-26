package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/testaquatic/NetworkProgrammingWithGo/ch12/housework/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr, certFn, keyFn string

func init() {
	flag.StringVar(&addr, "address", "localhost:34443", "listen port")
	flag.StringVar(&certFn, "cert", "cert.pem", "certificate file")
	flag.StringVar(&keyFn, "key", "key.pem", "private key file")
}

// 책의 코드가 작동하지 않아서 https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go 이곳의 코드를 보고 변경했다.
func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(certFn, keyFn)
	if err != nil {
		log.Fatal(err)
	}
	tlsConfig := credentials.NewTLS(&tls.Config{
		Certificates:     []tls.Certificate{cert},
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP256},
	})

	grpcServer := grpc.NewServer(grpc.Creds(tlsConfig))
	housework.RegisterRobotMaidServer(grpcServer, &Rosie{})

	listener, err := net.Listen("tcp", addr)
	fmt.Printf("Listening for TLS connection on %s ...", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

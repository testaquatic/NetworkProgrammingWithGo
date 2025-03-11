package main

import (
	"flag"
	"log"
	"os"

	"github.com/testaquatic/NetworkProgrammingWithGo/ch06/tftp"
)

var (
	address = flag.String("a", ":6999", "listening address")
	payload = flag.String("p", "payload.svg", "file to serve to client")
)

func main() {
	flag.Parse()

	p, err := os.ReadFile(*payload)
	if err != nil {
		log.Fatal(err)
	}

	s := tftp.Server{Payload: p}
	log.Fatal(s.ListenAndServe(*address))
}

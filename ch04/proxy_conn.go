package main

import (
	"io"
	"log"
	"net"
)

func proxyConn(source, destination string) error {
	connSource, err := net.Dial("tcp", source)
	if err != nil {
		return err
	}
	defer connSource.Close()

	connDestination, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer connDestination.Close()

	// `connDestination` -> `connSource`
	go func() {
		_, err := io.Copy(connSource, connDestination)
		if err != nil {
			log.Println(err)
		}
	}()

	// `connSource` -> `connDestination`
	_, err = io.Copy(connDestination, connSource)

	return err
}

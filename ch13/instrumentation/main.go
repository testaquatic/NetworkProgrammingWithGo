package main

import (
	"flag"
	"net/http"

	"github.com/testaquatic/NetworkProgrammingWithGo/ch13/instrumentation/metrics"
)

var (
	metricsAddr = flag.String("metrics", "127.0.0.1:8081", "metrics listen address")
	webAddr = flag.String("web", "127.0.0.1:8082", "web listen address")
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	metrics.Requests
}
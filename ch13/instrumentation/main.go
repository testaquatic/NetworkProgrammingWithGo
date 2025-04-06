package main

import "flag"

var (
	metricsAddr = flag.String("metrics", "127.0.0.1:8081", "metrics listen address")
	webAddr = flag.String("web", "127.0.0.1:8082", "web listen address")
)


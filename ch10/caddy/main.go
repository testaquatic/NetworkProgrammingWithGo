package main

import (
	cmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/testaquatic/NetworkProgrammingWithGo/ch10/caddy-restrict-prefix"
	_ "github.com/testaquatic/NetworkProgrammingWithGo/ch10/caddy-toml-adapter"
)

func main() {
	cmd.Main()
}

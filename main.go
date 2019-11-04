package main

import "github.com/genku-m/sample_server/server"

func main() {
	cfg := server.NewConfig()
	svr := server.NewServer(cfg)
	svr.Listen()
}

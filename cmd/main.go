package main

import (
	"webservice/internals/app"
	"webservice/internals/cfg"
)

func main() { 
	config := cfg.LoadConfig()

	server := app.NewServer(config)

	server.Start()

	defer server.Shutdown()
}

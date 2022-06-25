package main

import (
	"webservice/internal/app"
	"webservice/internal/cfg"
)

func main() { 
	config := cfg.LoadConfig()

	server := app.NewServer(config)

	server.Start()

	defer server.Shutdown()
}

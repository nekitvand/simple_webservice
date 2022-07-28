package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"webservice/internal/app"
	"webservice/internal/cfg"
)

func main() { 

	ctx,cancel := context.WithCancel(context.Background())

	end := make(chan os.Signal,1)
	signal.Notify(end,syscall.SIGINT,syscall.SIGQUIT,syscall.SIGTERM)

	config := cfg.LoadConfig()

	server := app.NewServer(ctx,config)

	go func(server *app.AppServer) {
		select{
		case sig := <-end:
			fmt.Println("Shutdown with ",sig)
			server.Shutdown()
		case <-ctx.Done():
			fmt.Println("Shutdown with ctx")
			server.Shutdown()
		}
		defer cancel()

	}(server)

	server.Start()
	
}

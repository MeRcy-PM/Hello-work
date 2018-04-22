package main

import (
	"context"
	"fmt"
)

func main() {
	config := NewDefaultConfiguration()
	config.ParseArguments()
	fmt.Println(config)

	if config.IsServer {
		server := NewServer(config)
		if err := server.Run(context.Background()); err != nil {
			return
		}
	}
}

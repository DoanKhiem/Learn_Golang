package main

import (
	"context"
	"fmt"
	"github.com/dreamsoftcode-io/orders-api/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start the server: ", err)
	}
}

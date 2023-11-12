package main

import (
	"gotodo/internal/app"
	"log"
	"os"
)

func main() {
	if err := app.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

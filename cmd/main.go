package main

import (
	"github.com/memclutter/gotodo/internal/cli"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := cli.NewApp().Run(os.Args); err != nil {
		logrus.Fatalf("error app run: %v", err)
	}
}

package main

import (
	"context"
	"github.com/dracher/easy-commit/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Cli().Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

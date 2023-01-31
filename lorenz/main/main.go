package main

import (
	"fmt"
	"log"

	"github.com/Abound-art/starter-go/abound"
	"github.com/Abound-art/starter-go/lorenz"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error running algo: %v", err)
	}
}

func run() error {
	var config *lorenz.Config
	if err := abound.LoadConfig(&config); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	img := config.Run()

	if err := abound.WriteImage(img); err != nil {
		return fmt.Errorf("writing image: %w", err)
	}
	return nil
}

// Command algo implements a basic algorithm that can be run on the ABOUND
// platform.
package main

import (
	"fmt"
	"log"

	"github.com/Abound-art/starter-go/abound"
	"github.com/Abound-art/starter-go/algo"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error running algo: %v", err)
	}
}

func run() error {
	var config *algo.Config
	if err := abound.LoadConfig(&config); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	img := algo.Run(config)

	if err := abound.WritePNG(img); err != nil {
		return fmt.Errorf("writing image: %w", err)
	}

	return nil
}

// Package abound provides some helper utilities for reading configuration and
// writing output images.
package abound

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

const (
	ConfigPathEnvVar = "ABOUND_CONFIG_PATH"
	OutputPathEnvVar = "ABOUND_OUTPUT_PATH"
)

// LoadConfig parses an ABOUND configuration from its specified location into
// the provided structure.
func LoadConfig(v any) error {
	cfgPath := os.Getenv(ConfigPathEnvVar)
	if cfgPath == "" {
		return fmt.Errorf("environment variable for algo configuration (%q) was not set, are you running on ABOUND?", ConfigPathEnvVar)
	}
	f, err := os.Open(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&v); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	return nil
}

// WriteSVG writes an SVG-formatted string to the target output path.
func WriteSVG(svg string) error {
	f, err := os.Create(os.Getenv(OutputPathEnvVar))
	if err != nil {
		return fmt.Errorf("failed to create file for svg: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(svg); err != nil {
		return fmt.Errorf("failed to write svg content to file: %w", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close svg file: %w", err)
	}
	return nil
}

// WritePNG writes a PNG-formatted image to the target output path.
func WritePNG(img image.Image) error {
	f, err := os.Create(os.Getenv(OutputPathEnvVar))
	if err != nil {
		return fmt.Errorf("failed to create ouput image file: %w", err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode output image: %w", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close output image: %w", err)
	}

	return nil
}

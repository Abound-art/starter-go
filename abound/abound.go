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
	ConfigPathEnvVar = "THEAPP_CONFIG_PATH"
	OutputPathEnvVar = "THEAPP_OUTPUT_PATH"
)

func LoadConfig(v interface{}) error {
	f, err := os.Open(os.Getenv(ConfigPathEnvVar))
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&v); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	return nil
}

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

func WriteImage(img image.Image) error {
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

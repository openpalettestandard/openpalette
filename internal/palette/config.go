package palette

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/openpalettestandard/openpalette/internal/types"
)

type ConfigFile struct {
	Version  string                   `json:"version"`
	Variants map[string]ConfigVariant `json:"variants"`
}

type ConfigVariant struct {
	Name   string                 `json:"name"`
	Emoji  string                 `json:"emoji"`
	Dark   bool                   `json:"dark"`
	Colors map[string]ConfigColor `json:"colors"`
}

type ConfigColor struct {
	Name   string `json:"name"`
	Hex    string `json:"hex"`
	Accent bool   `json:"accent"`
}

func LoadFromFile(filename string) ([]types.RawVariant, string, error) {
	if filename == "" {
		return getRawVariants(), "", nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", fmt.Errorf("reading config file: %w", err)
	}

	var config ConfigFile
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, "", fmt.Errorf("parsing JSON config: %w", err)
	}

	return convertConfigToRawVariants(config), config.Version, nil
}

func convertConfigToRawVariants(config ConfigFile) []types.RawVariant {
	var variants []types.RawVariant

	for id, variant := range config.Variants {
		rawVariant := types.RawVariant{
			ID:    id,
			Name:  variant.Name,
			Emoji: variant.Emoji,
			Dark:  variant.Dark,
		}

		for colorID, color := range variant.Colors {
			rawVariant.PaletteColors = append(rawVariant.PaletteColors, types.RawPaletteColor{
				ID:     colorID,
				Name:   color.Name,
				Hex:    color.Hex,
				Accent: color.Accent,
			})
		}

		variants = append(variants, rawVariant)
	}

	return variants
}

func GenerateExampleConfig(filename string) error {
	config := ConfigFile{
		Version: "1.0.0",
		Variants: map[string]ConfigVariant{
			"latte": {
				Name:  "Latte",
				Emoji: "🌻",
				Dark:  false,
				Colors: map[string]ConfigColor{
					"rosewater": {Name: "Rosewater", Hex: "#dc8a78", Accent: true},
					"flamingo":  {Name: "Flamingo", Hex: "#dd7878", Accent: true},
					"pink":      {Name: "Pink", Hex: "#ea76cb", Accent: true},
					"red":       {Name: "Red", Hex: "#d20f39", Accent: true},
					"text":      {Name: "Text", Hex: "#4c4f69", Accent: false},
					"base":      {Name: "Base", Hex: "#eff1f5", Accent: false},
				},
			},
			"mocha": {
				Name:  "Mocha",
				Emoji: "🌙",
				Dark:  true,
				Colors: map[string]ConfigColor{
					"rosewater": {Name: "Rosewater", Hex: "#f5e0dc", Accent: true},
					"flamingo":  {Name: "Flamingo", Hex: "#f2cdcd", Accent: true},
					"pink":      {Name: "Pink", Hex: "#f5c2e7", Accent: true},
					"red":       {Name: "Red", Hex: "#f38ba8", Accent: true},
					"text":      {Name: "Text", Hex: "#cdd6f4", Accent: false},
					"base":      {Name: "Base", Hex: "#1e1e2e", Accent: false},
				},
			},
		},
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling config: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

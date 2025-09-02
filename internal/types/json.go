package types

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (pr PaletteResult) MarshalJSON() ([]byte, error) {
	var buf strings.Builder
	buf.WriteString("{")

	buf.WriteString(fmt.Sprintf(`"version":"%s"`, pr.Version))

	variantOrder := []string{"latte", "frappe", "macchiato", "mocha"}
	for _, variantName := range variantOrder {
		if variant, exists := pr.Variants[variantName]; exists {
			buf.WriteString(",")
			buf.WriteString(fmt.Sprintf(`"%s":`, variantName))

			variantJSON, err := variant.MarshalJSON()
			if err != nil {
				return nil, err
			}
			buf.Write(variantJSON)
		}
	}

	buf.WriteString("}")
	return []byte(buf.String()), nil
}

func (pv PaletteVariant) MarshalJSON() ([]byte, error) {
	var buf strings.Builder
	buf.WriteString("{")

	buf.WriteString(fmt.Sprintf(`"name":"%s"`, pv.Name))
	buf.WriteString(fmt.Sprintf(`,"emoji":"%s"`, pv.Emoji))
	buf.WriteString(fmt.Sprintf(`,"order":%d`, pv.Order))
	buf.WriteString(fmt.Sprintf(`,"dark":%t`, pv.Dark))

	buf.WriteString(`,"colors":{`)
	colorOrder := []string{
		"rosewater", "flamingo", "pink", "mauve", "red", "maroon", "peach", "yellow",
		"green", "teal", "sky", "sapphire", "blue", "lavender", "text", "subtext1",
		"subtext0", "overlay2", "overlay1", "overlay0", "surface2", "surface1",
		"surface0", "base", "mantle", "crust",
	}

	first := true
	for _, colorName := range colorOrder {
		if color, exists := pv.PaletteColors[colorName]; exists {
			if !first {
				buf.WriteString(",")
			}
			first = false

			buf.WriteString(fmt.Sprintf(`"%s":`, colorName))
			colorJSON, err := json.Marshal(color)
			if err != nil {
				return nil, err
			}
			buf.Write(colorJSON)
		}
	}
	buf.WriteString("}")

	buf.WriteString(`,"ansiColors":{`)
	ansiOrder := []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}

	first = true
	for _, ansiName := range ansiOrder {
		if ansiColor, exists := pv.AnsiPaletteColors[ansiName]; exists {
			if !first {
				buf.WriteString(",")
			}
			first = false

			buf.WriteString(fmt.Sprintf(`"%s":`, ansiName))
			ansiJSON, err := json.Marshal(ansiColor)
			if err != nil {
				return nil, err
			}
			buf.Write(ansiJSON)
		}
	}
	buf.WriteString("}")

	buf.WriteString("}")
	return []byte(buf.String()), nil
}

func WriteJSON(palette PaletteResult, writer io.Writer) error {
	var jsonData []byte
	var err error

	jsonData, err = json.MarshalIndent(palette, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	_, err = writer.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}

func WriteJSONFile(palette PaletteResult, filename string) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(palette, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

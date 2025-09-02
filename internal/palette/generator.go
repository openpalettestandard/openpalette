package palette

import (
	"math"
	"strings"

	"github.com/openpalettestandard/openpalette/internal/color"
	"github.com/openpalettestandard/openpalette/internal/types"
)

func Generate() types.PaletteResult {
	rawVariants, _, _ := LoadFromFile("")
	return GenerateFromVariants(rawVariants, "")
}

func GenerateFromConfig(configFile string) (types.PaletteResult, error) {
	rawVariants, configVersion, err := LoadFromFile(configFile)
	if err != nil {
		return types.PaletteResult{}, err
	}

	return GenerateFromVariants(rawVariants, configVersion), nil
}

func GenerateFromVariants(rawVariants []types.RawVariant, version string) types.PaletteResult {
	ansiMappings := getANSIMappings()

	result := types.PaletteResult{
		Version:  version,
		Variants: make(map[string]types.PaletteVariant),
	}

	for variantIndex, rawVariant := range rawVariants {
		variant := types.PaletteVariant{
			Name:              rawVariant.Name,
			Emoji:             rawVariant.Emoji,
			Order:             variantIndex,
			Dark:              rawVariant.Dark,
			PaletteColors:     make(map[string]types.PaletteColor),
			AnsiPaletteColors: make(map[string]types.ANSIColor),
		}

		for colorIndex, rawColor := range rawVariant.PaletteColors {
			variant.PaletteColors[rawColor.ID] = ProcessColor(rawColor, colorIndex)
		}

		for ansiIndex, ansiName := range []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"} {
			ansiMapping := ansiMappings[ansiName]
			variant.AnsiPaletteColors[ansiName] = ProcessANSIColor(ansiName, ansiMapping, ansiIndex, variant, rawVariant.Dark)
		}

		result.Variants[rawVariant.ID] = variant
	}

	return result
}

func ProcessColor(rawColor types.RawPaletteColor, order int) types.PaletteColor {
	newColor := color.NewColor(rawColor.Hex)

	coords := newColor.ToSRGBGamut()
	rgb := types.RGB{
		R: int(math.Round(coords[0] * 255)),
		G: int(math.Round(coords[1] * 255)),
		B: int(math.Round(coords[2] * 255)),
	}

	hsl := color.TinyColorHSL(rawColor.Hex)

	return types.PaletteColor{
		Name:   rawColor.Name,
		Order:  order,
		Hex:    rawColor.Hex,
		RGB:    rgb,
		HSL:    hsl,
		Accent: rawColor.Accent,
	}
}

func ProcessANSIColor(ansiName string, mapping types.ANSIMapping, order int, variant types.PaletteVariant, isDark bool) types.ANSIColor {
	var normalColor *color.Color
	var normalName string

	if ansiName == "black" {
		if isDark {
			normalColor = color.NewColor(findColorHex(variant, "surface1"))
		} else {
			normalColor = color.NewColor(findColorHex(variant, "subtext1"))
		}
		normalName = "Black"
	} else if ansiName == "white" {
		if isDark {
			normalColor = color.NewColor(findColorHex(variant, "subtext0"))
		} else {
			normalColor = color.NewColor(findColorHex(variant, "surface2"))
		}
		normalName = "White"
	} else {
		normalColor = color.NewColor(findColorHex(variant, mapping.Mapping))
		normalName = strings.Title(ansiName)
	}

	brightColor := normalColor.Clone()

	if ansiName != "black" && ansiName != "white" {
		lch := brightColor.GetLCH()

		if isDark {
			lch.SetL(lch.L() * 0.94)
			lch.SetC(lch.C() + 8)
		} else {
			lch.SetL(lch.L() * 1.09)
			lch.SetC(lch.C() + 0)
		}
		lch.SetH(lch.H() + 2)
	} else {
		if ansiName == "black" {
			if isDark {
				brightColor = color.NewColor(findColorHex(variant, "surface2"))
			} else {
				brightColor = color.NewColor(findColorHex(variant, "subtext0"))
			}
		} else {
			if isDark {
				brightColor = color.NewColor(findColorHex(variant, "subtext1"))
			} else {
				brightColor = color.NewColor(findColorHex(variant, "surface1"))
			}
		}
	}

	normalVariant := colorToANSIVariant(normalColor, normalName, mapping.NormalCode)
	brightVariant := colorToANSIVariant(brightColor, "Bright "+normalName, mapping.BrightCode)

	return types.ANSIColor{
		Name:   normalName,
		Order:  order,
		Normal: normalVariant,
		Bright: brightVariant,
	}
}

func getRawVariants() []types.RawVariant {
	return []types.RawVariant{
		{
			ID:    "latte",
			Name:  "Latte",
			Emoji: "🌻",
			Dark:  false,
			PaletteColors: []types.RawPaletteColor{
				{ID: "rosewater", Name: "Rosewater", Hex: "#dc8a78", Accent: true},
				{ID: "flamingo", Name: "Flamingo", Hex: "#dd7878", Accent: true},
				{ID: "pink", Name: "Pink", Hex: "#ea76cb", Accent: true},
				{ID: "mauve", Name: "Mauve", Hex: "#8839ef", Accent: true},
				{ID: "red", Name: "Red", Hex: "#d20f39", Accent: true},
				{ID: "maroon", Name: "Maroon", Hex: "#e64553", Accent: true},
				{ID: "peach", Name: "Peach", Hex: "#fe640b", Accent: true},
				{ID: "yellow", Name: "Yellow", Hex: "#df8e1d", Accent: true},
				{ID: "green", Name: "Green", Hex: "#40a02b", Accent: true},
				{ID: "teal", Name: "Teal", Hex: "#179299", Accent: true},
				{ID: "sky", Name: "Sky", Hex: "#04a5e5", Accent: true},
				{ID: "sapphire", Name: "Sapphire", Hex: "#209fb5", Accent: true},
				{ID: "blue", Name: "Blue", Hex: "#1e66f5", Accent: true},
				{ID: "lavender", Name: "Lavender", Hex: "#7287fd", Accent: true},
				{ID: "text", Name: "Text", Hex: "#4c4f69", Accent: false},
				{ID: "subtext1", Name: "Subtext 1", Hex: "#5c5f77", Accent: false},
				{ID: "subtext0", Name: "Subtext 0", Hex: "#6c6f85", Accent: false},
				{ID: "overlay2", Name: "Overlay 2", Hex: "#7c7f93", Accent: false},
				{ID: "overlay1", Name: "Overlay 1", Hex: "#8c8fa1", Accent: false},
				{ID: "overlay0", Name: "Overlay 0", Hex: "#9ca0b0", Accent: false},
				{ID: "surface2", Name: "Surface 2", Hex: "#acb0be", Accent: false},
				{ID: "surface1", Name: "Surface 1", Hex: "#bcc0cc", Accent: false},
				{ID: "surface0", Name: "Surface 0", Hex: "#ccd0da", Accent: false},
				{ID: "base", Name: "Base", Hex: "#eff1f5", Accent: false},
				{ID: "mantle", Name: "Mantle", Hex: "#e6e9ef", Accent: false},
				{ID: "crust", Name: "Crust", Hex: "#dce0e8", Accent: false},
			},
		},
	}
}

func getANSIMappings() map[string]types.ANSIMapping {
	return map[string]types.ANSIMapping{
		"black": {
			NormalCode: 0,
			BrightCode: 8,
			Mapping:    "",
		},
		"red": {
			NormalCode: 1,
			BrightCode: 9,
			Mapping:    "red",
		},
		"green": {
			NormalCode: 2,
			BrightCode: 10,
			Mapping:    "green",
		},
		"yellow": {
			NormalCode: 3,
			BrightCode: 11,
			Mapping:    "yellow",
		},
		"blue": {
			NormalCode: 4,
			BrightCode: 12,
			Mapping:    "blue",
		},
		"magenta": {
			NormalCode: 5,
			BrightCode: 13,
			Mapping:    "pink",
		},
		"cyan": {
			NormalCode: 6,
			BrightCode: 14,
			Mapping:    "teal",
		},
		"white": {
			NormalCode: 7,
			BrightCode: 15,
			Mapping:    "",
		},
	}
}

func findColorHex(variant types.PaletteVariant, colorID string) string {
	if color, exists := variant.PaletteColors[colorID]; exists {
		return color.Hex
	}
	return "#000000"
}

func colorToANSIVariant(c *color.Color, name string, code int) types.ANSIVariant {
	coords := c.ToSRGBGamut()
	rgb := types.RGB{
		R: int(math.Round(coords[0] * 255)),
		G: int(math.Round(coords[1] * 255)),
		B: int(math.Round(coords[2] * 255)),
	}

	hex := c.ToString()
	hsl := color.TinyColorHSL(hex)

	return types.ANSIVariant{
		Name: name,
		Hex:  hex,
		RGB:  rgb,
		HSL:  hsl,
		Code: code,
	}
}

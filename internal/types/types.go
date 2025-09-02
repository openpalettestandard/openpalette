package types

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type HSL struct {
	H float64 `json:"h"`
	S float64 `json:"s"`
	L float64 `json:"l"`
}

type PaletteColor struct {
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Hex    string `json:"hex"`
	RGB    RGB    `json:"rgb"`
	HSL    HSL    `json:"hsl"`
	Accent bool   `json:"accent"`
}

type ANSIVariant struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
	RGB  RGB    `json:"rgb"`
	HSL  HSL    `json:"hsl"`
	Code int    `json:"code"`
}

type ANSIColor struct {
	Name   string      `json:"name"`
	Order  int         `json:"order"`
	Normal ANSIVariant `json:"normal"`
	Bright ANSIVariant `json:"bright"`
}

type ANSIMapping struct {
	NormalCode int
	BrightCode int
	Mapping    string
}

type PaletteVariant struct {
	Name              string                  `json:"name"`
	Emoji             string                  `json:"emoji"`
	Order             int                     `json:"order"`
	Dark              bool                    `json:"dark"`
	PaletteColors     map[string]PaletteColor `json:"colors"`
	AnsiPaletteColors map[string]ANSIColor    `json:"ansiColors"`
}

type PaletteResult struct {
	Version  string                    `json:"version"`
	Variants map[string]PaletteVariant `json:"-"`
}

type RawPaletteColor struct {
	ID     string
	Name   string
	Hex    string
	Accent bool
}

type RawVariant struct {
	ID            string
	Name          string
	Emoji         string
	Dark          bool
	PaletteColors []RawPaletteColor
}


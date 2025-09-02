package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/openpalettestandard/openpalette/internal/types"
)

type Color struct {
	hex   string
	lch   [3]float64
	isLCH bool
}

func NewColor(hex string) *Color {
	return &Color{
		hex:   strings.TrimPrefix(hex, "#"),
		isLCH: false,
	}
}

func (c *Color) Clone() *Color {
	clone := &Color{
		hex:   c.hex,
		isLCH: c.isLCH,
	}
	if c.isLCH {
		clone.lch = c.lch
	}
	return clone
}

func (c *Color) ToString() string {
	if c.isLCH {
		return c.lchToHex()
	}
	return "#" + c.hex
}

func (c *Color) ToSRGBGamut() [3]float64 {
	var r, g, b float64

	if c.isLCH {
		r, g, b = c.lchToSRGB()
	} else {
		r, g, b = c.hexToSRGB()
	}

	r = clampFloat(r, 0, 1)
	g = clampFloat(g, 0, 1)
	b = clampFloat(b, 0, 1)

	return [3]float64{r, g, b}
}

func (c *Color) GetLCH() *LCHColor {
	if !c.isLCH {
		c.lch = c.hexToLCH()
		c.isLCH = true
	}

	return &LCHColor{color: c}
}

type LCHColor struct {
	color *Color
}

func (lch *LCHColor) L() float64 {
	return lch.color.lch[0]
}

func (lch *LCHColor) SetL(value float64) {
	lch.color.lch[0] = value
}

func (lch *LCHColor) C() float64 {
	return lch.color.lch[1]
}

func (lch *LCHColor) SetC(value float64) {
	lch.color.lch[1] = value
}

func (lch *LCHColor) H() float64 {
	return lch.color.lch[2]
}

func (lch *LCHColor) SetH(value float64) {
	lch.color.lch[2] = value
}

func (c *Color) hexToSRGB() (float64, float64, float64) {
	r, _ := strconv.ParseInt(c.hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(c.hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(c.hex[4:6], 16, 0)

	return float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0
}

func TinyColorHSL(hex string) types.HSL {
	hex = strings.TrimPrefix(hex, "#")
	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)

	rNorm := bound01(float64(r), 255)
	gNorm := bound01(float64(g), 255)
	bNorm := bound01(float64(b), 255)

	max := math.Max(math.Max(rNorm, gNorm), bNorm)
	min := math.Min(math.Min(rNorm, gNorm), bNorm)

	var h, s, l float64
	l = (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case rNorm:
			h = (gNorm - bNorm) / d
			if gNorm < bNorm {
				h += 6
			}
		case gNorm:
			h = (bNorm-rNorm)/d + 2
		case bNorm:
			h = (rNorm-gNorm)/d + 4
		}
		h /= 6
	}

	return types.HSL{
		H: h * 360,
		S: s,
		L: l,
	}
}

func bound01(value, max float64) float64 {
	if max == 360 {
		n := value
		if math.Abs(n-max) < 0.000001 {
			return 1.0
		}
		if n < 0 {
			return (math.Mod(n, max) + max) / max
		}
		return math.Mod(n, max) / max
	}

	n := math.Min(max, math.Max(0, value))
	if math.Abs(n-max) < 0.000001 {
		return 1.0
	}
	return math.Mod(n, max) / max
}

func clampFloat(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (c *Color) hexToLCH() [3]float64 {
	r, g, b := c.hexToSRGB()

	r = srgbToLinear(r)
	g = srgbToLinear(g)
	b = srgbToLinear(b)

	x, y, z := linearRGBToXYZ(r, g, b)
	l, a, labB := xyzToLab(x, y, z)
	lch_l, lch_c, lch_h := labToLCH(l, a, labB)

	return [3]float64{lch_l, lch_c, lch_h}
}

func (c *Color) lchToSRGB() (float64, float64, float64) {
	l := c.lch[0]
	ch := c.lch[1]
	h := c.lch[2]

	lab_l, lab_a, lab_b := lchToLab(l, ch, h)

	x, y, z := labToXYZ(lab_l, lab_a, lab_b)

	r, g, b := xyzToLinearRGB(x, y, z)

	r = linearToSRGB(r)
	g = linearToSRGB(g)
	b = linearToSRGB(b)

	return r, g, b
}

func (c *Color) lchToHex() string {
	r, g, b := c.lchToSRGB()

	rInt := int(math.Round(clampFloat(r, 0, 1) * 255))
	gInt := int(math.Round(clampFloat(g, 0, 1) * 255))
	bInt := int(math.Round(clampFloat(b, 0, 1) * 255))

	return fmt.Sprintf("#%02x%02x%02x", rInt, gInt, bInt)
}

func srgbToLinear(val float64) float64 {
	if val <= 0.04045 {
		return val / 12.92
	}
	return math.Pow((val+0.055)/1.055, 2.4)
}

func linearToSRGB(val float64) float64 {
	if val <= 0.0031308 {
		return 12.92 * val
	}
	return 1.055*math.Pow(val, 1.0/2.4) - 0.055
}

func linearRGBToXYZ(r, g, b float64) (float64, float64, float64) {
	x := 0.4360747*r + 0.3850649*g + 0.1430804*b
	y := 0.2225045*r + 0.7168786*g + 0.0606169*b
	z := 0.0139322*r + 0.0971045*g + 0.7141733*b

	return x, y, z
}

func xyzToLinearRGB(x, y, z float64) (float64, float64, float64) {
	r := 3.1338561*x - 1.6168667*y - 0.4906146*z
	g := -0.9787684*x + 1.9161415*y + 0.0334540*z
	b := 0.0719453*x - 0.2289914*y + 1.4052427*z

	return r, g, b
}

func xyzToLab(x, y, z float64) (float64, float64, float64) {
	const xn = 0.9642956
	const yn = 1.0
	const zn = 0.8251046

	const epsilon = 216.0 / 24389.0
	const kappa = 24389.0 / 27.0

	fx := x / xn
	fy := y / yn
	fz := z / zn

	if fx > epsilon {
		fx = math.Cbrt(fx)
	} else {
		fx = (kappa*fx + 16) / 116
	}

	if fy > epsilon {
		fy = math.Cbrt(fy)
	} else {
		fy = (kappa*fy + 16) / 116
	}

	if fz > epsilon {
		fz = math.Cbrt(fz)
	} else {
		fz = (kappa*fz + 16) / 116
	}

	l := 116*fy - 16
	a := 500 * (fx - fy)
	b := 200 * (fy - fz)

	return l, a, b
}

func labToXYZ(l, a, b float64) (float64, float64, float64) {
	const xn = 0.9642956
	const yn = 1.0
	const zn = 0.8251046

	const epsilon = 216.0 / 24389.0
	const kappa = 24389.0 / 27.0
	const epsilon3 = 24.0 / 116.0

	fy := (l + 16) / 116
	fx := a/500 + fy
	fz := fy - b/200

	var x, y, z float64

	if fx*fx*fx > epsilon {
		x = fx * fx * fx
	} else {
		x = (116*fx - 16) / kappa
	}

	if l > 8 {
		y = math.Pow((l+16)/116, 3)
	} else {
		y = l / kappa
	}

	if fz*fz*fz > epsilon {
		z = fz * fz * fz
	} else {
		z = (116*fz - 16) / kappa
	}

	return x * xn, y * yn, z * zn
}

func labToLCH(l, a, b float64) (float64, float64, float64) {
	c := math.Sqrt(a*a + b*b)
	h := math.Atan2(b, a) * 180 / math.Pi

	if h < 0 {
		h += 360
	}

	return l, c, h
}

func lchToLab(l, c, h float64) (float64, float64, float64) {
	hRad := h * math.Pi / 180
	a := c * math.Cos(hRad)
	b := c * math.Sin(hRad)

	return l, a, b
}

package gotransform

import "go/types"

type ValueOption struct {
	i    int
	d    float64
	kind types.BasicKind
}

func NewValueOptionInt(v int) *ValueOption {
	return &ValueOption{
		i:    v,
		d:    0,
		kind: types.Int,
	}
}

func NewValueOptionDouble(v float64) *ValueOption {
	return &ValueOption{
		i:    0,
		d:    v,
		kind: types.Float64,
	}
}

func (t *ValueOption) IsInt() bool {
	return t.kind == types.Int
}

func (t *ValueOption) IsDouble() bool {
	return t.kind == types.Float64
}

func (t *ValueOption) Int() int {
	return t.i
}

func (t *ValueOption) Double() float64 {
	return t.d
}

type GaussianBlur struct {
	Sigma        float64
	MinAmplitude float64
}

type AffineTransform struct {
	A            float64
	B            float64
	C            float64
	D            float64
	OffsetX      float64
	OffsetY      float64
	OffsetWidth  float64
	OffsetHeight float64
}

type Sharpen struct {
	Radius             float64
	FlatJaggyThreshold float64
	Brightening        float64
	Darkening          float64
	FlatSlope          float64
	JaggySlope         float64
}

type RGBA struct {
	R byte
	G byte
	B byte
	A byte
}

func NewRGB(red, green, blue int) RGBA {
	return NewRGBA(red, green, blue, 255)
}

func NewRGBA(red, green, blue, alpha int) RGBA {
	return RGBA{R: byte(red), G: byte(green), B: byte(blue), A: byte(alpha)}
}

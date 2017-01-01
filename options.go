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

type Options struct {
	CenterSampling bool
	Gamma          float64
	GaussianBlur   GaussianBlur
	Gravity        Gravity
	Interpolator   Interpolator
	Kernel         Kernel
	ResizeStrategy ResizeStrategy
	Sharpen        Sharpen
	Width          *ValueOption
	Height         *ValueOption
	CropOffsetX    *ValueOption
	CropOffsetY    *ValueOption
}

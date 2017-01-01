package gotransform

type ResizeStrategy int

const (
	ResizeStrategyDefault ResizeStrategy = iota
	ResizeStrategyCrop
	ResizeStrategyMax
	ResizeStrategyMin
	ResizeStrategyPad
	ResizeStrategyStretch
)

type Gravity int

const (
	GravityCenter Gravity = iota
	GravityN
	GravityE
	GravityS
	GravityW
	GravithNE
	GravitySE
	GravitySW
	GravityNW
)

type Kernel int

const (
	KernelDefault Kernel = iota
	KernelCubic
	KernelLanczos2
	KernelLanczos3
)

type Interpolator int

const (
	InterpolatorDefault Interpolator = iota
	InterpolatorBicubic
	InterpolatorBilinear
	InterpolatorLBB
	InterpolatorNearestNeighbor
	InterpolatorNohalo
	InterpolatorVSQBS
)

type CropStrategy int

const (
	CropStrategyEntropy CropStrategy = iota
	CropStrategyAttention
)

type GaussianBlur struct {
	Sigma float64
}

type Sharpen struct {
	Sigma float64
}

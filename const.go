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

type Flip int

const (
	FlipNone Flip = iota
	FlipX
	FlipY
	FlipXY
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

type ImageFormat int

const (
	ImageFormatUnknown ImageFormat = iota
	ImageFormatBmp
	ImageFormatJpeg
	ImageFormatPng
	ImageFormatWebp
)

type HistogramOperation int

const (
	HistogramOperationNone HistogramOperation = iota
	HistogramOperationCumulative
	HistogramOperationNormalize
)

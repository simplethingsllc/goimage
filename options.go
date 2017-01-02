package gotransform

type Options struct {
	BackgroundColor    *RGBA
	Blur               int
	CenterSampling     bool
	CropOffsetX        *ValueOption
	CropOffsetY        *ValueOption
	Flip               Flip
	Gamma              float64
	GaussianBlur       *GaussianBlur
	Gravity            Gravity
	Height             *ValueOption
	HistogramOperation HistogramOperation
	Interpolator       Interpolator
	Kernel             Kernel
	OutputFormat       ImageFormat
	PostFlip           int
	PostRotate         int
	Quality            int
	ResizeStrategy     ResizeStrategy
	Rotate             int
	Sharpen            *Sharpen
	Transform          *AffineTransform
	Width              *ValueOption
}

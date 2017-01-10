package gotransform

// Options specifies how to process the input image
type Options struct {
	BackgroundColor    *RGBA
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

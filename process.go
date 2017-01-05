package gotransform

import (
	"math"

	moreMath "github.com/davidbyttow/gomore/math"
	"github.com/davidbyttow/govips"
)

func Process(buf []byte, options *Options) ([]byte, error) {
	defer govips.ShutdownThread()

	image, err := govips.NewImageFromBuffer(buf, nil)
	if err != nil {
		return nil, err
	}

	image, err = makeImage(buf, image, options)
	if err != nil {
		return nil, err
	}

	return image.WriteToBuffer(".jpeg", nil)
}

func makeImage(sourceBytes []byte, image *govips.Image, options *Options) (*govips.Image, error) {

	// TODO(d): Rotation

	imageWidth := float64(image.Width())
	imageHeight := float64(image.Height())

	xFactor := 1.0
	yFactor := 1.0
	desiredWidth := getOptionValue(imageWidth, options.Width)
	desiredHeight := getOptionValue(imageHeight, options.Height)
	if desiredWidth > 0 && desiredHeight > 0 {
		xFactor = imageWidth / desiredWidth
		yFactor = imageHeight / desiredHeight
		switch options.ResizeStrategy {
		case ResizeStrategyCrop:
			fallthrough
		case ResizeStrategyPad:
			crop := options.ResizeStrategy == ResizeStrategyCrop
			if (crop && xFactor < yFactor) || (!crop && xFactor > yFactor) {
				desiredHeight = moreMath.Round(imageHeight / xFactor)
				yFactor = xFactor
			} else {
				desiredWidth = moreMath.Round(imageWidth / yFactor)
				xFactor = yFactor
			}
		case ResizeStrategyMax:
			fallthrough
		case ResizeStrategyMin:
			max := options.ResizeStrategy == ResizeStrategyMax
			if (max && xFactor > yFactor) || (!max && xFactor < yFactor) {
				desiredHeight = moreMath.Round(imageHeight / xFactor)
				options.Height = NewValueOptionInt(int(desiredHeight))
				yFactor = xFactor
			} else {
				desiredWidth = moreMath.Round(imageWidth / yFactor)
				options.Width = NewValueOptionInt(int(desiredWidth))
				xFactor = yFactor
			}
		case ResizeStrategyStretch:
			// Nothing to do unless there's a rotation
		}
	} else if desiredWidth > 0 {
		xFactor = imageWidth / desiredWidth
		if options.ResizeStrategy == ResizeStrategyStretch {
			desiredHeight = imageHeight
			options.Height = NewValueOptionInt(int(desiredHeight))
		} else {
			yFactor = xFactor
			desiredHeight = moreMath.Round(imageHeight / yFactor)
			options.Height = NewValueOptionInt(int(desiredHeight))
		}
	} else if desiredHeight > 0 {
		yFactor = imageHeight / desiredHeight
		if options.ResizeStrategy == ResizeStrategyStretch {
			desiredWidth = imageWidth
			options.Width = NewValueOptionInt(int(desiredWidth))
		} else {
			xFactor = yFactor
			desiredWidth = moreMath.Round(imageWidth / xFactor)
			options.Width = NewValueOptionInt(int(desiredWidth))
		}
	} else {
		options.Width = NewValueOptionInt(int(imageWidth))
		options.Height = NewValueOptionInt(int(imageHeight))
	}

	xShrink := int(math.Max(1.0, math.Floor(xFactor)))
	yShrink := int(math.Max(1.0, math.Floor(yFactor)))

	xResidual := float64(xShrink) / xFactor
	yResidual := float64(yShrink) / yFactor

	// Optionally prevent enlargement

	// imageType := image.Type()
	// canShrinkOnLoad := (imageType == govips.ImageTypeJpeg || imageType == govips.ImageTypeWebp) &&
	// 	!hasGammaAdjustment

	// shrinkFactor := 1
	// if canShrinkOnLoad && xShrink == yShrink && xShrink >= 2 {
	// 	if xShrink >= 8 {
	// 		shrinkFactor = 8
	// 	} else if xShrink >= 4 {
	// 		shrinkFactor = 4
	// 	} else if xShrink >= 2 {
	// 		shrinkFactor = 2
	// 	}
	// 	if shrinkFactor > 1 {
	// 		xFactor /= float64(shrinkFactor)
	// 		yFactor = xFactor
	// 	}
	// }

	// Reload the image with a shrink factor on load
	// if shrinkFactor > 1 {
	// 	var err error
	// 	if imageType == vips.ImageTypeJpeg {
	// 		image, err = vips.LoadJpegBuffer(sourceBytes, shrinkFactor)
	// 	} else {
	// 		image, err = vips.LoadWebpBuffer(sourceBytes, shrinkFactor)
	// 	}
	// 	if err != nil {
	// 		return image, err
	// 	}
	// 	shrunkWidth := image.Width()
	// 	shrunkHeight := image.Height()
	// 	xFactor = float64(shrunkWidth) / desiredWidth
	// 	yFactor = float64(shrunkHeight) / desiredHeight
	// 	xShrink = int(math.Max(1.0, math.Floor(xFactor)))
	// 	yShrink = int(math.Max(1.0, math.Floor(yFactor)))
	// 	xResidual = float64(xShrink) / xFactor
	// 	yResidual = float64(yShrink) / yFactor
	// }

	// TODO(d): Remove alpha channel?

	// TODO(d): Negate image if needed

	// hasGammaAdjustment := options.Gamma > 0
	// TODO(d): Gamma darkening

	// TODO(d): Greyscale

	// TODO(d): Overlay setup

	shrink := xShrink > 1 || yShrink > 1
	// reduce := xResidual != 1.0 || yResidual != 1.0
	// blur := options.GaussianBlur.Sigma != 0.0
	// sharpen := options.Sharpen.Sigma != 0.0

	// TODO(d): Premultiply alpha if needed

	if shrink {
		if yShrink > 1 {
			image = image.Shrinkv(yShrink)
		}
		if xShrink > 1 {
			image = image.Shrinkh(xShrink)
		}
		shrunkWidth := image.Width()
		shrunkHeight := image.Height()
		xResidual = desiredWidth / float64(shrunkWidth)
		yResidual = desiredHeight / float64(shrunkHeight)
	}

	if xResidual != 1.0 || yResidual != 1.0 {
		if xResidual < 1.0 || yResidual < 1.0 {
			// TODO(d): Pass in kernel
			if yResidual < 1.0 {
				image = image.Reducev(1.0 / yResidual)
			}
			if xResidual < 1.0 {
				image = image.Reduceh(1.0 / xResidual)
			}
		}
	}

	return image, nil
}

func getOptionValue(source float64, option *ValueOption) float64 {
	out := source
	if option != nil {
		if option.IsInt() {
			out = float64(option.Int())
		} else {
			out = out * option.Double()
		}
	}
	return out
}

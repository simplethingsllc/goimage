package gotransform

import (
	"io/ioutil"
	"testing"

	"github.com/davidbyttow/govips"
	"github.com/stretchr/testify/require"
)

func TestResize(t *testing.T) {
	govips.Startup(nil)
	defer govips.Shutdown()

	originalBuf, err := ioutil.ReadFile("testdata/canyon.jpg")
	require.Nil(t, err)

	options := &Options{
		Width: NewValueOptionDouble(0.25),
	}

	buf, err := Process(originalBuf, options)
	require.Nil(t, err)

	resized, err := govips.NewImageFromBuffer(buf, nil)
	require.Nil(t, err)
	require.Equal(t, 640, resized.Width())
	require.Equal(t, 400, resized.Height())
	GoldenCompare(t, buf, "testdata/resize-relative-golden.jpg")

	options = &Options{
		Width:          NewValueOptionDouble(0.1),
		ResizeStrategy: ResizeStrategyStretch,
	}

	buf, err = Process(originalBuf, options)
	require.Nil(t, err)

	scaled, err := govips.NewImageFromBuffer(buf, nil)
	require.Nil(t, err)
	require.Equal(t, 256, scaled.Width())
	require.Equal(t, 1600, scaled.Height())
	GoldenCompare(t, buf, "testdata/resize-scaled-golden.jpg")

	options = &Options{
		Width:  NewValueOptionInt(400),
		Height: NewValueOptionInt(400),
	}

	buf, err = Process(originalBuf, options)
	require.Nil(t, err)

	cropped, err := govips.NewImageFromBuffer(buf, nil)
	require.Nil(t, err)
	require.Equal(t, 400, cropped.Width())
	require.Equal(t, 250, cropped.Height())
	GoldenCompare(t, buf, "testdata/resize-crop-golden.jpg")
}

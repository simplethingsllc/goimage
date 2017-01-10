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

	buf, err := ioutil.ReadFile("testdata/canyon.jpg")
	require.Nil(t, err)

	options := &Options{
		Width: NewValueOptionDouble(0.25),
	}

	buf, err = Process(buf, options)
	require.Nil(t, err)

	resized, err := govips.NewImageFromBuffer(buf, nil)
	require.Nil(t, err)

	require.Equal(t, 640, resized.Width())
	require.Equal(t, 400, resized.Height())

	GoldenCompare(t, buf, "testdata/resize-golden.jpg")
}

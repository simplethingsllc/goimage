package gotransform

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResize(t *testing.T) {
	buf, err := ioutil.ReadFile("fixtures/canyon.jpg")
	require.Nil(t, err)

	options := &Options{
		Width: NewValueOptionDouble(0.25),
	}

	buf, err = Process(buf, options)
	require.Nil(t, err)

	// TODO(d): Figure out why this image is failing to load
	// resized, err := govips.NewImageFromBuffer(buf, nil)
	// require.Nil(t, err)

	// assert.Equal(t, 320, resized.Width())
}

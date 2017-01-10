package gotransform

import (
	go_debug "github.com/tj/go-debug"
)

var debugFn = go_debug.Debug("gotransform")

func debug(fmt string, values ...interface{}) {
	if len(values) > 0 {
		debugFn(fmt, values...)
	} else {
		debugFn(fmt)
	}
}

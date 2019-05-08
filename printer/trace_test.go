package printer

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("", t, func() {
		content := []byte(`
HEADER-----------
stack traceback:
	stdin:1	 in main chunk
	[C]: in ?
IGNORE-----------`)
		trace(bytes.NewReader(content), "")
	})
	Convey("", t, func() {
		content := []byte(`
HEADER-----------
stack traceback:
	stdin:1	 in main chunk
	[C]: in ?
stack traceback:
	stdin:1	 in main chunk
	[C]: in ?
IGNORE-----------`)
		trace(bytes.NewReader(content), "")
	})
}

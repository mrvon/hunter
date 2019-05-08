package printer

import (
	"bytes"

	"github.com/gookit/color"
	"github.com/mrvon/hunter/finder"
)

func Traceback(serverName string) {
	buf, err := finder.Fetch(serverName)
	if err != nil {
		color.Error.Println(err)
		return
	}
	color.Success.Printf("TRACEBACK | %s | OK\n", serverName)
	trace(bytes.NewReader(buf), "")
}

package printer

import (
	"bufio"
	"io"
	"strings"

	"github.com/gookit/color"
)

func trace(reader io.Reader, source string) {
	scanner := bufio.NewScanner(reader)
	prev := ""
	trace := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "stack traceback:" || strings.HasPrefix(line, "\t") {
			if !trace {
				trace = true
				if len(source) > 0 {
					color.Comment.Printf(source)
				}
				color.Danger.Println(prev)
			}
			color.Note.Println(line)
		} else {
			trace = false
		}
		prev = line
	}
	if err := scanner.Err(); err != nil {
		color.Error.Println(err)
	}
}

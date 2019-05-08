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
	block := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "stack traceback:" {
			if !block {
				block = true
				if len(source) > 0 {
					color.Comment.Printf(source)
				}
				color.Danger.Println(prev)
			}
			color.Note.Println(line)
		} else if strings.HasPrefix(line, "\t") {
			if block {
				color.Note.Println(line)
			}
		} else {
			block = false
		}
		prev = line
	}
	if err := scanner.Err(); err != nil {
		color.Error.Println(err)
	}
}

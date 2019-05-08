package printer

import (
	"bufio"
	"strings"

	"github.com/gookit/color"
	"github.com/mrvon/hunter/config"
	"github.com/mrvon/hunter/finder"
)

func tracecheck(serverName string) {
	color.Note.Printf("CHECK | %s | start\n", serverName)
	reportOK := func() {
		color.Success.Printf("CHECK | %s | OK\n", serverName)
	}
	reportFailed := func() {
		color.Error.Printf("CHECK | %s | failed\n", serverName)
	}
	rc, err := finder.TailCheck(serverName)
	if err != nil {
		reportFailed()
		return
	}
	defer rc.Close()
	scanner := bufio.NewScanner(rc)
	if !scanner.Scan() {
		reportFailed()
		return
	}
	line := scanner.Text()
	if strings.HasPrefix(line, "==>") && strings.HasSuffix(line, "<==") {
		reportOK()
	} else {
		reportFailed()
	}
}

func Tracecheck(serverName string) {
	c, err := config.Get()
	if err != nil {
		return
	}
	if serverName == "" {
		for name := range c.Server {
			tracecheck(name)
		}
	} else {
		tracecheck(serverName)
	}
}

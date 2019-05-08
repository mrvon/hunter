package printer

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/mrvon/hunter/config"
	"github.com/mrvon/hunter/finder"
)

func tracetail(serverName string) {
	rc, err := finder.Tail(serverName)
	if err != nil {
		color.Error.Println(err)
		return
	}
	color.Success.Printf("TRACETAIL | %s\n", serverName)
	go func() {
		source := fmt.Sprintf("FROM | %s\n", serverName)
		trace(rc, source)
		defer rc.Close()
	}()
}

func Tracetail(serverName string) {
	c, err := config.Get()
	if err != nil {
		return
	}
	if serverName == "" {
		for name := range c.Server {
			tracetail(name)
		}
	} else {
		tracetail(serverName)
	}
}

package finder

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/mrvon/hunter/config"
)

func tail(serverName string, arg string) (rc io.ReadCloser, err error) {
	var cmd *exec.Cmd
	c, err := config.Get()
	if err != nil {
		err = fmt.Errorf("read config failed | %v", err)
		return
	}
	item, existed := c.Server[serverName]
	if !existed {
		err = fmt.Errorf("config not found | server: %s", serverName)
		return
	}
	if item.Type == "local" {
		cmd = exec.Command("tail", arg, item.Path)
	} else {
		cmd = exec.Command("ssh", item.SSH, "tail", arg, item.Path)
	}
	rc, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	go func() {
		cmd.Wait()
	}()
	return
}

func Tail(serverName string) (rc io.ReadCloser, err error) {
	rc, err = tail(serverName, "-F")
	return
}

func TailCheck(serverName string) (rc io.ReadCloser, err error) {
	rc, err = tail(serverName, "-v")
	return
}

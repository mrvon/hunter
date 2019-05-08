package finder

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/mrvon/hunter/config"
)

var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextRandom() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func randomFileName() string {
	return fmt.Sprintf("/tmp/hunter_%s", nextRandom())
}

func Fetch(serverName string) (buf []byte, err error) {
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
	var cmd *exec.Cmd
	dstPath := randomFileName()
	if item.Type == "local" {
		srcPath := item.Path
		cmd = exec.Command("cp", srcPath, dstPath)
	} else {
		srcPath := fmt.Sprintf("%s:%s", item.SSH, item.Path)
		cmd = exec.Command("scp", srcPath, dstPath)
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	err = cmd.Wait()
	if err != nil {
		return
	}
	defer func() {
		os.Remove(dstPath)
	}()
	buf, err = ioutil.ReadFile(dstPath)
	if err != nil {
		return
	}
	return
}

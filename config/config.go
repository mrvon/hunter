package config

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

var configList = []string{
	"~/.hunter.json",
	"hunter.json",
}

type ConfigItem struct {
	Type string
	SSH  string
	Path string
}

type Config struct {
	Server map[string]ConfigItem
}

func expandTilde(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if path == "~" {
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(dir, path[2:])
	}
	return path
}

func Get() (*Config, error) {
	var buf []byte
	var err error
	for _, path := range configList {
		path = expandTilde(path)
		buf, err = ioutil.ReadFile(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}
	for name, item := range c.Server {
		if item.Type == "local" {
			item.Path = expandTilde(item.Path)
			c.Server[name] = item
		}
	}
	return &c, nil
}

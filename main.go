package main

import (
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/gookit/color"
	"github.com/mrvon/hunter/config"
	"github.com/mrvon/hunter/printer"
)

const (
	traceback  = "traceback"
	tracetail  = "tracetail"
	tracecheck = "check"
	quit       = "quit"
)

func listServers() func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		c, err := config.Get()
		if err != nil {
			return names
		}
		for name := range c.Server {
			names = append(names, name)
		}
		return names
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem(traceback,
		readline.PcItemDynamic(listServers()),
	),
	readline.PcItem(tracetail,
		readline.PcItemDynamic(listServers()),
	),
	readline.PcItem(tracecheck,
		readline.PcItemDynamic(listServers()),
	),
	readline.PcItem(quit),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:              color.Note.Sprint("Hunter> "),
		HistoryFile:         "/tmp/readline.tmp",
		AutoComplete:        completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	log.SetOutput(l.Stderr())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	outerloop:
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, traceback):
			serverName := strings.TrimSpace(line[len(traceback):])
			printer.Traceback(serverName)
		case strings.HasPrefix(line, tracetail):
			serverName := strings.TrimSpace(line[len(tracetail):])
			printer.Tracetail(serverName)
		case strings.HasPrefix(line, tracecheck):
			serverName := strings.TrimSpace(line[len(tracecheck):])
			printer.Tracecheck(serverName)
		case strings.HasPrefix(line, quit[:1]):
			break outerloop;
		}
	}
}

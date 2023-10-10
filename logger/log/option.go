package log

import (
	"log"
	"os"
)

type Option func(g *logger)

func SyncConsole() Option {
	return func(g *logger) {
		g.Logger = log.New(os.Stderr, "", log.Llongfile|log.LstdFlags)
	}
}

func SetCallDepth(depth int) Option {
	return func(g *logger) {
		g.callDepth = depth
	}
}

func SyncFile(file string) Option {
	return func(g *logger) {
		openFile, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if err != nil {
			panic(err)
		}
		g.Logger = log.New(openFile, "", log.Lshortfile|log.LstdFlags)
	}
}

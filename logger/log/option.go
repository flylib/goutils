package log

import (
	"log"
	"os"
)

type Option func(g *logger)

type option struct {
	syncFile    string
	syncConsole bool
	//syncEmail   string //同步邮箱
	//syncHttp    string //同步http
	timeFormat    string
	maxFileSize   int //日志文件最大多大
	maxAge        int //文件最多保存多少天
	outJsonStyle  bool
	minPrintLevel Level
}

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
		g.Logger.SetOutput()
	}
}

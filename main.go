package main

import (
	"log"
	"os"

	"go-completion/vim"

	"github.com/neovim/go-client/nvim/plugin"
)

// 1. create /tmp/go-completion folder if err
// 2. init all source logger
func initialPlugin() {
	dispatch = &Dispatcher{}
}

var logger *log.Logger

var dispatch *Dispatcher

type LuaSource struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func main() {
	// define log output place
	f, err := os.Open("/tmp/go-completion.log")
	if err != nil {
		f, _ = os.Create("/tmp/go-completion.log")
	}
	// log.SetOutput(f)
	logger = log.New(f, "[go-completion]", log.Ltime|log.Lshortfile)

	// defer f.Close()

	plugin.Main(func(pm *plugin.Plugin) error {
		// logger when start
		logger.Println("go-completion have start")

		// store global vim pointer
		vim.SetVim(pm.Nvim)

		initialPlugin()
		logger.Println("go-completion have done all initial work ")

		//
		// register global function
		//
		// send source
		pm.HandleFunction(&plugin.FunctionOptions{Name: "GoCompletionSend"}, func(arg any) error {
			logger.Println("get source")
			logger.Print(arg)
			return nil
		})

		pm.HandleFunction(&plugin.FunctionOptions{Name: "GoStartFn"}, func(arg any) (string, error) {
			logger.Print("hello")
			f.Close()
			return "go start", nil
		})

		return nil
	})
}

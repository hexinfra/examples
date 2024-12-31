package main

import (
	"os"

	"github.com/hexinfra/gorox/hemi/procmgr"
	"myprg/test"

	_ "myprg/apps"
	_ "myprg/exts"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "test" {
		test.Main()
	} else {
		procmgr.Main(&procmgr.Opts{
			ProgramName:  "myprg",
			ProgramTitle: "MyProgram",
			DebugLevel:   0,
			CmdUIAddr:    "127.0.0.1:9527",
			WebUIAddr:    "127.0.0.1:9528",
		})
	}
}

package main

import (
	"phoenix/cmd"
	"phoenix/lib"
)

var (
	version     string
	dotconfname string
)

func main() {
	lib.Initialize(lib.Variables{Dotconfname: dotconfname})
	cmd.Initialize(cmd.Variables{Version: version})
	cmd.Execute()
}

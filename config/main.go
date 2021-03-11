package main

import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/rwxrob/cmdtab-config"
)

func main() {
	cmdtab.Execute("config")
}

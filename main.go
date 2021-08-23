package main

import (
	"github.com/sibprogrammer/uws/cmd"
)

var (
	commit  string
	date    string
	version string
)

func init() {
	cmd.Commit = commit
	cmd.BuildTime = date
	cmd.Version = version
}

func main() {
	cmd.Execute()
}

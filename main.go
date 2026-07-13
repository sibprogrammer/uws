package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/sibprogrammer/uws/cmd"
)

var (
	commit = "000000"
	date   = ""
)

//go:embed version
var version string

func init() {
	fullVersion := strings.TrimSpace(version)
	if date != "" {
		fullVersion += fmt.Sprintf(" (%s, %s)", date, commit)
	}
	cmd.Version = fullVersion
}

func main() {
	cmd.Execute()
}

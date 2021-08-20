package main

import (
	_ "embed"
	"fmt"
	"github.com/asecurityteam/env-mapper/pkg/mapper"
	"os"
)

//go:embed usage.txt
var usage string
func main() {
	cmd, err := mapper.CommandWithEnvOverrides(os.Args[1:], os.Environ())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n\n%s", err.Error(), usage)
		os.Exit(1)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

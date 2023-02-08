package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/asecurityteam/env-mapper/pkg/mapper"
)

//go:embed usage.txt
var usage string

func main() {
	var envSep string
	flag.StringVar(&envSep, "envSep", ":", "a string used to separate target from source for mapping. defaults to ':'")

	var complexVar bool
	flag.BoolVar(&complexVar, "complex", false, "a boolean to show what mode to interpret source variables in. "+
		"true allows for combining multiple target variables but requires  '||VALUE||' as a delimiter. ex: "+
		"TARGET:cat||SOURCE1||:||SOURCE2|| where SOURCE1 is foo and SOURCE2 is bar becomes 'catfoo:bar'. defaults to false")

	flag.Parse()

	conf := mapper.Config{
		EnvSeparator: envSep,
		ComplexVar:   complexVar,
	}
	cmd, err := mapper.CommandWithEnvOverrides(conf, flag.Args(), os.Environ()) // flag.Args() gives positional arguments left after parsing defined flags
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

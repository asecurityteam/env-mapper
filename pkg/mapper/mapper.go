package mapper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type envMapping struct {
	To   string
	From string
}

// Config is used to configure the complex mapper via flags parsed in main.go
type Config struct {
	EnvSeparator string
	ComplexVar   bool
}

func parseMappings(src []string, sep string) ([]envMapping, error) {
	mapping := make([]envMapping, 0, len(src))
	targets := make(map[string]bool, len(src))
	for _, s := range src {
		kv := strings.SplitN(s, sep, 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("unable to parse mapping: '%v' using separator '%v'", s, sep)
		}
		target, source := kv[0], kv[1]
		if _, ok := targets[target]; ok {
			return nil, fmt.Errorf("duplicate target for mapping: '%v' in mapping '%v'", target, s)
		}
		targets[target] = true
		mapping = append(mapping, envMapping{
			To:   target,
			From: source,
		})
	}
	return mapping, nil
}

func resolveMappings(mappings []envMapping, resolver func(string) string) []string {
	resolved := make([]string, 0, len(mappings))
	for _, mapping := range mappings {
		resolved = append(resolved, fmt.Sprintf("%s=%s", mapping.To, resolver(mapping.From)))
	}
	return resolved
}

// complexResolver will always attempt an os.GetEnv lookup for anything being substituted
func complexResolver(unsubbed string) string {
	delim := "||"
	//Check that we have balanced delimiters before attempting
	if strings.Count(unsubbed, delim)%2 == 0 {
		subbed := unsubbed
		for strings.Contains(subbed, delim) {
			//Get whatever was before the substitution needed
			pre, left, _ := strings.Cut(subbed, delim)
			//Get the environment variable to sub and whatever was left
			sub, after, _ := strings.Cut(left, delim)
			lookup := os.Getenv(sub)

			//This will sub in whatever the lookup got, and remove the placeholders thanks to cut
			subbed = pre + lookup + after
		}
		return subbed
	}
	return unsubbed
}

func bisectSlice(src []string, sep string) ([]string, []string) {
	for pos := range src {
		if src[pos] == sep {
			return src[:pos], src[pos+1:]
		}
	}
	return nil, nil
}

// CommandWithEnvOverrides creates exec.Cmd with new environment
// based on mappings passed in inputArgs and existing values in inputEnv.
//
// The values from inputEnv (usually return value from os.Environ()) can be extended and overridden
// based on the TARGET:SOURCE pairs provided in the inputArgs.
//
// Environment variable TARGET in the resulting environment in exec.Cmd is
// set to the value of SOURCE in input environment or empty string if the value is not set or is empty.
//
// For example, calling:
//
//	CommandWithEnvOverrides(
//		[]string{"A:PATH","B:UNKNOWN_VARIABLE","--","/bin/sh"},
//		[]string{"PATH=/bin"}
//		)
//
// Will result in exec.Cmd with path "/bin/sh" and with environment
//
//	{"A=/bin","B=", "PATH=/bin"}
//
// Where A is assigned to whatever value PATH is set to, and B is set to empty value as the UNKNOWN_VARIABLE is not
// defined in inputEnv.
//
// Separator:
//
//	"--"
//
// is used to divide "TARGET:SOURCE" mappings and the path to the command with (optional) arguments.
func CommandWithEnvOverrides(conf Config, inputArgs []string, inputEnv []string) (*exec.Cmd, error) {

	mapNames, cmdLine := bisectSlice(inputArgs, "--")
	if len(cmdLine) < 1 || len(cmdLine[0]) < 1 {
		return nil, fmt.Errorf("missing command path")
	}
	cmdPath := cmdLine[0]
	cmdArgs := cmdLine[1:]

	mappings, err := parseMappings(mapNames, conf.EnvSeparator)
	if err != nil {
		return nil, err
	}

	resolver := os.Getenv
	if conf.ComplexVar {
		resolver = complexResolver
	}

	envOverrides := resolveMappings(mappings, resolver)
	command := exec.Command(cmdPath, cmdArgs...)
	command.Env = append(inputEnv, envOverrides...) //envOverrides take precedence
	return command, nil
}

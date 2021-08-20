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
// 		CommandWithEnvOverrides(
//			[]string{"A:PATH","B:UNKNOWN_VARIABLE","--","/bin/sh"},
//			[]string{"PATH=/bin"}
//			)
// Will result in exec.Cmd with path "/bin/sh" and with environment
// 		{"A=/bin","B=", "PATH=/bin"}
// Where A is assigned to whatever value PATH is set to, and B is set to empty value as the UNKNOWN_VARIABLE is not
// defined in inputEnv.
//
// Separator:
//		"--"
// is used to divide "TARGET:SOURCE" mappings and the path to the command with (optional) arguments.
func CommandWithEnvOverrides(inputArgs []string, inputEnv []string) (*exec.Cmd, error) {
	mapNames, cmdLine := bisectSlice(inputArgs, "--")
	if len(cmdLine) < 1 || len(cmdLine[0]) < 1 {
		return nil, fmt.Errorf("missing command path")
	}
	cmdPath := cmdLine[0]
	cmdArgs := cmdLine[1:]

	mappings, err := parseMappings(mapNames, ":")
	if err != nil {
		return nil, err
	}
	envOverrides := resolveMappings(mappings, os.Getenv)
	command := exec.Command(cmdPath, cmdArgs...)
	command.Env = append(inputEnv, envOverrides...) //envOverrides take precedence
	return command, nil
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ReadDir(dir string) (map[string]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	envs := make(map[string]string)
	for _, file := range files {
		value, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		envs[file.Name()] = string(value)
	}

	return envs, nil
}

func RunCmd(cmdv []string, env map[string]string) int {
	cmd := exec.Command(cmdv[0], cmdv[1:]...)
	cmd.Env = os.Environ()
	for key, value := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		eerr, ok := err.(*exec.ExitError)
		if ok {
			return eerr.ExitCode()
		} else {
			fmt.Fprint(os.Stderr, err)
			return 1
		}
	}
	return 0
}

func main() {
	envDir := os.Args[1]
	cmdv := os.Args[2:]

	env, err := ReadDir(envDir)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	exitCode := RunCmd(cmdv, env)

	os.Exit(exitCode)
}

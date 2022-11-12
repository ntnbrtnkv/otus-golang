package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDirGetEnv(t *testing.T) {
	env, err := ReadDir("./examples")

	res := map[string]string{
		"MULTILINE": "qweqwe\r\nasdasd\r\nzxcxzc",
		"TEST_ENV":  "123",
		"YOTA":      "very nice",
	}

	require.Nil(t, err)
	require.Equal(t, res, env)
}

func TestReadDirNotExisting(t *testing.T) {
	_, err := ReadDir("./eamples")

	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestRunCmd(t *testing.T) {
	exitCode := RunCmd([]string{"powershell", "dir", "env:"}, make(map[string]string))

	require.Equal(t, 0, exitCode)
}

func TestRunCmdNotExisitinCommand(t *testing.T) {
	exitCode := RunCmd([]string{"kek123"}, make(map[string]string))

	require.Equal(t, 1, exitCode)
}

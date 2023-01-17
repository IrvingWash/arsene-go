package utils

import (
	"os/user"
	"path/filepath"
	"strings"
)

func AppendHomeDirIfNeeded(path string) (string, error) {
	if path[0:1] != "~" {
		return path, nil
	}

	user, err := user.Current()

	if err != nil {
		return "", err
	}

	homeDir := user.HomeDir

	if path == "~" {
		return homeDir, nil
	}

	if strings.HasPrefix(path, "~/") {
		return filepath.Join(homeDir, path[2:]), nil
	}

	return "", err
}

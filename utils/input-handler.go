package utils

import (
	"fmt"
	"log"
	"os"
)

func InputHandler() (url, downloadPath string) {
	args := os.Args[1:]

	if len(args) == 1 {
		fmt.Println("Usage: arsene bandcamp-url download-path")
		os.Exit(0)
	}

	if len(args) < 2 {
		log.Fatal("Too few arguments. Use 'arsene' to see help")
	}

	path, err := AppendHomeDirIfNeeded(args[1])

	if err != nil {
		log.Fatal("Failed to parse path: ", err)
	}

	return args[0], path
}

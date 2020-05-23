package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
)

const (
	DefaultFilename = ".dockerbuild"
	DefaultFileType = "json"
)

func DefaultConfigPath() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return home
}

func DefaultConfigFilePath() string {
	return DefaultConfigPath() + "/" + DefaultFilename + "." + DefaultFileType
}

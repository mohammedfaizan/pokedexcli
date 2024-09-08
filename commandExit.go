package main

import (
	"os"
)

func commandExit(cfg *Config) error {
	os.Exit(1)
	return nil
}

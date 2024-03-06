//go:build windows
// +build windows

package main

import (
	"os"

	"golang.org/x/sys/windows"
	"golang.org/x/term"
)

func readPassword() ([]byte, error) {
	return term.ReadPassword(int(windows.Handle(os.Stdin.Fd())))
}

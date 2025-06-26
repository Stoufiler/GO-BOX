package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

// GetPassword prompts the user for a password and returns it.
func GetPassword() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	fmt.Println()
	return password, nil
}
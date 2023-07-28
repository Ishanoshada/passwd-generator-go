package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	defaultPasswordLength = 12
)

var (
	
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers          = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~"
)

func generatePassword(length int, useLowercase, useUppercase, useNumbers, useSpecialChars bool) (string, error) {
	charset := ""

	if useLowercase {
		charset += lowercaseLetters
	}
	if useUppercase {
		charset += uppercaseLetters
	}
	if useNumbers {
		charset += numbers
	}
	if useSpecialChars {
		charset += specialChars
	}

	if charset == "" {
		return "", fmt.Errorf("you must select at least one character set")
	}

	password := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}

func main() {
	passwordLength := defaultPasswordLength

	// Customize these boolean values to include or exclude character sets
	useLowercase := true
	useUppercase := true
	useNumbers := true
	useSpecialChars := true

	password, err := generatePassword(passwordLength, useLowercase, useUppercase, useNumbers, useSpecialChars)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}

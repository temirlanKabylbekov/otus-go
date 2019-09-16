package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpack(s string) (string, error) {
	var result strings.Builder

	var letter string
	var pairs int8

	for _, rune := range s {
		pairs += 1

		if pairs == 1 {
			if unicode.IsLetter(rune) {
				letter = string(rune)
			} else if unicode.IsDigit(rune) {
				return "", fmt.Errorf("Passed string %q starts with digit or have two digits next to each other", s)
			} else {
				return "", fmt.Errorf("Passed string %q should contain only letters and digits", s)
			}
		}

		if pairs == 2 {
			if unicode.IsLetter(rune) {
				if _, err := result.WriteString(letter); err != nil {
					return "", err
				}
				letter = string(rune)
				pairs = 1
			} else if unicode.IsDigit(rune) {
				digit, err := strconv.Atoi(string(rune))
				if err != nil {
					return "", err
				}
				if _, err := result.WriteString(strings.Repeat(letter, digit)); err != nil {
					return "", err
				}
				letter = ""
				pairs = 0
			} else {
				return "", fmt.Errorf("Passed string %q should contain only letters and digits", s)
			}
		}
	}

	if letter != "" {
		if _, err := result.WriteString(letter); err != nil {
			return "", err
		}
	}

	return result.String(), nil
}

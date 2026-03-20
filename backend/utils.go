package main

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Encode a number to a Base62 string

func Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	var builder strings.Builder
	length := uint64(len(alphabet))

	for number > 0 {
		remainder := number % length
		builder.WriteByte(alphabet[remainder])
		number = number / length
	}

	return reverse(builder.String())
}

// Reverse a string

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1;i<j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
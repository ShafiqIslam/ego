package main

import (
	"fmt"
	"strings"
)

func encrypt(plaintext, keyword string) string {
	plaintext = strings.ToUpper(plaintext)
	keyword = strings.ToUpper(keyword)

	var ciphertext strings.Builder
	keyIdx := 0

	for _, char := range plaintext {
		if char < 'A' || char > 'Z' {
			continue
		}

		shift := keyword[keyIdx%len(keyword)] - 'A'
		fmt.Printf("Shift %c with %v\n", char, shift)
		encryptedChar := (char-'A'+rune(shift))%26 + 'A'

		ciphertext.WriteRune(encryptedChar)
		keyIdx++
	}

	return ciphertext.String()
}

func main() {
	plaintext := "LONGLIVEPOLYGON"
	keyword := "POYSA"
	ciphertext := encrypt(plaintext, keyword)

	fmt.Println("Plaintext:  ", strings.ToLower(plaintext))
	fmt.Println("Key:  ", strings.ToLower(keyword))
	fmt.Println("Ciphertext: ", strings.ToLower(ciphertext))
}

package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."

	// 1. Bytes y runas
	bytes := len(texto)
	runas := utf8.RuneCountInString(texto)

	// 2. Palabras
	palabras := len(strings.Fields(texto))

	// 3. Mayúsculas y dígitos
	mayusculas := 0
	digitos := 0

	for _, r := range texto {
		if unicode.IsUpper(r) {
			mayusculas++
		}
		if unicode.IsDigit(r) {
			digitos++
		}
	}

	// 4. Letras con tilde o ñ
	conTildeON := 0

	for _, r := range texto {
		if strings.ContainsRune("áéíóúÁÉÍÓÚñÑ", r) {
			conTildeON++
		}
	}

	fmt.Println("Bytes:", bytes)
	fmt.Println("Runas:", runas)
	fmt.Println("Palabras:", palabras)
	fmt.Println("Mayúsculas:", mayusculas)
	fmt.Println("Dígitos:", digitos)
	fmt.Println("Con tilde/ñ:", conTildeON)
}
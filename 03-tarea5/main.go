package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	LimiteSubtitulo = 42

	ErrUso    = 1
	ErrVacio  = 2
	ErrLimite = 3
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Uso: go run . \"texto\"")
		os.Exit(ErrUso)
	}

	texto := os.Args[1]

	if strings.TrimSpace(texto) == "" {
		fmt.Fprintln(os.Stderr, "Error: texto vacío")
		os.Exit(ErrVacio)
	}

	var letras int
	var digitos int
	var mayusculas int
	var runasMas1Byte int
	var palabras int

	enPalabra := false

	for _, r := range texto {

		if unicode.IsLetter(r) {
			letras++
		}

		if unicode.IsDigit(r) {
			digitos++
		}

		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			mayusculas++
		}

		if utf8.RuneLen(r) > 1 {
			runasMas1Byte++
		}

		if !unicode.IsSpace(r) {
			if !enPalabra {
				palabras++
				enPalabra = true
			}
		} else {
			enPalabra = false
		}
	}

	bytes := len(texto)
	runas := utf8.RuneCountInString(texto)

	fmt.Printf("texto: %q\n", texto)
	fmt.Printf("bytes: %d\n", bytes)
	fmt.Printf("runas: %d\n", runas)
	fmt.Printf("palabras: %d\n", palabras)
	fmt.Printf("letras: %d\n", letras)
	fmt.Printf("dígitos: %d\n", digitos)
	fmt.Printf("mayúsculas: %d\n", mayusculas)
	fmt.Printf("runas de más de 1 byte: %d\n", runasMas1Byte)

	if runas > LimiteSubtitulo {
		fmt.Fprintln(os.Stderr, "límite de 42 runas: fuera")
		os.Exit(ErrLimite)
	}

	fmt.Println("límite de 42 runas: dentro")
}

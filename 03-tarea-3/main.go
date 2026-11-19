package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ErrUso    = 1
	ErrValor  = 2
	ErrUnidad = 3
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Uso: go run . <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || valor < 0 {
		fmt.Fprintln(os.Stderr, "Valor inválido")
		os.Exit(ErrValor)
	}

	origen := strings.ToUpper(os.Args[2])
	destino := strings.ToUpper(os.Args[3])

	var bytesOrigen float64
	var bytesDestino float64

	switch origen {
	case "B":
		bytesOrigen = 1
	case "KB":
		bytesOrigen = 1024
	case "MB":
		bytesOrigen = 1024 * 1024
	case "GB":
		bytesOrigen = 1024 * 1024 * 1024
	default:
		fmt.Fprintln(os.Stderr, "Unidad inválida")
		os.Exit(ErrUnidad)
	}

	switch destino {
	case "B":
		bytesDestino = 1
	case "KB":
		bytesDestino = 1024
	case "MB":
		bytesDestino = 1024 * 1024
	case "GB":
		bytesDestino = 1024 * 1024 * 1024
	default:
		fmt.Fprintln(os.Stderr, "Unidad inválida")
		os.Exit(ErrUnidad)
	}

	resultado := valor * bytesOrigen / bytesDestino

	fmt.Printf("%g %s = %.2f %s\n",
		valor,
		origen,
		resultado,
		destino,
	)
}
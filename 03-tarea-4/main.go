package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"tarea4/internal/almacenamiento"
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

	bytesOrigen, ok := almacenamiento.Bytes(origen)
	if !ok {
		fmt.Fprintln(os.Stderr, "Unidad inválida")
		os.Exit(ErrUnidad)
	}

	bytesDestino, ok := almacenamiento.Bytes(destino)
	if !ok {
		fmt.Fprintln(os.Stderr, "Unidad inválida")
		os.Exit(ErrUnidad)
	}
	resultado := almacenamiento.Convertir(
		valor,
		bytesOrigen,
		bytesDestino,
	)

	fmt.Printf("%.2f %s = %.2f %s\n",
		valor,
		origen,
		resultado,
		destino,
	)
}

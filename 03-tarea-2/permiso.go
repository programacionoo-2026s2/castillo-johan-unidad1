package main

import "fmt"

type Permiso uint8

const (
	Leer     Permiso = 4
	Escribir Permiso = 2
	Ejecutar Permiso = 1
)

func describir(p Permiso) string {
	s := ""

	if p&Leer != 0 {
		s += "r"
	} else {
		s += "-"
	}

	if p&Escribir != 0 {
		s += "w"
	} else {
		s += "-"
	}

	if p&Ejecutar != 0 {
		s += "x"
	} else {
		s += "-"
	}

	return s
}

func main() {
	usuario := Leer | Escribir | Ejecutar
	grupo := Leer | Ejecutar
	otros := Leer

	fmt.Println(describir(usuario) + describir(grupo) + describir(otros))

	fmt.Printf("usuario: %d\n", usuario)
	fmt.Printf("grupo: %d\n", grupo)
	fmt.Printf("otros: %d\n", otros)

	numerico := int(usuario)*100 + int(grupo)*10 + int(otros)
	fmt.Printf("numérico: %d\n", numerico)

	otros |= Escribir
	grupo &^= Ejecutar

	fmt.Println(describir(usuario) + describir(grupo) + describir(otros))
}
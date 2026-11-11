package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Ñandú"
	fmt.Println("longitud:", len(s))                 // bytes
	fmt.Println("runas:", utf8.RuneCountInString(s)) // runas
	for i, r := range s {
		fmt.Printf("%d:%c ", i, r)
	}
	fmt.Println()
	fmt.Println(s[0]) // un byte, no una letra
}

package main

import (
	"fmt"

	"github.com/jsonandrx/gozero/variables"
)

func main() {
	estado, texto := variables.ConvierteATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
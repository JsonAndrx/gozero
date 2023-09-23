package main

import (
	"fmt"
	// "runtime"
	// "github.com/jsonandrx/gozero/variables"
	"github.com/jsonandrx/gozero/ejercicios"
)

func main() {
	// estado, texto := variables.ConvierteATexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("Esto es", os)
	// }
	// fmt.Println("Esto es windows")

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es linux")
	// case "darwin":
	// 	fmt.Println("Esto es darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	nuero, texto := ejercicios.ConvertirANumero("fff")
	fmt.Println(nuero)
	fmt.Println(texto)
}

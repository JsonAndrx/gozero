package main

import (
	"github.com/jsonandrx/gozero/middleware"
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

	// nuero, texto := ejercicios.ConvertirANumero("fff")
	// fmt.Println(nuero)
	// fmt.Println(texto)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// arreglos_slices.Capacidad()

	// mapas.MostrarMapas()

	// defer_panic.EjemploPanic()
	// canal :=  make(chan bool)
	// go gorutines.MiNombreLento("Yeison", canal)
	// defer func ()  {
	// 	<- canal
	// }()
	// fmt.Println("Estoy aqui")

	middleware.MiMiddleware()
}

package funciones

import "fmt"

func Calculos() {

	numero3 := 32
	numero4 := 245

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(suma(10, 25))
}

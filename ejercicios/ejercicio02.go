package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar () {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un numero")

	if input.Scan() {
		num, err := strconv.Atoi(input.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto "+err.Error())
		}

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d X %d = %d \n", num, i, num*i)
		}
	}
}
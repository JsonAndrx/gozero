package ejercicios

import (
	"strconv"
)

func ConvertirANumero(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	}
	return numero, "Es menor a 100"
}

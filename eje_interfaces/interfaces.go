package eje_interfaces

import (
	"fmt"

	"github.com/jsonandrx/gozero/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
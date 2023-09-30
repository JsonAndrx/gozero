package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(time.Millisecond * 1000)
        fmt.Println(letra)
	}
	canal1 <- true
}
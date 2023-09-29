package mapas

import "fmt"

func MostrarMapas() {
	paise := make(map[string]string)
    paise["Argentina"] = "Buenos Aires"
	paise["Brasil"] = "No se"
	// fmt.Println(paise)
	// fmt.Println(paise["Argentina"])

	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 40,
        "Atlético Madrid": 41,
	}

	fmt.Println(campeonato)
	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Barcelona"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
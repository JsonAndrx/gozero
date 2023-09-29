package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 50}

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := []string{"Yeison", "Andres"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i<len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}

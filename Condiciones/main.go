//Condiciones en Go
package main

import "fmt"

var estado bool

func main() {
	estado = true
	//Cambiar variable en el medio de la condicion
	if estado = false; estado == true {
		fmt.Println(estado)
	} else { //Las llaves deben comenzar en la misma linea
		fmt.Println("Es falso")
	}

	// No es necesario el break como en otros lenguajes
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
		//no es necesario el break
	case 2:
		fmt.Println(2)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("No cumplio al switch")
	}

	x := 42
	if x > 18 {
		fmt.Println("Condicion comun")
	}

	// Declaro la variable z en la misma condicion y después la evaluo
	if z := 42; z > 18 {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Not allowed")
	}

}

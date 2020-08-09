package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i : %d", i)
		if i == 5 {
			fmt.Printf(" Multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Printf("    Paso por aca! \n")
		i++

	}

	var z int = 0
	//Seccion, no hace nada a nivel ejecucion
RUTINA:
	fmt.Printf("Esto con un continue no se mostraria \n")
	for z < 10 {
		if z == 4 {
			z = z + 2
			fmt.Println("RUTINA sumando 2 a z")
			//Es como el continue pero en vez de ir al inicio del bucle, va a donde nosotros le digamos
			goto RUTINA
		}
		fmt.Printf("Valor z: %d \n", z)
		z++
	}

}

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
		fmt.Printf("    Paso por aca!")
	}
}

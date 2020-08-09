package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese numero 1: ")
	//Que tipo de dato recibe y a que variable ira
	//Bug en windows, no toma segundo valor
	//fmt.Scanf("%d", numero1)
	//Solucion bug en windows
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")

	//fmt.Scanf("%d", numero2)
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion de cabecera: ")
	//Creamos un nuevo scanner que se alimentara del "Standard Input", que es el teclado por defecto
	scanner := bufio.NewScanner(os.Stdin)
	//Se evalua si ejecutando el metodo scan del objeto se obtiene algo
	if scanner.Scan() {
		//Que scanner me traiga el texto y lo asigne a la variable
		leyenda = scanner.Text()

	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

}

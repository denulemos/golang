package main

import(
	"fmt"
	"strings"
	"time"
)

func main() {
	//Ejecutar la funcion de forma ASINCRONA
    go miNombreLento("Denisse Lemos")
	fmt.Println("Estoy aca")
	var x string
	//Luego de poner la palabra por consola, termina la ejecucion, aun si la otra funcion no termino 
	fmt.Scanln(&x)
}

//Mientras vamos mostrando el nombre por letra, van ocurriendo otras cosas
func miNombreLento(nombre string) {
	//Separamos nombre por letra en un vector
	letras: = string.Split(nombre, "")
	for _, letra: = range letras {
		//Pausa en la ejecucion de 1 segundo. Que cada letra tarde 1 segundo en mostrarse
	   time.Sleep(1000 * time.Millisecond)
	   fmt.Println(letra)
	}
}
package main

import (
	"fmt"
	"strconv"
)

//Todo lo seteado aca, seran variables a nivel package
var numero int
var texto string

//Los booleanos se inicializan en false
var status bool = false

//Si la variable empieza con mayuscula, podra ser exportado en paquete
var NumeroPack int = 8

func main() {

	//Los numero se inicializan en cero
	var numero2, numero4, numero5 int
	//Status aqui sera true
	status = true
	//Asigno directamente y el tipo de variable sera el valor que le di, esta asignacion solo se puede hacer 1 vez por variable, si quiero asignarle un nuevo valor a la variable
	//uso "=" sin los dos puntos, puedo asignar valores a varias variables a la vez
	numero5, numero1 := 2, 5

	var valorInt64 int64 = 0

	//No se puede
	//var numero int = valorInt64
	//Si se puede
	var ejemplo int = int(valorInt64)

	fmt.Println(ejemplo)
	fmt.Println(numero2, numero4, numero5, numero1)

	//Imprimir variable int como decimal con fmt
	fmt.Println(fmt.Sprintf("%d", numero2))

	//Convertir entero a String, solo conviente int, si es int64 da error
	fmt.Println(strconv.Itoa(numero2))

	fmt.Println(status)
	mostrarStatus()
	MostrarTexto()
}

func mostrarStatus() {
	fmt.Println("Hola! Soy la funcion")
	//status aca sera false
	fmt.Println(status)
}

//Lo mismo con las funciones, si empiezan en mayuscula, pueden ser exportados
func MostrarTexto() {
	fmt.Println("Hola! Puedo ser exportado")
	fmt.Println(NumeroPack)
}

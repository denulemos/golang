package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	//Asigno a dos variables los 2 valores que devuelve mi funcion
	numero, estado := dos(1)

	fmt.Println(numero)
	fmt.Println(estado)

	//Mando cualquier cantidad de parametros deseados
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(55, 44, 0))
	fmt.Println(Calculo(24, 33, 77, 100))
}

//Ya le asigno un nombre a la variable de salida
func uno(numero int) (z int) {
	z = numero * 2
	return
}

//Funcion que devuelve 2 datos
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

//Parametros variables, va a recibir parametros enteros, pero no sabemos cuantos
func Calculo(numero ...int) int {
	total := 0
	//Iterar por cada parametro recibido
	//El guion bajo aloja variables que no voy a usar, la devuelve range, y no la voy a usar
	for _, num := range numero {
		total = total + num
	}

	return total
}

/*
- En go no existe la sobrecarga de metodos ni de funciones, si tiene implementado el numero variable de parametros
- Range devuelve 2 valores, se usa mucho en vectores, para recorrer todos los elementos de una lista de datos. El primer valor, es el numero de elemento actual y
el segundo es el elemento en si
*/

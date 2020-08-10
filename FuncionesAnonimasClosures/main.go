package main

import "fmt"

//Variable del tipo funcion
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Sumo 5 + 7 = %d \n ", Calculo(5, 7))

	//Restamos
	//Vamos a redefinir "Calculo", algo que en una funcion normal no es posible
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Println("Restamos 6 - 4 = %d \n ", Calculo(6, 4))

	//Dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Println("Dividimos 12/3 = %d \n", Calculo(12, 3))

	Operaciones()

	tablaDel := 2
	//Mi tabla se convertira en una funcion, ya que tabla devuelve una funcion
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		//Mi tabla ejecutara la funcion de adentro de Tabla, no todo entero, por eso secuencia guarda el valor y no se inicializa siempre en cero
		fmt.Println(MiTabla)
	}
}

//Creamos una funcion con una variable de tipo funcion dentro del mismo
func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

//Ejemplo closure
//Creo una funcion tabla que devuelve una funcion que devuelve un int
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	//La declaracion de ambas firmas debe coincidir
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

/*
-Las funciones anonimas son funciones sin nombre, solo voy a esconder su codigo dentro de una variable y modificarlo en ejecucion si lo necesito
-Closures: Son un concepto de funciones anonimas, tiene que ver con la proteccion/isolacion del codigo. Equivalente al encapsulamiento
*/

/* Un Middleware es un interceptor. Tenemos un desarollo con muchas funciones, y queremos que en cada funcion ahora, se grabe la informacion en un log. Esto nos
llevaria mucho tiempo si hay mucho codigo, los middlewares nos permiten encapsular la ejecucion de una funcion ya existente con una funcion nueva.
Se usan para agregarle a todas las rutinas cosas de seguridad como encriptacion.*/

package main

import "fmt"

//Permite ejecutar instrucciones a muchas funciones que reciben y devuelven el mismo tipo de variable
var result int

func main() {
	//Decidimos que ademas de ejecutar las cuentas, se vea un valor en pantalla tambien
	fmt.Println("Inicio")

	result = operacionesMidd(sumar)(2, 3) //Cual es la funcion, los parametros de sumar
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3) //Cual es la funcion, los parametros de sumar
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3) //Cual es la funcion, los parametros de sumar
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

//Recibe una funcion que recibe en si mismo dos parametros propios de la funcion y devuelve un valor entero, y el Middleware a su vez devuelve la misma funcion del mismo tipo
//Lo que recibe y lo que devuelve TIENE que ser igual
func operacionesMidd(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {
		//Hago lo que necesito hacer
		fmt.Println("Soy el mensaje del Middleware")
		//Devuelvo la funcion para que haga su trabajo :)
		return f(a, b)
	}
}

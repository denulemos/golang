package main

import "fmt"

var tabla [10]int

//5 filas y 7 columnas
var matriz [5][7]int

func main() {
	//Vector
	tabla := [10]int{5, 6, 44, 55, 66, 33, 55, 34, 45, 66}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])

	}

	//Matriz
	matriz[3][5] = 1

	//Slices , vectores dinamicos, puedo ampliar las dimensiones en tiempo de ejecucion
	matrix := []int{1, 2, 3} //Si no asigno longitud, es un slice
	fmt.Println(matrix)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	//Crear un slide en funcion del vector elementos, desde la posicion 3 hasta la ultima
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	//Capacidad no es lo mismo de largo. En memoria se reserva la capacidad, aunque se trabaje inicialmente con 5 elementos
	elementos := make([]int, 5, 20)
	//Print F me permite formatear texto
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

}

func variante4() {
	//Creo un slice vacio
	nums := make([]int, 0, 0)
	//El for es para redimensionar el slice
	for i := 0; i < 100; i++ {
		//Append recibe un slice y devuelve un slice, el append es ideal usarlo cuando nos quedamos sin espacio, no siempre
		nums = append(nums, i)
	}
	//Termina con un largo 100 y capacidad de 128 por la numeracion binaria, siempre va a reservar un poco mas de capacidad.
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}

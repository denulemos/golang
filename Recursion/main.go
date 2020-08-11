package main

import "fmt"

//Recursion => Funcion que se llama a si misma

func main(){
exponencia(2)
}

func exponencia (numero int){ 
	//Le ponemos una condicion de salida para que no sea eterno
	if numero > 100000000{
		return
	}
	fmt.Println(numero)
	//Recursion, la funcion se llama a si misma
	exponencia(numero*2)
}
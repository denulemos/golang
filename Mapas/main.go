//Un mapa es una estructura que se puede serializar, como un vector, pero la diferencia es que consta de clave y valor. No indice y valor

package main

import "fmt"

func main(){
	//Crear mapa donde la key es un string y el valor es un string
	paises := make(map[string]string)
	fmt.PrintLn(paises)

	//Agregar valores al mapa
	paises["Mexico"] = "Distrito Federal"
	paises["Argentina"] = "Buenos Aires"

	fmt.PrintLn(paises)

	//Otra manera de crear mapas pero ya estableciendo la cantidad de elementos para reservar memoria
	paisesAux := make(map[string]string , 5)

	//Otra manera de crear mapas
	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 44,
		"Chivas": 67,
		"Boca": 55}

	fmt.PrintLn(paisesAux)
	

	//Agregar un item
	campeonato["River"] = 66
	//Cambiar un valor
	campeonato["Chivas"] = 45
	//remover item
	delete(campeonato, "Real Madrid")
fmt.PrintLn(campeonato) //Al mostrarlo, lo ordena alfabeticamente por la clave

//Dos valores ya que range devuelve 2 valores, la posicion y el valor
//Range no lo ordena alfabeticamente
for equipo, puntaje := range campeonato{
	fmt.Printf("El equipo %s tiene un puntaje de: %d" , equipo, puntaje)
}

//Ver si existe en mi mapa
puntaje, existe := campeonato["Mineiro"]
//%t es para booleanos 
fmt.Printf("El puntaje capturado es %d, y el equipo existe %t", puntaje, existe)

}
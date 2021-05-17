package main

// Los paquetes se pueden importar de a muchos asi..
import (
"fmt" 
"math"
)

// Si algo empieza con mayuscula, es exportado
func main() {
	// Println empieza con mayuscula ya que fue exportado por fmt
	fmt.Println("Hola Mundo!");
}
/*
-En go no hay try/catch ni excepciones
Defer: Es una instruccion que se va a ejecutar si o si cuando una funcion llega a su final
Panic: Es una funcion que fuerza un error, el sistema abortara
Recover: Cuando se detecta que hay un panic, toma el control del mensaje dado, es un salvavidas 
*/

package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	//Intentar abrir archivo que NO existe
/*	archivo: = "prueba.txt"
	f,err: = os.Open(archivo)
	defer f.Close() //Hay que liberar el buffer de memoria, aunque el archivo no exista

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1) //Cierra la aplicacion con status 1
	} */

	//Aca se ejecuta el defer, ya que aca es donde realmente termina la funcion, este o no el exit antes

ejemploPanic()
}

func ejemploPanic(){
	//Defer de una funcion anonima si quiero hacer mas de 1 instruccion
	defer func(){
		//Si se hubiera encontrado un panic en el fin del programa
		//Recover trae el resultado del ultimo panic, de lo contrario, es nil
reco := Recover()
if reco != nil{
	//%v => Un valor variable
	//Grabo en el log
log.Fatalf("Ocurrio un error que genero un panic \n %v", reco)
}
	}()
	
	a:= 1
	if a==1{
		//Aborta el programa go con status 2 y el mensaje
		panic("Se encontro el valor de 1")
	}
}
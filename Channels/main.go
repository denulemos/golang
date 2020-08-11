package main

import(
	"fmt"
	"time"
)

func main(){
/*Un CANAL es donde van a dialogar distintas rutinas, es un espacio de memoria. 
Transporta una variable, entonces es de un tipo de dato */
canal1 := make(chan time.Duration) //time.Duration calcula la diferencia entre distintas horas (Segundos, milisegundos, etc)
go bucle(canal1) //Pasamos como parametro al canal para dialogar con otras rutinas
fmt.Println("Llegue aca")

//Esperar a que el bucle devuelva un valor en el canal1 y se lo asigne a msg
//Escuchar al canal1, msg sera el resultado de escuchar a canal1
msg := <- canal1 //Similar a await de nodeJS
fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	//Hora inicio
inicio := time.Now()
for i := 0; i < 100000000000000 ; i++{

}

//Hora de final
final := time.Now()
//Sub es la funcion que retorna la duracion de la ejecucion(paquete time)
//<- es asignarle al canal 1 el valor resultante de la funcion
canal1 <- final.Sub(inicio)
}
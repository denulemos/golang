package main

import "net/http"

func main() {
	//Ruta que detecta cuando estoy en "/" y ejecuta home
	http.HandleFunc("/", home)        //Ruta y funcion a ejecutar en consecuencia
	http.ListenAndServe(":3000", nil) //Escucha un puerto, y si detecta info en el mismo, ejecuta algo. Escucha el puerto 3000 y no ejecuta nada por ahora
}

func home(w http.ResponseWriter, r *http.Request) {
	//Cuando se ejecute la funcion home, vamos a mostrar el html
	http.ServeFile(w, r, "./index.html")
}

/*Vamos al navegador y a localhost:3000, ahi va a aparecer el index.html*/

package main

import(
	"fmt"
	"bufio"
	"os"
	"io/ioutil"

)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

//Metodo 1 con ioutil, puedo abrirlo pero no manipularlo muchi
func leoArchivo() {
	//Archivo y variable de error por si lo hay
	//Si la lectura es OK, se guarda el contenido en archivo, si no, se guardara el log en err
	archivo,
	err: = ioutil.ReadFile("./texto.txt")
	//Null en go es nil
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		//Archivo NO es un string, asi que debe ser convertido
		fmt.Println(string(archivo))
	}
}

//Metodo 2 con os, puedo abrir el archivo y hacer algunas modificaciones
func leoArchivo2() {
	archivo,
	err: = os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		//Puedo recorrer secuencialmente registro por registro de mi txt
		scanner: = bufio.NewScanner(archivo)
		//Mientras la linea siguiente sea valida...
		for scanner.Scan() {
			//Convierte a cadena de texto y la guarda en registro
			registro: = scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")

		}
	}
	//Aca si debe ser cerrado
	archivo.Close()
}

func graboArchivo() {
	//Crear archivos
	archivo,
	err: = os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

//Agregamos texto a un archivo que ya existe, creado en graboArchivo
func graboArchivo2() {
	fileName: = "./nuevoArchivo.txt"
	if Append(fileName, "\n Soy otra linea nueva hecha con Append") == false {
		fmt.Println("Hubo un error en la segunda linea")
	}
}

//Append incorpora nuevas lineas
//true o false, si fue false, sucedio un error. Recibe nombre del archivo y el texto a agregar
func Append(archivo string, text string) bool {
	//Es de escritura y de lectura y lo vamos a abrir en modo append, para que no borre el contenido anterior, y el permiso en si (Unix)
	arch,
	err: = os.OpenFile(archivo, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	//Funcion para grabar un string dentro de archivo
	//No nos interesa el primer parametro que nos devuelve WriteString, el segundo si.
	_,
	err = arch.WriteString(text)
	//Si hubo un error..
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
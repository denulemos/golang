/* En Go no existe la POO como la conocemos, aca hay un manejo especial */
package main

//Paquete time maneja fechas
import (
	"fmt"
	"time"
//Referencia a User. Cuando hablamos de us, hablamos de user, dentro de esa carpeta, esta el objeto user
	us "./user"
)

//Herencia 
type pepe struct {
	us.Usuario
}

func main(){
	//Objeto u del tipo Pepe, que posee un puntero a Usuario, hereda de Usuario
	u:= new(pepe)
	u.AltaUsuario(1, "Denu Lemon" , time.Now(), true)
	fmt.Println(u.Usuario)
}
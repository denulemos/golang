package user

import "time"

type Usuario struct {
	Id int
	Nombre string
	FechaAlta time.Time
	Status bool
}

//Funcion, devuelve valores. Metodo no
/*Va a recibir parametros que voy a enviar desde el main cuado haga referencia en el main, pero esta funcion, hara referencia
Al objeto Usuario. Cuando hablo de this, hablo de la estructura usuario. El this *Usuario es un puntero */
func(this * Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	//Funciona como constructor
	//u := Usuario

	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
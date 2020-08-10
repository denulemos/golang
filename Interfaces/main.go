/*
Interfaces: 
*Definen conductas, comportamiento. Por ejemplo, los seres humanos tenemos comportamientos que nos definen como humanos, como pensar,
si creamos un objeto, y decimos que piensa, podemos decir que ese objeto implementa la interfaz ser humano. 
*/

package main
import "fmt"

type SerVivo interface {
	estaVivo() bool

}

type humano interface {
	//Definimos metodos para implementar la interfaz
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

//Genero humano 
type hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
	vivo bool
}


//Hago a una mujer
type mujer struct {
	//Heredamos todas las definiciones del hombre solo con el nombre
	hombre
}

//Funcion que pertenece al hombre, le va a setear true al "respirando" del hombre
func(h * hombre) respirar() {
	h.respirando = true
}
func(h * hombre) comer() {
	h.comiendo = true
}
func(h * hombre) pensar() {
	h.pensando = true
}
func(h * hombre) estaVivo() bool {
	return h.vivo
}
func(h * hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}

}
//El hombre, al tener estas 4 funciones declaradas, ya puede implementar humano, no es necesario el implements

//Esta funcion recibe a cualquier humano, puedo mandar hombre o mujer, ambos son humanos 
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

//Genero Animal
type perro struct {
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func(p * perro) respirar() {
	p.respirando = true
}
func(p * perro) comer() {
	p.comiendo = true
}
func(p * perro) EsCarnivoro() bool {
	return p.carnivoro
}
func(p * perro) estaVivo() bool {
	return p.vivo
}


func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	//Cualquiera de los dos objetos es valido 
	Pedro: = new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)
	Maria: = new(mujer)
	HumanosRespirando(Maria)

	totalcarnivoros: = 0
	Dogo: = new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	//Incrementa en el valor de totalcarnivoros
	totalcarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total Carnivoros %d \n ", totalcarnivoros)

	fmt.Printf("Estoy vivo = %t" , estoyVivo(Dogo))
}
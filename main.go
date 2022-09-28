//  go mod init sample-app  sirve para inicializar el proyecto y se crea el archivo go.mod

// declaramos que el codigo de este archivo estara en el package main
package main

import "fmt"

// todas las aplicaciones deben tener una función main que es el punto de entrada de la app
func main() {
	//camelcase es la forma de dar nombra a las variables
	var conferenceName = "Go conference"
	// para los valores que no van a cambiar se usa const
	const conferenceTickets = 50
	var remainingTickets int = 50
	fmt.Println("welcome to", conferenceName, "booking application")
	fmt.Println("we have", conferenceTickets, "and remains", remainingTickets)
	// printf sirve para sring formatting, es necesario agregar \n para el salto de linea
	fmt.Printf("hay %v de conference y quedan %v\n", conferenceTickets, remainingTickets)

	// se pueden declarar variables vacias pero SI o SI se debe declarar el data type en ese caso, se puede declarar el tipo
	// en variables con valor por si queremos un tipo especifico de int

	var userName string
	var userTickets int
	// asi se pueden declarar variables (NO CONSTANTES) de esta forma no se puede asignarsele un data type y es lo que interprete GO
	ejemplo := "sugar"
	// scan es como input en python, nos pedira que introduzcamos el nombre, el input debe hacerse con un POINTER
	fmt.Println("ingrese su nombre")
	fmt.Scan(&userName)
	// pointer es un tipo de variable especial que nos localiza el lugar de memoria en donde esta almacenada la variable
	fmt.Println("el pointer de userName es ", &userName)

	fmt.Println("ingrese cuantos tickets desea")
	fmt.Scan(&userTickets)
	// es posible saber los tipos con %T
	// https://www.geeksforgeeks.org/data-types-in-go/
	fmt.Printf("user name es %T y user tickets es %T y sugar es %T\n", userName, userTickets, ejemplo)
	fmt.Printf("user name es %v y user tickets es %v y sugar es %v\n", userName, userTickets, ejemplo)
	// es posible hacer esta operacion ya que ambos son int, solo se pueden hacer operaciones en variables iguales
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("quedan %v tickets disponibles", remainingTickets)

}

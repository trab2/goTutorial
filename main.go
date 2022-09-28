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
	var remainingTickets = 50
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

	userName = "tomacito"
	userTickets = 2
	// es posible saber los tipos con %T
	// https://www.geeksforgeeks.org/data-types-in-go/
	fmt.Printf("user name es %T y user tickets es %T y sugar es %T", userName, userTickets, ejemplo)

}

package main

import "fmt"

func main() {

	name := "Santiago"
	dir := "Buenos Aires, Argentina"

	fmt.Println(name, dir)
	clima()
}

func clima() {

	var temperatura float64
	var humedad float64
	var presion float64

	temperatura = 10.2
	humedad = 20.4
	presion = 11.3

	fmt.Println("temperatura:", temperatura, "\nhumedad:", humedad, "\npresion:", presion)

}

/*Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.
Necesita ayuda para:
Detectar cuáles de estas variables que declaró el alumno son correctas.*/
// Corregir las incorrectas.

// var 1nombre string --> Empieza con un numero
// var apellido string --> Esta bien
// var int edad --> Esta mal el orden
// 1apellido := 6 -->  Empieza con un numero
// var licencia_de_conducir = true --> Esta bien
// var estatura de la persona int --> El nombre tiene espacios
// cantidadDeHijos := 2 --> Bien

// Un estudiante de programación intentó realizar declaraciones de
// variables con sus respectivos tipos en Go pero tuvo varias dudas mi
// entras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un
// desarrollador experimentado que pueda:
// Verificar su código y realizar las correcciones necesarias.

//   var apellido string = "Gomez" --> bien
//   var edad int = "35" --> mal
//   boolean := "false"; --> mal
//   var sueldo string = 45857.90 --> mal
//   var nombre string = "Julián" --> bien

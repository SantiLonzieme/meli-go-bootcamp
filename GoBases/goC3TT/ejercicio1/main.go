package main

import (
	"fmt"
)

func main() {

	usuar := Usuario{}
	usuar.cambiarEdad(32)

	fmt.Println(usuar)
}

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *Usuario) cambiarNombre(nombre string, apellido string) {

	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) cambiarEdad(edad int) {

	u.Edad = edad
}

func (u *Usuario) cambiarCorreo(correo string) {

	u.Correo = correo
}

func (u *Usuario) cambiarContraseña(contraseña string) {

	u.Contraseña = contraseña
}

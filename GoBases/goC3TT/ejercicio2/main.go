package main

import (
	"fmt"
)

func main() {

	fmt.Println()
}

type Usuario struct {
	Nombre   string
	Apellido string
	Correo   string
	Producto []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p *Producto) nuevoProducto(nombre string, precio float64) Producto {

	producto := Producto{nombre, precio, 1}

	return producto
}

func (u *Usuario) agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {

	producto.Cantidad = cantidad
	usuario.Producto = append(usuario.Producto, *producto)

}

func (u *Usuario) borrarProductos(usuario *Usuario) {

	usuario.Producto = nil
}

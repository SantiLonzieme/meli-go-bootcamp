package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

func main() {

	p1 := Persona{Nombre: "Santi", Musica: Musica{Banda: "Pearl Jam"}}
	p1.Nombre = "Juan"
	p1.Musica.Disco = "Yield"
	p1.Musica.Estilo = "Grunge"

	miJosn, err := json.Marshal(p1)

	fmt.Println(string(miJosn), err)

	////////////////////////////////

	p2 := Persona{Nombre: "Santi", Musica: Musica{Banda: "Pearl Jam"}}
	p2.Nombre = "Juan"
	p2.Musica.Disco = "Yield"
	p2.Musica.Estilo = "Grunge"

	p := reflect.TypeOf(p2)

	fmt.Println("///////////////////////////")
	fmt.Println(p, "REFLECT")
	fmt.Println(p.NumField(), "Numero de campos")
	fmt.Println(p.Field(2), "Obtener un campo")
	fmt.Println(p.Field(2).Tag.Get("json"), "Etiqueta")
	fmt.Println("///////////////////////////")

	/////////////////////////////////////////////////

	circulo := Circulo{5}

	fmt.Println(circulo.area())
	fmt.Println(circulo.perim())
	circulo.setRadio(10)
	fmt.Println(circulo, "circulo")
	fmt.Println(circulo.area())
	fmt.Println(circulo.perim())

	fmt.Println("///////////////////////////")

	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()

	fmt.Println("///////////////////////////GEOMETRY")

	r := newGeometry(rectType, 2, 3)
	fmt.Println(r.area())
	fmt.Println(r.perim())
	c := newGeometry(circleType, 2)
	fmt.Println(c.area())
	fmt.Println(c.perim())

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
}

type Persona struct {
	Nombre    string `json:"primer_nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"Edad"`
	Localidad string `json:"Localidad"`
	Musica    Musica
}

type Musica struct {
	Banda  string
	Estilo string
	Disco  string
}

type Circulo struct {
	radio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func (c *Circulo) setRadio(r float64) {
	c.radio = r
}

type Vehiculo struct {
	km     float64
	tiempo float64
}

type Auto struct {
	vehiculo Vehiculo
}

type Moto struct {
	vehiculo Vehiculo
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, heigth float64
}

func (r rect) area() float64 {
	return r.heigth * r.width
}

func (r rect) perim() float64 {
	return 2*r.heigth + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], heigth: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (v Vehiculo) detalle() {
	fmt.Printf("km:\t%f\ntiempo:\t%f\n", v.km, v.tiempo) //ver en las diapos
}

func (a *Auto) Correr(minutos int) {
	a.vehiculo.tiempo = float64(minutos) / 60
	a.vehiculo.km = a.vehiculo.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.vehiculo.detalle()
}

func (m *Moto) Correr(minutos int) {
	m.vehiculo.tiempo = float64(minutos) / 60
	m.vehiculo.km = m.vehiculo.tiempo * 80
}

func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.vehiculo.detalle()
}

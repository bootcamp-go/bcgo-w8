package main

import "fmt"

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string	`json:"apellido"`
	Edad     int	`json:"edad"`
}
func (p Persona) NombreCompleto() string {
	return p.Nombre + " " + p.Apellido
}
func (p *Persona) NuevoNombre(name string) {
	(*p).Nombre = "Nuevo nombre"
}

func NombreCompleto(p Persona) string {
	return p.Nombre + " " + p.Apellido
}
func NuevoNombre(p *Persona) {
	(*p).Nombre = "Nuevo nombre"
}

func main() {
	p1 := Persona{
		Nombre:   "Juan",
		Apellido: "Perez",
		Edad:     20,
	}

	p1.NuevoNombre("Carlos")

	fmt.Println("Nombre completo:", p1.NombreCompleto())
}
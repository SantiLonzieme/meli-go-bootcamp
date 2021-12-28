package usuarios

import (
	"fmt"

	"github.com/SantiLonzieme/goweb/pkg/store"
)

type Usuario struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Edad     int    `json:"edad"`
	Altura   int    `json:"altura"`
	Activo   bool   `json:"activo"`
	Fecha    string `json:"fecha"`
}

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura int, activo bool, fecha string) (Usuario, error)
	LastId() (int, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura int, activo bool, fecha string) (Usuario, error)
	UpdateApellidoEdad(id int, apellido string, edad int) (Usuario, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

var usuarios []Usuario

// var lastID int

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Usuario, error) {
	var usuarios []Usuario
	r.db.Read(&usuarios)
	return usuarios, nil
}

func (r *repository) LastId() (int, error) {
	var usuarios []Usuario

	if err := r.db.Read(&usuarios); err != nil {
		return 0, err
	}

	if len(usuarios) == 0 {
		return 0, nil
	}

	return usuarios[len(usuarios)-1].Id, nil
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura int, activo bool, fecha string) (Usuario, error) {

	var usuarios []Usuario
	r.db.Read(&usuarios)

	usuario := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}

	usuarios = append(usuarios, usuario)

	if err := r.db.Write(usuarios); err != nil {
		return Usuario{}, err
	}
	// lastID = usuario.Id
	return usuario, nil
}

func (r *repository) Update(id int, nombre string, apellido string, email string, edad int,
	altura int, activo bool, fecha string) (Usuario, error) {

	u := Usuario{
		Nombre:   nombre,
		Apellido: apellido,
		Email:    email,
		Edad:     edad,
		Altura:   altura,
		Activo:   activo,
		Fecha:    fecha,
	}

	updated := false

	r.db.Read(&usuarios)

	for i := range usuarios {
		if usuarios[i].Id == id {
			u.Id = id
			usuarios[i] = u
			updated = true
		}
	}

	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrando", id)
	}

	if err := r.db.Write(usuarios); err != nil {
		return Usuario{}, err
	}

	return u, nil
}

func (r *repository) Delete(id int) error {
	deleted := false

	var index int

	r.db.Read(&usuarios)

	for i := range usuarios {
		if usuarios[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}

	usuarios = append(usuarios[:index], usuarios[:index+1]...)

	if err := r.db.Write(usuarios); err != nil {
		return err
	}

	return nil

}

func (r *repository) UpdateApellidoEdad(id int, apellido string, edad int) (Usuario, error) {
	var u Usuario

	updated := false

	r.db.Read(&usuarios)

	for i := range usuarios {
		if usuarios[i].Id == id {
			usuarios[i].Apellido = apellido
			usuarios[i].Edad = edad
			updated = true
			u = usuarios[i]
		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	if err := r.db.Write(usuarios); err != nil {
		return Usuario{}, err
	}
	return u, nil
}

package usuarios

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
	Store(id int, nombre string, apellido string, email string, edad int,
		altura int, activo bool, fecha string) (Usuario, error)
	LastId() (int, error)
}

type repository struct{}

var usuarios []Usuario
var lastID int

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Usuario, error) {
	return usuarios, nil
}

func (r *repository) LastId() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura int, activo bool, fecha string) (Usuario, error) {
	usuario := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha}
	usuarios = append(usuarios, usuario)
	lastID = usuario.Id
	return usuario, nil
}

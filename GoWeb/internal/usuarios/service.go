package usuarios

type Service interface {
	GetAll() ([]Usuario, error)
	//Store no recibe Id
	Store(id int, nombre string, apellido string, email string, edad int, altura int, activo bool, fecha string) (Usuario, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura int, activo bool, fecha string) (Usuario, error)
	UpdateApellidoEdad(id int, apellido string, edad int) (Usuario, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Usuario, error) {
	usuarios, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (s *service) Store(id int, nombre string, apellido string, email string, edad int,
	altura int, activo bool, fecha string) (Usuario, error) {
	lastID, err := s.repository.LastId()

	if err != nil {
		return Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fecha)

	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}

func (s *service) Update(id int, nombre string, apellido string, email string, edad int,
	altura int, activo bool, fecha string) (Usuario, error) {

	return s.repository.Update(id, nombre, apellido, email,
		edad, altura, activo, fecha)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateApellidoEdad(id int, apellido string, edad int) (Usuario, error) {
	return s.repository.UpdateApellidoEdad(id, apellido, edad)
}

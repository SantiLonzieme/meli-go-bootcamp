package usuarios

type Service interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre string, apellido string, email string, altura int, activo bool, fecha string) (Usuario, error)
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

func (s *service) Store(id int, nombre string, apellido string, email string,
	altura int, activo bool, fecha string) (Usuario, error) {
	lastID, err := s.repository.LastId()

	if err != nil {
		return Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nombre, apellido, email,
		altura, activo, fecha)

	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}

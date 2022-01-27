package usuarios

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

func (s *stubStore) Read(data interface{}) error {
	user1 := Usuario{1, "Santi", "Lonzieme", "san@gmail.com", 50, 180, true, "10/1/2022"}
	user2 := Usuario{2, "Juan", "Perez", "juan@gmail.com", 40, 178, true, "9/1/2022"}
	users := []Usuario{user1, user2}

	usersJson, _ := json.MarshalIndent(users, "", " ")

	err := json.Unmarshal(usersJson, &data)

	if err != nil {
		return err
	}

	return nil
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	usersRepo := repository{db: &stubStore{}}
	users, err := usersRepo.GetAll()

	assert.Equal(t, 2, len(users))
	assert.Error(t, err, "error")

}

type fileStoreSpy struct {
	Usuarios []Usuario
	Called   bool
}

func (s *fileStoreSpy) Read(data interface{}) error {
	s.Called = true

	userJson, _ := json.MarshalIndent(s.Usuarios, "", " ")
	err := json.Unmarshal(userJson, &data)

	if err != nil {
		return err
	}

	return nil
}

func (s *fileStoreSpy) Write(data interface{}) error {
	s.Usuarios, _ = data.([]Usuario)
	return nil
}

func TestUpdateCodigo(t *testing.T) {

	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	spy := &fileStoreSpy{Usuarios: users}
	repo := repository{db: spy}

	user, err := repo.UpdateApellidoEdad(1, "After", 24)

	assert.True(t, spy.Called)
	assert.Equal(t, user.Apellido, "After")
	assert.Equal(t, err, nil)
}

func TestLastIdRepo(t *testing.T) {

	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := repository{db: db}

	lastId, err := repo.LastId()

	assert.IsType(t, user.Id, lastId)
	assert.Nil(t, err)
}

func TestStore(t *testing.T) {
	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := repository{db: db}

	userStore, err := repo.Store(1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022")

	assert.Equal(t, user, userStore)
	assert.Equal(t, err, nil)

}

func TestUpdateRepo(t *testing.T) {
	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := repository{db: db}

	userUpdated, err := repo.Update(1, "SantiL", "Before", "san@gmail.com", 50, 180, true, "10/1/2022")

	assert.NotEqual(t, user.Nombre, userUpdated.Nombre)
	assert.Equal(t, err, nil)
}

func TestErrorRepo(t *testing.T) {
	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := repository{db: db}

	var fake int
	err := repo.db.Write(fake)

	assert.NotNil(t, err)
}

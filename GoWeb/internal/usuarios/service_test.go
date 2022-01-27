package usuarios

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Mock struct {
	Data []byte
	Err  error
}
type FileStore struct {
	FileName string
	Mock     *Mock
}

func (f *FileStore) Read(data interface{}) error {
	return nil
}

func (f *FileStore) Write(data interface{}) error {
	return nil
}

func TestUpdate(t *testing.T) {
	users := []Usuario{{1, "Santi", "Lonzieme", "san@gmail.com", 50, 180, true, "10/1/2022"}}

	userJson, _ := json.Marshal(users)

	dbMock := Mock{
		Data: userJson,
	}

	storeStub := FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	result, err := service.GetAll()

	assert.Equal(t, users, result)
	assert.Nil(t, err)

}

func TestDelete(t *testing.T) {
	users := []Usuario{}

	userJson, _ := json.Marshal(users)

	dbMock := Mock{
		Data: userJson,
	}

	storeStub := FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	result := service.Delete(1)

	assert.Equal(t, users, result)

}

func TestStoreSer(t *testing.T) {
	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := NewRepository(db)
	service := NewService(repo)

	userStore, err := service.Store(1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022")

	assert.Equal(t, user, userStore)
	assert.Equal(t, err, nil)

}

func TestUpdateServ(t *testing.T) {
	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := NewRepository(db)
	service := NewService(repo)

	userUpdated, err := service.Update(1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022")

	assert.IsType(t, Usuario{}, userUpdated)
	assert.Equal(t, err, nil)
}

func TestUpdateApellidoEdad(t *testing.T) {

	user := Usuario{1, "Santi", "Before", "san@gmail.com", 50, 180, true, "10/1/2022"}
	users := []Usuario{user}

	db := &fileStoreSpy{Usuarios: users}
	repo := NewRepository(db)
	service := NewService(repo)

	userUpdated, err := service.UpdateApellidoEdad(1, "Perez", 40)

	assert.NotEqual(t, user.Apellido, userUpdated.Apellido)
	assert.Equal(t, err, nil)
}

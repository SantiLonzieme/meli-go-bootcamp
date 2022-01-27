package product

import (
	"database/sql"
	"testing"

	"github.com/SantiLonzieme/sql/internal/models"
	"github.com/SantiLonzieme/sql/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {

	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	newProduct := models.Product{
		Name:  "Santi",
		Type:  "Auto",
		Count: 1,
		Price: 200.50,
	}

	serv := NewRepository(db)

	res, _ := serv.Store(newProduct)

	assert.Equal(t, newProduct.Name, res.Name)
	assert.Equal(t, newProduct.Type, res.Type)
	assert.Equal(t, newProduct.Count, res.Count)
	assert.Equal(t, newProduct.Price, res.Price)
}

func TestGetAll(t *testing.T) {

	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	serv := NewRepository(db)

	res, _ := serv.GetAll()

	assert.IsType(t, []models.Product{}, res)
}

func TestGetByName(t *testing.T) {

	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")

	newProduct := models.Product{
		Name:  "Santi",
		Type:  "Auto",
		Count: 1,
		Price: 200.50,
	}

	serv := NewRepository(db)

	res, _ := serv.GetByName(newProduct.Name)

	assert.Equal(t, newProduct.Name, res.Name)
	assert.IsType(t, models.Product{}, res)
}

func TestSqlRepositoryStore(t *testing.T) {

	db, err := util.InitDb()

	assert.NoError(t, err)

	repository := NewRepository(db)

	newProduct := models.Product{
		Name:  "Santi",
		Type:  "Auto",
		Count: 1,
		Price: 200.50,
	}

	respStore, err := repository.Store(newProduct)
	assert.NoError(t, err)
	assert.NotEqual(t, newProduct, respStore)

	respGet, err := repository.Get(newProduct.ID)

	assert.NoError(t, err)
	assert.Equal(t, newProduct.Name, respGet.Name)

}

func TestSqlRepositorySave(t *testing.T) {

	db, err := util.InitDb()

	assert.NoError(t, err)

	repository := NewRepository(db)

	newProduct := models.Product{
		Name:  "Santi",
		Type:  "Auto",
		Count: 1,
		Price: 200.50,
	}

	err = repository.Update(newProduct)
	assert.NoError(t, err)
}

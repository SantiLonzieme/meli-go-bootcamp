package product

import (
	"github.com/SantiLonzieme/sql/internal/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetByName(name string) (models.Product, error)
	GetById(id int) (models.Product, error)
	Create(name string, typeProduct string,
		count int, price float64) (models.Product, error)
	GetAll() ([]models.Product, error)
	Update(id int, name string, typeProduct string,
		count int, price float64) (models.Product, error)
	UpdateWithContext(ctx *gin.Context, id int, name string, typeProduct string,
		count int, price float64) (models.Product, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]models.Product, error) {

	products, err := s.repository.GetAll()

	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (s *service) GetByName(name string) (models.Product, error) {

	product, err := s.repository.GetByName(name)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) GetById(id int) (models.Product, error) {

	product, err := s.repository.Get(id)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

func (s *service) Create(name string, typeProduct string,
	count int, price float64) (models.Product, error) {

	newProduct := models.Product{
		Name:  name,
		Type:  typeProduct,
		Count: count,
		Price: price,
	}

	res, err := s.repository.Store(newProduct)

	if err != nil {
		return models.Product{}, err
	}

	return res, nil
}

func (s *service) Update(id int, name string, typeProduct string,
	count int, price float64) (models.Product, error) {

	product, err := s.repository.Get(int(id))

	if err != nil {
		return models.Product{}, err
	}

	productToUp := models.Product{Name: name, Type: typeProduct,
		Count: count, Price: price, ID: id}

	err = s.repository.Update(productToUp)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) UpdateWithContext(ctx *gin.Context, id int, name string, typeProduct string,
	count int, price float64) (models.Product, error) {

	product, err := s.repository.Get(int(id))

	if err != nil {
		return models.Product{}, err
	}

	productToUp := models.Product{Name: name, Type: typeProduct,
		Count: count, Price: price, ID: id}

	err = s.repository.UpdateWithContext(ctx, productToUp)

	if err != nil {
		return models.Product{}, err
	}

	product.ID = id

	return product, nil
}

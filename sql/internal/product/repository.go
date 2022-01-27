package product

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SantiLonzieme/sql/internal/models"
)

const (
	queryByName          = "select id, name, type, count, price from products where name = ?"
	queryInsertAllFields = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	queryGetAll          = "select id, name, type, count, price from products"
	queryUpdate          = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?"
	queryById            = "select id, name, type, count, price from products WHERE id=?;"
)

type Repository interface {
	GetByName(name string) (models.Product, error)
	Store(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Update(product models.Product) error
	UpdateWithContext(ctx context.Context, product models.Product) error
	Get(id int) (models.Product, error)
}
type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(name string) (models.Product, error) {
	var product models.Product
	db := r.db
	rows, err := db.Query(queryByName, name)
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return models.Product{}, err
		}
	}
	return product, nil
}

func (r *repository) Store(product models.Product) (models.Product, error) {
	db := r.db
	stmt, err := db.Prepare(queryInsertAllFields)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return models.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	var products []models.Product
	db := r.db
	rows, err := db.Query(queryGetAll)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Get(id int) (models.Product, error) {

	row := r.db.QueryRow(queryById, id)
	p := models.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

func (r *repository) Update(product models.Product) error {

	stmt, err := r.db.Prepare(queryUpdate)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", product)

	res, err := stmt.Exec(product.Name, product.Type, product.Count,
		product.Price, product.ID)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()

	if err != nil {
		return err
	}

	fmt.Print(ra)

	return nil
}

func (r *repository) UpdateWithContext(ctx context.Context, product models.Product) error {

	stmt, err := r.db.Prepare(queryUpdate)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

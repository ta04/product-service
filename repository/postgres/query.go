// product-service/repository/postgres/query.go

package postgres

import (
	"database/sql"
	"fmt"

	productPB "github.com/SleepingNext/product-service/proto"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) Index() (products []*productPB.Product, err error) {
	var id int32
	var name, description, picture string
	var price float64
	var status string

	query := "SELECT * FROM products"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &picture, &status)
		if err != nil {
			return nil, err
		}
		product := &productPB.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Picture:     picture,
			Price:       price,
			Status:      status,
		}
		products = append(products, product)
	}

	return products, err
}

func (repo *Repository) Show(product *productPB.Product) (*productPB.Product, error) {
	var id int32
	var name, description, picture string
	var price float64
	var status string

	query := fmt.Sprintf("SELECT * FROM products WHERE id=%d", product.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &name, &description, &price, &picture, &status)
	if err != nil {
		return nil, err
	}

	return &productPB.Product{
		Id:          id,
		Name:        name,
		Description: description,
		Picture:     picture,
		Price:       price,
		Status:      status,
	}, err
}

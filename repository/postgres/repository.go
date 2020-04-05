// product-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	productPB "github.com/SleepingNext/product-service/proto"
)

// Store will store a new product
func (repo *Repository) Store(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("INSERT INTO products (name, description, price, picture, status)"+
		" VALUES ('%s', '%s', %f, '%s', 'active')", product.Name, product.Description, product.Price, product.Picture)
	_, err := repo.DB.Exec(query)

	return product, err
}

// Update will update an existing product's data
func (repo *Repository) Update(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("UPDATE products SET name = '%s', description = '%s', price = %f, picture = '%s', status = 'active'"+
		" WHERE id = %d", product.Name, product.Description, product.Price, product.Picture, product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

// Destroy will update an existing product's status to inactive
func (repo *Repository) Destroy(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("UPDATE products SET status = 'inactive' WHERE id=%d", product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

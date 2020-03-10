// product-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	productPB "github.com/SleepingNext/product-service/proto"
)

func (repo *Repository) Store(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("INSERT INTO products (name, description, price, picture, status)"+
		" VALUES ('%s', '%s', %f, '%s', %s)", product.Name, product.Description, product.Price, product.Picture, product.Status)
	_, err := repo.DB.Exec(query)

	return product, err
}

func (repo *Repository) Update(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("UPDATE products SET name = '%s', description = '%s', price = %f, picture = '%s', status = %s"+
		" WHERE id = %d", product.Name, product.Description, product.Price, product.Picture, product.Status, product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

func (repo *Repository) Destroy(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("DELETE FROM products WHERE id=%d", product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package postgres

import (
	"fmt"

	productPB "github.com/SleepingNext/product-service/proto"
)

// Store stores a new product
func (repo *Postgres) Store(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("INSERT INTO products (name, description, price, picture, status)"+
		" VALUES ('%s', '%s', %f, '%s', 'active')", product.Name, product.Description, product.Price, product.Picture)
	_, err := repo.DB.Exec(query)

	return product, err
}

// Update updates a product
func (repo *Postgres) Update(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("UPDATE products SET name = '%s', description = '%s', price = %f, picture = '%s', status = 'active'"+
		" WHERE id = %d", product.Name, product.Description, product.Price, product.Picture, product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

// Destroy updates an existing product's status to inactive
func (repo *Postgres) Destroy(product *productPB.Product) (*productPB.Product, error) {
	query := fmt.Sprintf("UPDATE products SET status = 'inactive' WHERE id=%d", product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

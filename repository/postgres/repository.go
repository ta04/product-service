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

	proto "github.com/ta04/product-service/model/proto"
)

// CreateOne will create a new product
func (postgres *Postgres) CreateOne(product *proto.Product) (*proto.Product, error) {
	query := fmt.Sprintf("INSERT INTO products (name, description, price, picture, status)"+
		" VALUES ('%s', '%s', %f, '%s', 'active')", product.Name, product.Description, product.Price, product.Picture)
	_, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateOne will update a product
func (postgres *Postgres) UpdateOne(product *proto.Product) (*proto.Product, error) {
	query := fmt.Sprintf("UPDATE products SET name = '%s', description = '%s', price = %f, picture = '%s', status = '%s'"+
		" WHERE id = %d", product.Name, product.Description, product.Price, product.Status, product.Picture, product.Id)
	_, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	return product, nil
}

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
	"database/sql"
	"fmt"

	proto "github.com/ta04/product-service/model/proto"
)

// Postgres is the implementor of Repository interface
type Postgres struct {
	DB *sql.DB
}

// NewPostgres will create a new postgres instance
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

// GetAllByQuery will get all products by query
func (postgres *Postgres) GetAllByQuery(request *proto.GetAllProductsRequest) (*[]*proto.Product, error) {
	var id int32
	var name, description, picture, status string
	var price float64
	var products []*proto.Product

	query := fmt.Sprintf("SELECT * FROM products WHERE (LOWER(name) LIKE '%%%s%%' OR LOWER(description) LIKE" +
		" '%%%s%%') AND status = '%s'", request.Query, request.Query, request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &picture, &status)
		if err != nil {
			return nil, err
		}

		product := &proto.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Picture:     picture,
			Price:       price,
			Status:      status,
		}
		products = append(products, product)
	}

	return &products, err
}

// GetAll will get all products
func (postgres *Postgres) GetAll(request *proto.GetAllProductsRequest) (*[]*proto.Product, error) {
	var id int32
	var name, description, picture, status string
	var price float64
	var products []*proto.Product

	query := fmt.Sprintf("SELECT * FROM products WHERE status = '%s'", request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &picture, &status)
		if err != nil {
			return nil, err
		}

		product := &proto.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Picture:     picture,
			Price:       price,
			Status:      status,
		}
		products = append(products, product)
	}

	return &products, err
}

// GetOne will get a product by id
func (postgres *Postgres) GetOne(request *proto.GetOneProductRequest) (*proto.Product, error) {
	var id int32
	var name, description, picture, status string
	var price float64

	query := fmt.Sprintf("SELECT * FROM products WHERE id = %d", request.Id)
	err := postgres.DB.QueryRow(query).Scan(&id, &name, &description, &price, &picture, &status)
	if err != nil {
		return nil, err
	}

	return &proto.Product{
		Id:          id,
		Name:        name,
		Description: description,
		Picture:     picture,
		Price:       price,
		Status:      status,
	}, err
}

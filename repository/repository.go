package repository

import (
	"database/sql"
	"fmt"
	pb "github.com/SleepingNext/product-service/proto"
)

type Repository interface {
	Index() ([]*pb.Product, error)
	Show(*pb.Product) (*pb.Product, error)
	Store(*pb.Product) (*pb.Product, error)
	Update(*pb.Product) (*pb.Product, error)
	Destroy(*pb.Product) (*pb.Product, error)
}

type PostgresRepository struct {
	DB *sql.DB
}

func (repo *PostgresRepository) Index() (products []*pb.Product, err error) {
	var id int32
	var name, description, picture string
	var price float64
	var status bool

	query := "SELECT * FROM product"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &picture, &status)
		if err != nil {
			return nil, err
		}
		product := pb.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Picture:     picture,
			Price:       price,
			Status:      status,
		}
		products = append(products, &product)
	}

	return products, err
}

func (repo *PostgresRepository) Show(product *pb.Product) (*pb.Product, error) {
	var id int32
	var name, description, picture string
	var price float64
	var status bool

	query := fmt.Sprintf("SELECT * FROM product WHERE id=%d", product.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &name, &description, &price, &picture, &status)
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          id,
		Name:        name,
		Description: description,
		Picture:     picture,
		Price:       price,
		Status:      status,
	}, err
}

func (repo *PostgresRepository) Store(product *pb.Product) (*pb.Product, error) {
	query := fmt.Sprintf("INSERT INTO product (name, description, price, picture, status)"+
		" VALUES ('%s', '%s', %f, '%s', %t)", product.Name, product.Description, product.Price, product.Picture, product.Status)
	_, err := repo.DB.Exec(query)

	return product, err
}

func (repo *PostgresRepository) Update(product *pb.Product) (*pb.Product, error) {
	query := fmt.Sprintf("UPDATE product SET name = '%s', description = '%s', price = %f, picture = '%s', status = %t"+
		" WHERE id = %d", product.Name, product.Description, product.Price, product.Picture, product.Status, product.Id)
	_, err := repo.DB.Exec(query)

	return product, err
}

func (repo *PostgresRepository) Destroy(product *pb.Product) (*pb.Product, error) {
	query := fmt.Sprintf("DELETE FROM product WHERE id=%d", product.Id)
	_, err := repo.DB.Exec(query)

	return product, err

}

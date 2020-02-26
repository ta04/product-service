// product-service/main.go

package main

import (
	"context"
	"database/sql"
	"fmt"
	pb "github.com/SleepingNext/product-service/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

const (
	port = ":50051"
)

type repository interface {
	Index() ([]*pb.Product, error)
	Show(*pb.Product) (*pb.Product, error)
	Store(*pb.Product) (*pb.Product, error)
	Update(*pb.Product) (*pb.Product, error)
	Destroy(*pb.Product) (*pb.Product, error)
}

type Repository struct {
	mu sync.RWMutex
	db *sql.DB
}

func (repo *Repository) Index() (products []*pb.Product, err error) {
	var id int32
	var name, description, picture string
	var price float64
	var status bool

	query := "SELECT * FROM product"
	rows, err := repo.db.Query(query)
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

func (repo *Repository) Show(product *pb.Product) (*pb.Product, error) {
	var id int32
	var name, description, picture string
	var price float64
	var status bool

	query := fmt.Sprintf("SELECT * FROM product WHERE id=%d", product.Id)
	err := repo.db.QueryRow(query).Scan(&id, &name, &description, &price, &picture, &status)
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

func (repo *Repository) Store(product *pb.Product) (*pb.Product, error) {
	repo.mu.Lock()
	query := fmt.Sprintf("INSERT INTO product (name, description, price, picture, status)"+
		" VALUES ('%s', '%s', %f, '%s', %t)", product.Name, product.Description, product.Price, product.Picture, product.Status)
	_, err := repo.db.Exec(query)
	repo.mu.Unlock()

	return product, err
}

func (repo *Repository) Update(product *pb.Product) (*pb.Product, error) {
	repo.mu.Lock()
	query := fmt.Sprintf("UPDATE product SET name = '%s', description = '%s', price = %f, picture = '%s', status = %t"+
		" WHERE id = %d", product.Name, product.Description, product.Price, product.Picture, product.Status, product.Id)
	_, err := repo.db.Exec(query)
	repo.mu.Unlock()

	return product, err
}

func (repo *Repository) Destroy(product *pb.Product) (*pb.Product, error) {
	repo.mu.Lock()
	query := fmt.Sprintf("DELETE FROM product WHERE id=%d", product.Id)
	_, err := repo.db.Exec(query)
	repo.mu.Unlock()

	return product, err

}

type service struct {
	repo repository
}

func (s *service) IndexProducts(ctx context.Context, req *pb.IndexProductsRequest) (*pb.Response, error) {
	products, err := s.repo.Index()

	return &pb.Response{
		Products: products,
		Error:    nil,
	}, err
}

func (s *service) ShowProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	product, err := s.repo.Show(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Product: product,
		Error:   nil,
	}, nil
}

func (s *service) StoreProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	product, err := s.repo.Store(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Product: product,
		Error:   nil,
	}, err
}

func (s *service) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	product, err := s.repo.Update(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Product: product,
		Error:   nil,
	}, nil
}

func (s *service) DestroyProduct(ctx context.Context, req *pb.Product) (*pb.Response, error) {
	product, err := s.repo.Destroy(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Product: product,
		Error:   nil,
	}, err
}

func connectPostgres() (*sql.DB, error) {
	connStr := "user=sleepingnext dbname=product sslmode=disable password=kevin99123"
	db, err := sql.Open("postgres", connStr)

	return db, err
}

func main() {
	db, err := connectPostgres()
	repo := &Repository{}
	repo.db = db

	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterProductServiceServer(s, &service{repo})

	reflection.Register(s)

	log.Println("running on port", port)
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

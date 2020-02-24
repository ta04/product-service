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
	Create(*pb.Product) (*pb.Product, error)
	GetAll() ([]*pb.Product, error)
}

type Repository struct {
	mu sync.RWMutex
	db *sql.DB
}

func (repo *Repository) Create(product *pb.Product) (*pb.Product, error) {
	repo.mu.Lock()
	query := fmt.Sprintf("INSERT INTO product (name, description, price, picture, status)" +
		"VALUES ('%s', '%s', %f, '%s', %t)", product.Name, product.Description, product.Price, product.Picture, product.Status)
	_, err := repo.db.Exec(query)
	repo.mu.Unlock()

	return product, err
}

func (repo *Repository) GetAll() (products []*pb.Product, err error) {
	var id int32
	var name, description, picture string
	var price float64
	var status bool

	query := "SELECT * FROM product"
	rows, err := repo.db.Query(query)

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &picture, &status)
		if err != nil {
			return nil, err
		}
		product := pb.Product{
			Id: id,
			Name: name,
			Description: description,
			Picture: picture,
			Price: price,
			Status: status,
		}
		products = append(products, &product)
	}

	return products, err
}

type service struct {
	repo repository
}

func (s *service) CreateProduct(ctx context.Context, req *pb.Product) (*pb.CreateProductResponse, error) {
	product, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		CreatedProduct: product,
		Error: nil,
	}, err
}

func (s *service) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	products, err := s.repo.GetAll()

	return &pb.GetAllProductsResponse{
		Products: products,
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

	log.Println("Running on port:", port)
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
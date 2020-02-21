// product-service/main.go

package main

import (
	"context"
	pb "github.com/SleepingNext/product-service/proto"
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
}

type Repository struct {
	mu sync.RWMutex
	products []*pb.Product
}

func (repo *Repository) Create(product *pb.Product) (*pb.Product, error) {
	repo.mu.Lock()
	updated := append(repo.products, product)
	repo.products = updated
	repo.mu.Unlock()

	return product, nil
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

func main() {
	repo := &Repository{}

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
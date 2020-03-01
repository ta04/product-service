package handler

import (
	"context"
	pb "github.com/SleepingNext/product-service/proto"
	"github.com/SleepingNext/product-service/repository"
)

type Handler struct {
	Repo repository.Repository
	ProductClient pb.ProductServiceClient
}

func (h *Handler) IndexProducts(ctx context.Context, req *pb.IndexProductsRequest, res *pb.Response) error {
	products, err := h.Repo.Index()
	if err != nil {
		return err
	}

	res.Products = products
	res.Error = nil

	return err
}

func (h *Handler) ShowProduct(ctx context.Context, req *pb.Product, res *pb.Response) error {
	product, err := h.Repo.Show(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error =   nil

	return nil
}

func (h *Handler) StoreProduct(ctx context.Context, req *pb.Product, res *pb.Response) error {
	product, err := h.Repo.Store(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error =   nil

	return err
}

func (h *Handler) UpdateProduct(ctx context.Context, req *pb.Product, res *pb.Response) error {
	product, err := h.Repo.Update(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error =   nil

	return nil
}

func (h *Handler) DestroyProduct(ctx context.Context, req *pb.Product, res *pb.Response) error {
	product, err := h.Repo.Destroy(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error =   nil

	return err
}
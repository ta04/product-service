// product-service/handler/handler.go

package handler

import (
	"context"

	productPB "github.com/SleepingNext/product-service/proto"
	productRepo "github.com/SleepingNext/product-service/repository"
)

type handler struct {
	repository productRepo.Repository
}

func NewHandler(repo productRepo.Repository) *handler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) IndexProducts(ctx context.Context, req *productPB.IndexProductsRequest, res *productPB.Response) error {
	products, err := h.repository.Index()
	if err != nil {
		return err
	}

	res.Products = products
	res.Error = nil

	return err
}

func (h *handler) ShowProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return nil
}

func (h *handler) StoreProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return err
}

func (h *handler) UpdateProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return nil
}

func (h *handler) DestroyProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return err
}

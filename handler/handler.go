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

package handler

import (
	"context"

	productPB "github.com/ta04/product-service/proto"
	productRepo "github.com/ta04/product-service/repository"
)

// Handler is the handler of order service
type Handler struct {
	repository productRepo.Repository
}

// NewHandler creates a new product service handler
func NewHandler(repo productRepo.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

// IndexProducts indexes the products
func (h *Handler) IndexProducts(ctx context.Context, req *productPB.IndexProductsRequest, res *productPB.Response) error {
	products, err := h.repository.Index(req)
	if err != nil {
		return err
	}

	res.Products = products
	res.Error = nil

	return err
}

// ShowProduct shows a product by ID
func (h *Handler) ShowProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return nil
}

// StoreProduct stores a new product
func (h *Handler) StoreProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return err
}

// UpdateProduct updates a product
func (h *Handler) UpdateProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return nil
}

// DestroyProduct destroys a product
func (h *Handler) DestroyProduct(ctx context.Context, req *productPB.Product, res *productPB.Response) error {
	product, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.Product = product
	res.Error = nil

	return err
}

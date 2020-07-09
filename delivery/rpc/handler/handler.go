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
	"errors"

	proto "github.com/ta04/product-service/model/proto"
	"github.com/ta04/product-service/usecase"
)

// Handler is the handler of product service
type Handler struct {
	Usecase usecase.Usecase
}

// NewHandler will create a new product handler
func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

// GetAllProducts will get all products
func (handler *Handler) GetAllProducts(ctx context.Context, req *proto.GetAllProductsRequest, res *proto.Response) error {
	products, err := handler.Usecase.GetAll(req)
	if err != nil {
		res.Products = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Products = products
	res.Error = nil

	return nil
}

// GetOneProduct will get a product
func (handler *Handler) GetOneProduct(ctx context.Context, req *proto.GetOneProductRequest, res *proto.Response) error {
	product, err := handler.Usecase.GetOne(req)
	if err != nil {
		res.Product = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Product = product
	res.Error = nil

	return nil
}

// CreateOneProduct will create a new product
func (handler *Handler) CreateOneProduct(ctx context.Context, req *proto.Product, res *proto.Response) error {
	product, err := handler.Usecase.CreateOne(req)
	if err != nil {
		res.Product = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Product = product
	res.Error = nil

	return nil
}

// UpdateOneProduct will update a product
func (handler *Handler) UpdateOneProduct(ctx context.Context, req *proto.Product, res *proto.Response) error {
	product, err := handler.Usecase.UpdateOne(req)
	if err != nil {
		res.Product = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Product = product
	res.Error = nil

	return nil
}

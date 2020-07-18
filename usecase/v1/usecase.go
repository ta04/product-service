package v1

import (
	"net/http"

	proto "github.com/ta04/product-service/model/proto"
	"github.com/ta04/product-service/repository"
)

// usecase is the struct of product usecase
type Usecase struct {
	Repository repository.Repository
}

// NewUsecase will create a new product usecase
func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (usecase *Usecase) GetAll(request *proto.GetAllProductsRequest) (*[]*proto.Product, *proto.Error) {
	if request == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	if request.Status == "" {
		request.Status = "active"
	}

	var products *[]*proto.Product
	var err error
	if request.Query != "" {
		products, err = usecase.Repository.GetAllByQuery(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		products, err = usecase.Repository.GetAll(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return products, nil
}

func (usecase *Usecase) GetOne(request *proto.GetOneProductRequest) (*proto.Product, *proto.Error) {
	if request == nil || (request.Id == 0) {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	product, err := usecase.Repository.GetOne(request)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return product, nil
}

func (usecase *Usecase) CreateOne(product *proto.Product) (*proto.Product, *proto.Error) {
	if product == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	product, err := usecase.Repository.CreateOne(product)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return product, nil
}

func (usecase *Usecase) UpdateOne(product *proto.Product) (*proto.Product, *proto.Error) {
	if product == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	product, err := usecase.Repository.UpdateOne(product)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return product, nil
}

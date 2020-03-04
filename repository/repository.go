// product-service/repository/repository.go

package repository

import (
	productPB "github.com/SleepingNext/product-service/proto"
)

type Repository interface {
	Index() ([]*productPB.Product, error)
	Show(*productPB.Product) (*productPB.Product, error)
	Store(*productPB.Product) (*productPB.Product, error)
	Update(*productPB.Product) (*productPB.Product, error)
	Destroy(*productPB.Product) (*productPB.Product, error)
}

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

package repository

import (
	proto "github.com/ta04/product-service/model/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	GetAllByQuery(request *proto.GetAllProductsRequest) (*[]*proto.Product, error)
	GetAll(request *proto.GetAllProductsRequest) (*[]*proto.Product, error)
	GetOne(request *proto.GetOneProductRequest) (*proto.Product, error)
	CreateOne(product *proto.Product) (*proto.Product, error)
	UpdateOne(product *proto.Product) (*proto.Product, error)
}

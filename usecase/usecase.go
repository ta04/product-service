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

package usecase

import proto "github.com/ta04/product-service/model/proto"

// usecase is the interface of usecases.
// As there are many version of usecases can be made.
type Usecase interface {
	GetAll(request *proto.GetAllProductsRequest) (*[]*proto.Product, *proto.Error)
	GetOne(request *proto.GetOneProductRequest) (*proto.Product, *proto.Error)
	CreateOne(user *proto.Product) (*proto.Product, *proto.Error)
	UpdateOne(user *proto.Product) (*proto.Product, *proto.Error)
}

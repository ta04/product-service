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
	productPB "github.com/SleepingNext/product-service/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Index() ([]*productPB.Product, error)
	Show(*productPB.Product) (*productPB.Product, error)
	Store(*productPB.Product) (*productPB.Product, error)
	Update(*productPB.Product) (*productPB.Product, error)
	Destroy(*productPB.Product) (*productPB.Product, error)
}

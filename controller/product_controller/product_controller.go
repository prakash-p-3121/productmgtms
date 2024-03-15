package product_controller

import "github.com/prakash-p-3121/restlib"

type ProductController interface {
	CreateProduct(ctx restlib.RestContext)
	FindProduct(ctx restlib.RestContext)
}

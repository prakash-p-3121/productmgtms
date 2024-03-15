package product_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func CreateProduct(c *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(c)
	controller := NewProductController()
	controller.CreateProduct(ginRestCtx)
}

func FindProduct(c *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(c)
	controller := NewProductController()
	controller.FindProduct(ginRestCtx)
}

package category_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func CreateCategory(c *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(c)
	controller := NewCategoryController()
	controller.CreateCategory(ginRestCtx)
}

func FindCategory(c *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(c)
	controller := NewCategoryController()
	controller.FindCategory(ginRestCtx)
}

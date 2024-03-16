package marketplace_listing_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func CreateMarketplaceListing(c *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(c)
	controller := NewMarketplaceListingController()
	controller.CreateMarketplaceListing(ginRestCtx)
}

func FindMarketplaceListingByID(c *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(c)
	controller := NewMarketplaceListingController()
	controller.FindMarketplaceListingByID(ginRestCtx)
}

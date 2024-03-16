package marketplace_listing_controller

import (
	"github.com/prakash-p-3121/productmgtms/controller/marketplace_listing_controller/impl"
	"github.com/prakash-p-3121/productmgtms/service/marketplace_listing_service"
)

func NewMarketplaceListingController() MarketPlaceListingController {
	service := marketplace_listing_service.NewMarketplaceListingService()
	return &impl.MarketplaceListingControllerImpl{MarketplaceListingService: service}
}

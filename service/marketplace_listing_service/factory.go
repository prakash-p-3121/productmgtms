package marketplace_listing_service

import (
	"github.com/prakash-p-3121/productmgtms/repository/marketplace_listing_repository"
	"github.com/prakash-p-3121/productmgtms/service/marketplace_listing_service/impl"
)

func NewMarketplaceListingService() MarketplaceListingService {
	repository := marketplace_listing_repository.NewMarketplaceListingRepository()
	return &impl.MarketplaceListingServiceImpl{MarketplaceListingRepository: repository}
}

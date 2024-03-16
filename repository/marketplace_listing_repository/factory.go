package marketplace_listing_repository

import (
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/productmgtms/repository/marketplace_listing_repository/impl"
)

func NewMarketplaceListingRepository() MarketplaceListingRepository {
	return &impl.MarketplaceListingRepositoryImpl{SingleStoreConnection: database.GetSingleStoreConnection()}
}

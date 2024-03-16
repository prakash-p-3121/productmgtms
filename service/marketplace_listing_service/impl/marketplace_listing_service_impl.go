package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenclient"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
	"github.com/prakash-p-3121/productmgtms/cfg"
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/productmgtms/repository/marketplace_listing_repository"
)

type MarketplaceListingServiceImpl struct {
	MarketplaceListingRepository marketplace_listing_repository.MarketplaceListingRepository
}

func (service *MarketplaceListingServiceImpl) CreateMarketplaceListing(
	req *productmgtmodel.MarketplaceListingCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	idGenMSCfg, err := cfg.GetMsConnectionCfg("idgenms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	idGenClient := idgenclient.NewIDGenClient(idGenMSCfg.Host, uint(idGenMSCfg.Port))
	idGenResp, appErr := idGenClient.NextID(database.MarketplaceListingsTable)
	if appErr != nil {
		return nil, appErr
	}

	appErr = service.MarketplaceListingRepository.CreateMarketplaceListing(idGenResp, req)
	if appErr != nil {
		return nil, appErr
	}
	return idGenResp, nil
}

func (service *MarketplaceListingServiceImpl) FindMarketplaceListingByID(listingID string) (*productmgtmodel.MarketplaceListing,
	errorlib.AppError) {
	marketplaceListing, appErr := service.MarketplaceListingRepository.FindMarketplaceListingByID(listingID)
	if appErr != nil {
		return nil, appErr
	}
	return marketplaceListing, nil
}

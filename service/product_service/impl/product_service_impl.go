package impl

import (
	database_clustermgt_client "github.com/prakash-p-3121/database-clustermgt-client"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenclient"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
	"github.com/prakash-p-3121/productmgtms/cfg"
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/productmgtms/repository/product_repository"
)

type ProductServiceImpl struct {
	ProductRepository product_repository.ProductRepository
}

func (service *ProductServiceImpl) CreateProduct(req *productmgtmodel.ProductCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	idGenMSCfg, err := cfg.GetMsConnectionCfg("idgenms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	idGenClient := idgenclient.NewIDGenClient(idGenMSCfg.Host, uint(idGenMSCfg.Port))
	resp, appErr := idGenClient.NextID(database.ProductsTable)
	if appErr != nil {
		return nil, appErr
	}
	productID := resp.ID

	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.ProductsTable, productID)
	if appErr != nil {
		return nil, appErr
	}

	appErr = service.ProductRepository.CreateProduct(*shardPtr.ID, resp, req)
	if appErr != nil {
		return nil, appErr
	}
	return resp, nil
}

func (service *ProductServiceImpl) FindProduct(productID string) (*productmgtmodel.Product, errorlib.AppError) {
	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.ProductsTable, productID)
	if appErr != nil {
		return nil, appErr
	}

	productPtr, appErr := service.ProductRepository.FindProduct(*shardPtr.ID, productID)
	if appErr != nil {
		return nil, appErr
	}
	return productPtr, nil
}

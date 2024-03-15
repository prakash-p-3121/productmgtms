package impl

import (
	database_clustermgt_client "github.com/prakash-p-3121/database-clustermgt-client"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenclient"
	"github.com/prakash-p-3121/productmgtmodel"
	"github.com/prakash-p-3121/productmgtms/cfg"
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/productmgtms/repository/category_repository"
	restlib_model "github.com/prakash-p-3121/restlib/model"
)

type CategoryServiceImpl struct {
	CategoryRepository category_repository.CategoryRepository
}

func (service *CategoryServiceImpl) CreateCategory(req *productmgtmodel.CategoryCreateReq) (*restlib_model.IDResponse, errorlib.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	idGenMSCfg, err := cfg.GetMsConnectionCfg("idgenms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	idGenClient := idgenclient.NewIDGenClient(idGenMSCfg.Host, uint(idGenMSCfg.Port))
	resp, appErr := idGenClient.NextID(database.CategoriesTable)
	if appErr != nil {
		return nil, appErr
	}
	categoryID := resp.ID

	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.CategoriesTable, categoryID)
	if appErr != nil {
		return nil, appErr
	}

	appErr = service.CategoryRepository.CreateCategory(*shardPtr.ID, resp, req)
	if appErr != nil {
		return nil, appErr
	}

	return &restlib_model.IDResponse{ResourceID: resp.ID}, nil
}

func (service *CategoryServiceImpl) FindCategory(categoryID string) (*productmgtmodel.Category, errorlib.AppError) {

	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.CategoriesTable, categoryID)
	if appErr != nil {
		return nil, appErr
	}

	categoryPtr, appErr := service.CategoryRepository.FindCategory(*shardPtr.ID, categoryID)
	if appErr != nil {
		return nil, appErr
	}
	return categoryPtr, nil
}

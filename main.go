package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/productmgtms/cfg"
	"github.com/prakash-p-3121/productmgtms/controller/category_controller"
	"github.com/prakash-p-3121/productmgtms/controller/product_controller"
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/restlib"
)

func main() {

	msConnectionsMap, err := restlib.CreateMsConnectionCfg("conf/microservice.toml")
	if err != nil {
		panic(err)
	}
	cfg.SetMsConnectionsMap(msConnectionsMap)

	databaseInst, err := mysqllib.CreateDatabaseConnectionWithRetryByCfg("conf/database.toml")
	if err != nil {
		panic(err)
	}
	database.SetSingleStoreConnection(databaseInst)

	hostPortCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		panic(err)
	}

	connectionsMap, err := mysqllib.CreateShardConnectionsWithRetry(database.GetShardedTableList(), hostPortCfg.Host, hostPortCfg.Port)
	if err != nil {
		panic(err)
	}
	database.SetShardConnectionsMap(connectionsMap)

	router := gin.Default()
	routerGroup := router.Group("/productmgtms")

	routerGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroup.POST("/v1/category", category_controller.CreateCategory)
	routerGroup.GET("/v1/category", category_controller.FindCategory)
	routerGroup.POST("/v1/product", product_controller.CreateProduct)
	routerGroup.GET("/v1/product", product_controller.FindProduct)

	err = router.Run("127.0.0.1:3003")
	if err != nil {
		panic("Error Starting ProductMgtMs")
	}

}

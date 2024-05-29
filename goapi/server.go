package main

import (
	"goapi/controller"
	"goapi/service"
	"log"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	// "goapi/middlewares"
	"goapi/config"
	"goapi/repositories"
)

func main() {
	mongoClient, err := config.ConnectToMongo()
	if err != nil {
		log.Fatal("Failed to connect to the mongoDB!")
	}

	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect from the MongoDB! Error: ", err)
		}
	}()

	db := mongoClient.Database("f1data") // can use .env HERE!!!!!

	driverDataRepo := repositories.NewDriverDataRepository(db)
	driverDataService := service.NewDriverDataService(driverDataRepo)
	driverDataController := controller.NewDriverDataController(driverDataService)


	seasonProfileRepo := repositories.NewSeasonProfileRepository(db)
	seasonProfileService := service.NewSeasonProfileService(seasonProfileRepo)
	seasonProfileController := controller.NewSeasonProfileController(seasonProfileService)

	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	server.Use(cors.New(corsConfig))
	//server := gin.New()

	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/availableSeasonProfiles", func(ctx *gin.Context) {
		ctx.JSON(200, seasonProfileController.GetAvailableSeasonProfiles())
	})

	server.GET("/seasonTeamDriverProfile/:season", func(ctx *gin.Context) {
		season := ctx.Param("season")
		ctx.JSON(200, seasonProfileController.GetSeasonProfile(season))
	})

	server.GET("/driverSeasonData/:season/:firstname/:surname", func(ctx *gin.Context) {
		season := ctx.Param("season")
		driver := ctx.Param("firstname") + " " + ctx.Param("surname")
		ctx.JSON(200, driverDataController.GetDriverSeasonData(driver, season))
	})

	server.GET("/driverRaceData/:driver", func(ctx *gin.Context) {
		driver := ctx.Param("driver")
		ctx.JSON(200, driverDataController.GetDriverRaceData(driver))
	})

	// server.GET("/driverComparionData/:driver1/:driver2", func(ctx *gin.Context) {
	//     driver1 := ctx.Param("driver1")
	//     driver2 := ctx.Param("driver2")
	//     ctx.JSON(200, driverDataController.GetDriverComparison(driver1, driver2))
	// })

	server.Run(":7999")
}

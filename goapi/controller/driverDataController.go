package controller

import (
    "goapi/entity"
    "goapi/service"
)

type DriverDataController interface {
    GetDriverSeasonData(driver string, season string) entity.DriverSeasonData
    GetDriverRaceData(driver string) entity.DriverRaceData
    // GetDriverComparison(driver1 string, driver2 string) []entity.DriverData
}

type driverDataController struct {
    driverDataService service.DriverDataService
}

func NewDriverDataController(driverDataService service.DriverDataService) DriverDataController {
    return &driverDataController{
        driverDataService: driverDataService,
    }
}

func (c *driverDataController) GetDriverSeasonData(driver string, season string) entity.DriverSeasonData {
    // Check for driver, if not found return 404
    return c.driverDataService.GetDriverSeasonData(driver, season)
}


func (c *driverDataController) GetDriverRaceData(driver string) entity.DriverRaceData {
    // Check for driver, if not found return 404
    return c.driverDataService.GetDriverRaceData(driver)
}


// func (c *driverDataController) GetDriverComparison(driver string) []entity.DriverData {
//     var info entity.DriverData
//     return entity.DriverData
// }

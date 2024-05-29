package controller

import (
    "goapi/entity"
    "goapi/service"
)

type SeasonProfileController interface {
    GetAvailableSeasonProfiles()    []string
    GetSeasonProfile(season string) entity.SeasonProfile
}

type seasonProfileController struct {
    seasonProfileService service.SeasonProfileService
}

func NewSeasonProfileController(seasonProfileService service.SeasonProfileService) SeasonProfileController {
    return &seasonProfileController{
        seasonProfileService: seasonProfileService,
    }
}

func (c* seasonProfileController) GetSeasonProfile(season string) entity.SeasonProfile {
    return c.seasonProfileService.GetSeasonProfile(season)
}

func (c* seasonProfileController) GetAvailableSeasonProfiles() []string {
    return c.seasonProfileService.GetAvailableSeasonProfiles()
}
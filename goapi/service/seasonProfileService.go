package service

import (
	"goapi/entity"
	"goapi/repositories"
)

// SeasonProfileService defines the interface for the service
type SeasonProfileService interface {
	GetAvailableSeasonProfiles() []string
	GetSeasonProfile(season string) entity.SeasonProfile
}

// seasonProfileService is the struct that implements the SeasonProfileService interface
type seasonProfileService struct {
	seasonProfileRepository repositories.SeasonProfileRepository
}

// NewSeasonProfileService creates a new instance of seasonProfileService
func NewSeasonProfileService(repo repositories.SeasonProfileRepository) SeasonProfileService {
	return &seasonProfileService{
		seasonProfileRepository: repo,
	}
}

func (s *seasonProfileService) GetSeasonProfile(season string) entity.SeasonProfile {
	seasonProfile := s.seasonProfileRepository.GetSeasonProfile(season)
	seasonProfile.NumRaces = len(seasonProfile.Races)
	return seasonProfile


}

func (s *seasonProfileService) GetAvailableSeasonProfiles() []string {
	return s.seasonProfileRepository.GetAvailableSeasonProfiles()
}


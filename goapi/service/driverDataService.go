package service

import (
	"goapi/entity"
	"goapi/repositories"
	"math"
	"fmt"
)

// DriverDataService defines the interface for the service
type DriverDataService interface {
	GetDriverSeasonData(driver string, season string) entity.DriverSeasonData
	GetDriverRaceData(driver string) entity.DriverRaceData
}

// driverDataService is the struct that implements the DriverDataService interface
type driverDataService struct {
	driverDataRepository repositories.DriverDataRepository
}

// NewDriverDataService creates a new instance of driverDataService
func NewDriverDataService(repo repositories.DriverDataRepository) DriverDataService {
	return &driverDataService{
		driverDataRepository: repo,
	}
}

// GetDriverSeasonData retrieves season data for a specific driver
func (s *driverDataService) GetDriverSeasonData(driver string, season string) entity.DriverSeasonData {

	var driverData entity.DriverSeasonData
	driverData.Profile.Name = driver

	// Retrieve RaceSeasonData from db
	driverData.Races = s.driverDataRepository.GetDriverSeasonRaceStats(driver, season)
	driverData.Races.Consistency = calculateResultsConsistency(driverData.Races.Results)
	driverData.Qualifyings = s.driverDataRepository.GetDriverSeasonQualyStats(driver, season)
	driverData.SprintRaces = s.driverDataRepository.GetDriverSeasonSprintRaceStats(driver, season)

	return driverData
}

// GetDriverRaceData retrieves race data for a specific driver
func (s *driverDataService) GetDriverRaceData(driver string) entity.DriverRaceData {
	var driverData entity.DriverRaceData
	// call func in repo
	driverData.Profile.Name = driver
	return driverData
}

func calculateDatasetMean(dataset []int) float64 {
	sum := 0.0
	for _, num := range dataset {
		sum += float64(num)
	}
	return sum / float64(len(dataset))
}

func calculateDatasetStdDev(dataset []int, mean float64) float64 {
	sum := 0.0
	for _, num := range dataset {
		sum += math.Pow(float64(num)-mean, 2)
	}
	variance := sum / float64(len(dataset))
	return math.Sqrt(variance)
}

func calculateResultsConsistency(results [] entity.DriverSeasonRaceResult) float64 {
	// We use coefficient of variance subtracted from 100% 
	// as a metric of consistency.

	// extract an array of result positions.
	var positions []int;
	for _, race := range results{
		if (!race.Dnf){
			positions = append(positions, race.Position)
		}
	}

	fmt.Println(positions)

	if (len(results) == 0 ){
		return 100.0
	}

	mean := calculateDatasetMean(positions)
	stdDev := calculateDatasetStdDev(positions, mean)
	coefficientVariance := (stdDev / mean) * 100
	consistencyPercentage := 100 - coefficientVariance

	return consistencyPercentage
}

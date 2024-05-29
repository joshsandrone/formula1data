package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"goapi/entity"
)

type DriverDataRepository interface {
	GetRaceWeekendsInSeasonCursor(season string, projection bson.M) (*mongo.Cursor, error)
	GetRacesDocumentsInSeason(season string) ([]entity.RaceDocument, error)
	GetQualyDocumentsInSeason(season string) ([]entity.QualyDocument, error)
	
	GetDriverSeasonRaceStats(name string, season string) (entity.DriverSeasonRaceData)
	GetDriverSeasonSprintRaceStats(name string, season string) (entity.DriverSeasonSprintRaceData)
	GetDriverSeasonQualyStats(name string, season string) (entity.DriverSeasonQualyData)
}

type driverDataRepository struct {
	seasonsCollection *mongo.Collection
	racesCollection   *mongo.Collection
}

func NewDriverDataRepository(db *mongo.Database) DriverDataRepository {
	return &driverDataRepository{
		seasonsCollection: db.Collection("seasons"),
		racesCollection:   db.Collection("races"),
	}
}


func (r *driverDataRepository) GetDriverSeasonRaceStats(name string, season string) (entity.DriverSeasonRaceData) {
	racesInSeason, err := r.GetRacesDocumentsInSeason(season)

	if err != nil {
		log.Fatal(err)
	}

	var driverSeasonData entity.DriverSeasonRaceData
	driverSeasonData.HighestRacePos = 100
	driverSeasonCumulativePos := 0
	driverSeasonCumulativePosExecludingDnfs := 0

	for _, raceWeekend := range racesInSeason {

		// Standard Races
		for _, racePosition := range raceWeekend.RaceResults {

			if racePosition.Driver == name {

				// Record the result
				raceResult := entity.DriverSeasonRaceResult{
					Location:     raceWeekend.Location,
					Position:     racePosition.Position,
					Dnf:          racePosition.Dnf,
					Win:          racePosition.Position < 2,
					Podium:       racePosition.Position < 4,
					PointsFinish: racePosition.Position < 11,
				}
				driverSeasonData.Results = append(driverSeasonData.Results, raceResult)

				// Highest finish position
				if racePosition.Position < driverSeasonData.HighestRacePos {
					driverSeasonData.HighestRacePos = racePosition.Position
				}

				// Number of wins
				if racePosition.Position == 1 {
					driverSeasonData.Wins++
				}

				// Number of podiums
				if racePosition.Position < 4 {
					driverSeasonData.Podiums++
				}

				// Number of points finishes
				if racePosition.Position < 11 {
					driverSeasonData.PointsFinishes++
				}

				// Number of DNFs
				if racePosition.Dnf {
					driverSeasonData.Dnfs++
				} else {
					driverSeasonCumulativePosExecludingDnfs += racePosition.Position
				}

				driverSeasonData.Points += racePosition.Points
				driverSeasonCumulativePos += racePosition.Position
			}
		}
	}

	driverSeasonData.AvgRacePos = float32(driverSeasonCumulativePos) / float32(len(racesInSeason))
	driverSeasonData.AvgRacePosExecludingDnfs = float32(driverSeasonCumulativePosExecludingDnfs) / float32(len(racesInSeason) - driverSeasonData.Dnfs)
	return driverSeasonData
}



func (r *driverDataRepository) GetDriverSeasonSprintRaceStats(name string, season string) (entity.DriverSeasonSprintRaceData) {
	racesInSeason, err := r.GetRacesDocumentsInSeason(season)
	totalSprintRaces := 0

	if err != nil {
		log.Fatal(err)
	}

	var driverSeasonData entity.DriverSeasonSprintRaceData

	driverSeasonData.HighestSprintRacePos = 100
	driverSeasonCumulativeSprintPos := 0
	driverSeasonCumulativePosExecludingDnfs := 0

	for _, raceWeekend := range racesInSeason {

		for _, sprintRacePosition := range raceWeekend.SprintRaceResults {

			if sprintRacePosition.Driver == name {

				totalSprintRaces += 1
				
				// Highest sprint race finish position
				if sprintRacePosition.Position < driverSeasonData.HighestSprintRacePos {
					driverSeasonData.HighestSprintRacePos = sprintRacePosition.Position
				}

				// Number of sprint wins
				if sprintRacePosition.Position == 1 { // sort out MAGIC NUMS
					driverSeasonData.SprintWins.Total++
					driverSeasonData.SprintWins.Locations = append(driverSeasonData.SprintWins.Locations, raceWeekend.Location)
				}

				// Number of sprint podiums
				if sprintRacePosition.Position < 4 { // sort out MAGIC NUMS
					driverSeasonData.SprintPodiums.Total++
					driverSeasonData.SprintPodiums.Locations = append(driverSeasonData.SprintPodiums.Locations, raceWeekend.Location)
				}

				// Number of sprint points finishes
				if sprintRacePosition.Position < 9 { // sort out MAGIC NUMS
					driverSeasonData.SprintPointsFinishes.Total++
					driverSeasonData.SprintPointsFinishes.Locations = append(driverSeasonData.SprintPointsFinishes.Locations, raceWeekend.Location)
				}

				// Number of DNFs
				if sprintRacePosition.Dnf {
					driverSeasonData.SprintDnfs.Total++
					driverSeasonData.SprintDnfs.Locations = append(driverSeasonData.SprintDnfs.Locations, raceWeekend.Location)
				} else {
					driverSeasonCumulativePosExecludingDnfs += sprintRacePosition.Position
				}


				driverSeasonData.SprintPoints += sprintRacePosition.Points
				driverSeasonCumulativeSprintPos += sprintRacePosition.Position
			}
		}
	}

	driverSeasonData.AvgSprintRacePos = float32(driverSeasonCumulativeSprintPos) / float32(totalSprintRaces)
	driverSeasonData.AvgSprintRacePosExecludingDnfs = float32(driverSeasonCumulativePosExecludingDnfs) / float32(totalSprintRaces - driverSeasonData.SprintDnfs.Total)
	return driverSeasonData
}


func (r *driverDataRepository) GetDriverSeasonQualyStats(name string, season string) (entity.DriverSeasonQualyData) {
	qualysInSeason, err := r.GetQualyDocumentsInSeason(season)

	if err != nil {
		log.Fatal(err)
	}

	var driverSeasonData entity.DriverSeasonQualyData
	driverSeasonData.HighestQualyPos = 100
	driverSeasonData.GapsToPole = make(map[string]float32)
	driverSeasonCumulativePos := 0

	for _, raceWeekend := range qualysInSeason {
		for _, qualyPosition := range raceWeekend.Results {

			if qualyPosition.Driver == name {

				var qualyResult entity.Result
				qualyResult.Location = raceWeekend.Location
				qualyResult.Round = 334 // NEEDS ADDRESSING
				qualyResult.Position = qualyPosition.Position
				qualyResult.Dnf = false
				driverSeasonData.Results = append(driverSeasonData.Results, qualyResult)

				// Highest finish position
				if qualyPosition.Position < driverSeasonData.HighestQualyPos {
					driverSeasonData.HighestQualyPos = qualyPosition.Position
				}

				// Number of poles
				if qualyPosition.Position == 1 {
					driverSeasonData.Poles.Total++
					driverSeasonData.Poles.Locations = append(driverSeasonData.Poles.Locations, raceWeekend.Location)
				}

				// Number of front rows
				if qualyPosition.Position < 3 {
					driverSeasonData.FrontRows.Total++
					driverSeasonData.FrontRows.Locations = append(driverSeasonData.FrontRows.Locations, raceWeekend.Location)
				}

				// Gap To Pole
				// ...... calculate it
				driverSeasonData.GapsToPole[raceWeekend.Location] = 0.512

				// Total finsh pos
				driverSeasonCumulativePos += qualyPosition.Position
			}
		}
	}

	driverSeasonData.AvgQualyPos = float32(driverSeasonCumulativePos) / float32(len(qualysInSeason))
	return driverSeasonData
}


func (r *driverDataRepository) GetRacesDocumentsInSeason(season string) ([]entity.RaceDocument, error) {
	// We are only interested in a race data so only project that
	projection := bson.M{"_id" : 0, "Race" : 1, "location" : 1, "season" : 1, "SprintRace" : 1}
	cursor, err := r.GetRaceWeekendsInSeasonCursor(season, projection)
	if err != nil {
		log.Fatal("Failed to get the cursor required from mongodb pipeline! Err: ", err)
	}
	defer cursor.Close(context.Background())

	var raceDocuments []entity.RaceDocument
	// Iterate through the cursor to process the documents
	for cursor.Next(context.Background()) {
		var raceDoc entity.RaceDocument
		if err := cursor.Decode(&raceDoc); err != nil {
			log.Fatal(err)
			return nil, err
		}
		raceDocuments = append(raceDocuments, raceDoc)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return raceDocuments, nil
}


func (r *driverDataRepository) GetQualyDocumentsInSeason(season string) ([]entity.QualyDocument, error) {
	// We are only interested in a race data so only project that
	projection := bson.M{"_id" : 0, "Qualifying" : 1, "location" : 1, "season" : 1}
	cursor, err := r.GetRaceWeekendsInSeasonCursor(season, projection)
	if err != nil {
		log.Fatal("Failed to get the cursor required from mongodb pipeline! Err: ", err)
	}
	defer cursor.Close(context.Background())

	var qualyDocuments []entity.QualyDocument
	// Iterate through the cursor to process the documents
	for cursor.Next(context.Background()) {
		var qualyDoc entity.QualyDocument
		if err := cursor.Decode(&qualyDoc); err != nil {
			log.Fatal(err)
			return nil, err
		}
		qualyDocuments = append(qualyDocuments, qualyDoc)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return qualyDocuments, nil
}


func (r *driverDataRepository) GetRaceWeekendsInSeasonCursor(season string, projection bson.M) (*mongo.Cursor, error) {
	filter := bson.M{"season": season}
	findOptions := options.Find().SetProjection(projection)

	cursor, err := r.racesCollection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

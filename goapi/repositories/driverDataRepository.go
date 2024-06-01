package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"goapi/entity"
	"fmt"
	"strconv"
	"strings"
	"time"
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
	driverSeasonData.PointsPerRace = float32(driverSeasonData.Points) / float32(len(racesInSeason))
	driverSeasonData.FinishRate = float32(len(racesInSeason) - driverSeasonData.Dnfs) / float32(len(racesInSeason))
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

				sprintRaceResult := entity.DriverSeasonRaceResult{
					Location:     raceWeekend.Location,
					Position:     sprintRacePosition.Position,
					Dnf:          sprintRacePosition.Dnf,
					Win:          sprintRacePosition.Position < 2,
					Podium:       sprintRacePosition.Position < 4,
					PointsFinish: sprintRacePosition.Position < 11,
				}
				driverSeasonData.Results = append(driverSeasonData.Results, sprintRaceResult)


				
				// Highest sprint race finish position
				if sprintRacePosition.Position < driverSeasonData.HighestSprintRacePos {
					driverSeasonData.HighestSprintRacePos = sprintRacePosition.Position
				}

				// Number of sprint wins
				if sprintRacePosition.Position == 1 { // sort out MAGIC NUMS
					driverSeasonData.SprintWins++
				}

				// Number of sprint podiums
				if sprintRacePosition.Position < 4 { // sort out MAGIC NUMS
					driverSeasonData.SprintPodiums++
				}

				// Number of sprint points finishes
				if sprintRacePosition.Position < 9 { // sort out MAGIC NUMS
					driverSeasonData.SprintPointsFinishes++
				}

				// Number of DNFs
				if sprintRacePosition.Dnf {
					driverSeasonData.SprintDnfs++
				} else {
					driverSeasonCumulativePosExecludingDnfs += sprintRacePosition.Position
				}

				driverSeasonData.SprintPoints += sprintRacePosition.Points
				driverSeasonCumulativeSprintPos += sprintRacePosition.Position
			}
		}
	}

	driverSeasonData.AvgSprintRacePos = float32(driverSeasonCumulativeSprintPos) / float32(totalSprintRaces)
	driverSeasonData.AvgSprintRacePosExecludingDnfs = float32(driverSeasonCumulativePosExecludingDnfs) / float32(totalSprintRaces - driverSeasonData.SprintDnfs)
	return driverSeasonData
}


func (r *driverDataRepository) GetDriverSeasonQualyStats(name string, season string) (entity.DriverSeasonQualyData) {
	qualysInSeason, err := r.GetQualyDocumentsInSeason(season)

	if err != nil {
		log.Fatal(err)
	}

	var driverSeasonData entity.DriverSeasonQualyData
	driverSeasonData.HighestQualyPos = 100
	driverSeasonCumulativePos := 0
	driverSeasonCumulativeGapToPole := 0.0

	var poleLapTime string

	for _, raceWeekend := range qualysInSeason {

		for index, qualyPosition := range raceWeekend.Results {

			if index == 0 {
				poleLapTime = qualyPosition.Q3
			}

			if qualyPosition.Driver == name {

				var qualyDelta time.Duration

				if qualyPosition.Q3 != "" {
					qualyDelta =  findLaptimeDifference(qualyPosition.Q3, poleLapTime)
				} else if qualyPosition.Q2 != ""{
					qualyDelta =  findLaptimeDifference(qualyPosition.Q2, poleLapTime)
				} else if qualyPosition.Q1 != ""{
					qualyDelta =  findLaptimeDifference(qualyPosition.Q1, poleLapTime)
				} else {
					qualyDelta = 0
				}

				driverSeasonCumulativeGapToPole += qualyDelta.Seconds()

				qualyResult := entity.DriverSeasonQualyResult{
					Location:     raceWeekend.Location,
					Position:     qualyPosition.Position,
					Pole:          qualyPosition.Position < 2,
					FrontRow:       qualyPosition.Position < 3,
					GapToPole: 	qualyDelta.Seconds(), // ...... Requires logic!!!!!
				}

				driverSeasonData.Results = append(driverSeasonData.Results, qualyResult)

				// Highest finish position
				if qualyPosition.Position < driverSeasonData.HighestQualyPos {
					driverSeasonData.HighestQualyPos = qualyPosition.Position
				}

				// Number of poles
				if qualyPosition.Position == 1 {
					driverSeasonData.Poles++
				}

				// Number of front rows
				if qualyPosition.Position < 3 {
					driverSeasonData.FrontRows++
				}

				// Total finsh pos
				driverSeasonCumulativePos += qualyPosition.Position
			}
		}
	}

	driverSeasonData.AvgQualyPos = float32(driverSeasonCumulativePos) / float32(len(qualysInSeason))
	driverSeasonData.AvgGapToPole = float32(driverSeasonCumulativeGapToPole) / float32(len(qualysInSeason))
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


// ParseLapTime parses a lap time string in the format "m:ss.SSS" into a time.Duration.
func parseLapTime(lapTimeStr string) (time.Duration, error) {
	parts := strings.Split(lapTimeStr, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid lap time format")
	}

	// Parse minutes
	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	// Parse seconds and milliseconds
	secondsParts := strings.Split(parts[1], ".")
	if len(secondsParts) != 2 {
		return 0, fmt.Errorf("invalid lap time format")
	}
	seconds, err := strconv.Atoi(secondsParts[0])
	if err != nil {
		return 0, err
	}
	milliseconds, err := strconv.Atoi(secondsParts[1])
	if err != nil {
		return 0, err
	}

	// Convert to time.Duration
	duration := time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond

	return duration, nil
}

// FindDifference returns the difference between two lap times.
func findLaptimeDifference(lapTime1 string, lapTime2 string) (time.Duration) {

	fmt.Println("laptime 1 = ", lapTime1)
	fmt.Println("laptime 2 = ", lapTime2)



	time1, err := parseLapTime(lapTime1)
	if err != nil {
		return 0
	}
	time2, err := parseLapTime(lapTime2)
	if err != nil {
		return 0
	}

	// Find the difference
	difference := time1 - time2

	return difference
}

package repositories

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"goapi/entity"
)

type SeasonProfileRepository interface {
	GetAvailableSeasonProfiles() []string
	GetSeasonProfile(season string) (entity.SeasonProfile)
}

type seasonProfileRepository struct {
	seasonsCollection *mongo.Collection
}

func NewSeasonProfileRepository(db *mongo.Database) SeasonProfileRepository {
	return &seasonProfileRepository{
		seasonsCollection: db.Collection("seasons"),
	}
}


func (r *seasonProfileRepository) GetAvailableSeasonProfiles() []string {
	filter := bson.M{}
	projection := bson.M{"_id" : 0, "season" : 1}
	findOptions := options.Find().SetProjection(projection)
	var seasons []string
	cursor, err := r.seasonsCollection.Find(context.Background(), filter, findOptions)

	if err != nil {
        return nil
    }

    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var seasonYear entity.SeasonYear
        if err := cursor.Decode(&seasonYear); err != nil {
            return nil
        }
        seasons = append(seasons, seasonYear.Year)
    }

	return seasons
}

func (r *seasonProfileRepository) GetSeasonProfile(season string) (entity.SeasonProfile) {
	filter := bson.M{"season": season}
	var seasonProfile entity.SeasonProfile // Use entity.SeasonProfile
	err := r.seasonsCollection.FindOne(context.Background(), filter).Decode(&seasonProfile)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching document found")
			return entity.SeasonProfile{} // Return an empty SeasonProfile
		}
		log.Fatal(err)
		return entity.SeasonProfile{}
	}
	return seasonProfile
}

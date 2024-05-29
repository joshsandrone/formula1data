package entity

type SeasonProfile struct {
    Season string         `json:"season" bson:"season"`
    Teams  []SeasonTeams `json:"teams" bson:"teams"`
}

type SeasonTeams struct {
    Name         string           `json:"name" bson:"name"`
    Drivers      []SeasonDrivers `json:"drivers" bson:"drivers"`
    PrimaryColor string           `json:"primaryColor" bson:"primaryColor"`
}

type SeasonDrivers struct {
    Name string `json:"name" bson:"name"`
}


type SeasonYear struct {
    Year string `bson:"season"`
}
package entity

type SeasonProfile struct {
    Season string         `json:"season" bson:"season"`
    Teams  []SeasonTeams `json:"teams" bson:"teams"`
    Races  []SeasonRaces `json:"races" bson:"races`
    NumRaces    int     `json:"numRaces"`
}

type SeasonTeams struct {
    Name         string           `json:"name" bson:"name"`
    Drivers      []SeasonDrivers `json:"drivers" bson:"drivers"`
    PrimaryColor string           `json:"primaryColor" bson:"primaryColor"`
    Points       int              `json:"points" bson:"points"  ` 
}

type SeasonDrivers struct {
    Name string `json:"name" bson:"name"`
}


type SeasonYear struct {
    Year string `bson:"season"`
}

type SeasonRaces struct {
    Round       int         `json:"round"`
    Location    string      `json:"location"`
    Date        string      `json:"date"`
    Laps        int         `json:"laps"`
}
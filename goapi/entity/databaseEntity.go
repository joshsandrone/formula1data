package entity



// Structs for Race
type FastestLap struct {
    TimeOfDay string `bson:"timeOfDay"`
    LapTime   string `bson:"lapTime"`
    LapNumber string `bson:"lapNumber"`
    AvgSpeed  string `bson:"avgSpeed"`
}

type RaceResult struct {
    Position   int        `bson:"Position"`
    Driver     string     `bson:"Driver"`
    Team       string     `bson:"Team"`
    FastestLap FastestLap     `bson:"FastestLap"`
    Points      int       `bson:"Points"` 
    Dnf         bool      `bson:"Dnf"`
}

type SprintRaceResult struct {
    Position   int        `bson:"Position"`
    Driver     string     `bson:"Driver"`
    Team       string     `bson:"Team"`
    Points      int       `bson:"Points"` 
    Dnf         bool      `bson:"Dnf"`
}

type RaceDocument struct {
    Location string `bson:"location"`
    Season   string `bson:"season"`
    RaceResults  []RaceResult `bson:"Race"`
    SprintRaceResults []SprintRaceResult `bson:"SprintRace"`
}


// Structs for Qualy
type QualyResult struct {
    Position   	int       	`bson:"Position"`
    Driver     	string     	`bson:"Driver"`
    Team       	string     	`bson:"Team"`
    Q1		  	string	   	`bson:"Q1 Time"`
	Q2			string		`bson:"Q2 Time"`
	Q3			string		`bson:"Q3 Time"`
}

type QualyDocument struct {
    Location string `bson:"location"`
    Season   string `bson:"season"`
    Results  []QualyResult `bson:"Qualifying"`
}










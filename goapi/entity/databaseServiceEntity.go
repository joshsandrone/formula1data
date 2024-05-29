package entity

// These structs are passed between the repositoires (database) layer and the service layer

type DriverSeasonRaceData struct {
	Wins 			int  	`json:"wins"` 
	Podiums 		int 	`json:"podiums"`
	Dnfs 			int 	`json:"dnfs"`
	Points 			int 	`json:"points"`
	AvgRacePos 		float32 `json:"avgRacePos"`
	AvgRacePosExecludingDnfs float32 `json:"avgRacePosExecludingDnfs`
	HighestRacePos 	int 	`json:"highestRacePos"` // Can add where
	//FastestLaps int `json:"fastestLaps"` // Swap to List?
	PointsFinishes 	int 	`json:"pointsFinishes"`
	Consistency  	float64	`json:"raceConsistency"`
	Results			[]DriverSeasonRaceResult	`json:"results"`
}

type DriverSeasonSprintRaceData struct {
	SprintWins Occurences `json:"sprintWins"`
	SprintPodiums Occurences `json:"sprintPodiums"`
	SprintDnfs Occurences `json:"sprintDnfs"`
	SprintPoints int `json:"sprintPoints"`
	AvgSprintRacePos float32 `json:"avgSprintRacePos"`
	AvgSprintRacePosExecludingDnfs float32 `json:"avgSprintRacePosExecludingDnfs`
	HighestSprintRacePos int `json:"highestSprintRacePos"`
	//SprintFastestLaps int `json:"sprintFastestLaps"`
	SprintPointsFinishes Occurences `json:"sprintRacePointsFinsishes"`
}


type DriverSeasonQualyData struct {
	Poles PoleOccurences `json:"poles"`// swap to List?
	FrontRows FrontRowOccurences `json:"frontRows"`// Swap to List?
	AvgQualyPos float32 `json:"avgQualyPos"`
	AvgGapToPole float32 `json:"avgGapToPole"`
	HighestQualyPos int `json:"highestQualyPos"`// Can add where
	GapsToPole map[string]float32 `json:"gapsToPole"`
	Results		[]Result	`json:"results"`
}

type DriverSeasonRaceResult struct {
	Win				bool		`json:"win"`
	Podium			bool		`json:"podium"`
	PointsFinish	bool		`json:"pointsFinish"`
	Dnf				bool		`json:"dnf`
	Location		string		`json:"location"`
	Position		int			`json:"position"`
}



type delta struct{

}


type Occurences struct {
	Total int `json:"total"`
	Locations []string `json:locations`
}

type Result struct {
	Round int `json:"round"`
	Location string `json:"location"`
	Position int `json:"position"`
	Dnf bool
}


// alias
type PoleOccurences = Occurences
type FrontRowOccurences = Occurences
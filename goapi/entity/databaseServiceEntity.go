package entity

// import "time"

// These structs are passed between the repositoires (database) layer and the service layer

type DriverSeasonRaceData struct {
	Wins 			int  	`json:"wins"` 
	Podiums 		int 	`json:"podiums"`
	Dnfs 			int 	`json:"dnfs"`
	FinishRate		float32	`json:"finishRate"`
	Points 			int 	 `json:"points"`
	PointsPerRace	float32	 `json:"pointsPerRace"`
	AvgRacePos 		float32 `json:"avgRacePos"`
	AvgRacePosExecludingDnfs float32 `json:"avgRacePosExecludingDnfs`
	HighestRacePos 	int 	`json:"highestRacePos"` // Can add where
	//FastestLaps int `json:"fastestLaps"` // Swap to List?
	PointsFinishes 	int 	`json:"pointsFinishes"`
	Consistency  	float64	`json:"raceConsistency"`
	Results			[]DriverSeasonRaceResult	`json:"results"`
}

type DriverSeasonSprintRaceData struct {
	SprintWins int `json:"sprintWins"`
	SprintPodiums int `json:"sprintPodiums"`
	SprintDnfs int `json:"sprintDnfs"`
	SprintPoints int `json:"sprintPoints"`
	AvgSprintRacePos float32 `json:"avgSprintRacePos"`
	AvgSprintRacePosExecludingDnfs float32 `json:"avgSprintRacePosExecludingDnfs`
	HighestSprintRacePos int `json:"highestSprintRacePos"`
	//SprintFastestLaps int `json:"sprintFastestLaps"`
	SprintPointsFinishes int `json:"sprintRacePointsFinsishes"`
	Results			[]DriverSeasonRaceResult	`json:"results"`
}


type DriverSeasonQualyData struct {
	Poles int `json:"poles"`// swap to List?
	FrontRows int `json:"frontRows"`// Swap to List?
	AvgQualyPos float32 `json:"avgQualyPos"`
	AvgGapToPole float32 `json:"avgGapToPole"`
	HighestQualyPos int `json:"highestQualyPos"`// Can add where
	Results		[]DriverSeasonQualyResult	`json:"results"`
}

type DriverSeasonRaceResult struct {
	Win				bool		`json:"win"`
	Podium			bool		`json:"podium"`
	PointsFinish	bool		`json:"pointsFinish"`
	Dnf				bool		`json:"dnf`
	Location		string		`json:"location"`
	Position		int			`json:"position"`
}

type DriverSeasonQualyResult struct {
	Pole				bool		`json:"pole"`
	FrontRow			bool		`json:"frontRow"`
	GapToPole			float64		`json:"gapToPole"`
	Location			string		`json:"location"`
	Position			int			`json:"position"`
}

type delta struct{

}

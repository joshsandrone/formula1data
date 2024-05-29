package entity

type DriverProfile struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Team string `json:"team"`    

}

type DriverSeasonData struct {
    Profile DriverProfile `json:"profile"`
    Races DriverSeasonRaceData `json:"races"`
    Qualifyings DriverSeasonQualyData `json:"qualifyings"`
    SprintRaces DriverSeasonSprintRaceData `json:"sprintRaces`
}

type DriverRaceData struct {
    Profile DriverProfile `json:"profile"`
}

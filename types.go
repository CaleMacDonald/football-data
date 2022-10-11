package footballdata

import (
	"time"
)

type AreaResponse struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Code         string      `json:"code"`
	Flag         interface{} `json:"flag"`
	ParentAreaID interface{} `json:"parentAreaId"`
	ParentArea   interface{} `json:"parentArea"`
	ChildAreas   []struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		CountryCode  string      `json:"countryCode"`
		Flag         interface{} `json:"flag"`
		ParentAreaID int         `json:"parentAreaId"`
		ParentArea   string      `json:"parentArea"`
	} `json:"childAreas"`
}

type AreasResponse struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Areas []struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		CountryCode  string      `json:"countryCode"`
		Flag         interface{} `json:"flag"`
		ParentAreaID int         `json:"parentAreaId"`
		ParentArea   string      `json:"parentArea"`
	} `json:"areas"`
}

type CompetitionResponse struct {
	Area struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Flag string `json:"flag"`
	} `json:"area"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Type          string `json:"type"`
	Emblem        string `json:"emblem"`
	CurrentSeason struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"currentSeason"`
	Seasons []struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"seasons"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type CompetitionsResponse struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competitions []struct {
		ID   int `json:"id"`
		Area struct {
			ID   int         `json:"id"`
			Name string      `json:"name"`
			Code string      `json:"code"`
			Flag interface{} `json:"flag"`
		} `json:"area"`
		Name          string      `json:"name"`
		Code          string      `json:"code"`
		Type          string      `json:"type"`
		Emblem        interface{} `json:"emblem"`
		Plan          string      `json:"plan"`
		CurrentSeason struct {
			ID              int         `json:"id"`
			StartDate       string      `json:"startDate"`
			EndDate         string      `json:"endDate"`
			CurrentMatchday int         `json:"currentMatchday"`
			Winner          interface{} `json:"winner"`
		} `json:"currentSeason"`
		NumberOfAvailableSeasons int       `json:"numberOfAvailableSeasons"`
		LastUpdated              time.Time `json:"lastUpdated"`
	} `json:"competitions"`
}

type StandingsResponse struct {
	Filters struct {
		Season string `json:"season"`
	} `json:"filters"`
	Area struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Flag string `json:"flag"`
	} `json:"area"`
	Competition struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Emblem string `json:"emblem"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	Standings []struct {
		Stage string      `json:"stage"`
		Type  string      `json:"type"`
		Group interface{} `json:"group"`
		Table []struct {
			Position int `json:"position"`
			Team     struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				ShortName string `json:"shortName"`
				Tla       string `json:"tla"`
				Crest     string `json:"crest"`
			} `json:"team"`
			PlayedGames    int    `json:"playedGames"`
			Form           string `json:"form"`
			Won            int    `json:"won"`
			Draw           int    `json:"draw"`
			Lost           int    `json:"lost"`
			Points         int    `json:"points"`
			GoalsFor       int    `json:"goalsFor"`
			GoalsAgainst   int    `json:"goalsAgainst"`
			GoalDifference int    `json:"goalDifference"`
		} `json:"table"`
	} `json:"standings"`
}

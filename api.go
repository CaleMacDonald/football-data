package footballdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	apiKey string
}

//goland:noinspection GoUnusedExportedFunction
func NewFootballData(apiKey string) *Api {
	return &Api{
		apiKey: apiKey,
	}
}

func (a *Api) Area(id int) (AreaResponse, error) {
	resp, err := a.makeRequest(fmt.Sprintf("https://api.football-data.org/v4/areas/%d", id))
	if err != nil {
		return AreaResponse{}, err
	}

	area := &AreaResponse{}
	err = json.Unmarshal(resp, area)
	return *area, err
}

func (a *Api) Areas() (AreasResponse, error) {
	resp, err := a.makeRequest("https://api.football-data.org/v4/areas")

	if err != nil {
		return AreasResponse{}, err
	}

	areas := &AreasResponse{}
	err = json.Unmarshal(resp, areas)
	return *areas, err
}

func (a *Api) Competition(league string) (CompetitionResponse, error) {
	resp, err := a.makeRequest(fmt.Sprintf("https://api.football-data.org/v4/competitions/%s", league))

	if err != nil {
		return CompetitionResponse{}, err
	}

	competition := &CompetitionResponse{}
	err = json.Unmarshal(resp, competition)
	return *competition, err
}

func (a *Api) Competitions() (CompetitionsResponse, error) {
	resp, err := a.makeRequest("https://api.football-data.org/v4/competitions")

	if err != nil {
		return CompetitionsResponse{}, err
	}

	competitions := &CompetitionsResponse{}
	err = json.Unmarshal(resp, competitions)
	return *competitions, err
}

func (a *Api) Standings(league string) (StandingsResponse, error) {
	resp, err := a.makeRequest(fmt.Sprintf("https://api.football-data.org/v4/competitions/%s/standings", league))

	if err != nil {
		return StandingsResponse{}, err
	}

	standings := &StandingsResponse{}
	err = json.Unmarshal(resp, standings)
	return *standings, err
}

func (a *Api) makeRequest(url string) ([]byte, error) {
	req, err := a.createRequest(url)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the request was not successful")
	}

	return io.ReadAll(resp.Body)
}

func (a *Api) createRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Auth-Token", a.apiKey)

	return req, nil
}

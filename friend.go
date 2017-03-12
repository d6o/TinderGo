package tindergo

import (
	"encoding/json"
	"errors"
)

type Friend struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Photo  []struct {
		ProcessedFiles []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"processedFiles"`
	} `json:"photo"`
	InSquad bool `json:"in_squad"`
}

type FriendResponse struct {
	Status  int      `json:"status"`
	Results []Friend `json:"results"`
}

func (t *TinderGo) Friends() ([]Friend, error) {
	fr := FriendResponse{}
	url := "https://api.gotinder.com/group/friends"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return fr.Results, errs[0]
	}

	err := json.Unmarshal([]byte(b), &fr)
	if err != nil {
		return fr.Results, err
	}

	if fr.Status != 200 {
		return fr.Results, errors.New("Error getting Facebook Friends.")
	}

	return fr.Results, nil
}

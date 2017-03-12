package tindergo

import (
	"encoding/json"
	"time"
)

type Profile struct {
	ID              string        `json:"_id"`
	AgeFilterMax    int           `json:"age_filter_max"`
	AgeFilterMin    int           `json:"age_filter_min"`
	Badges          []interface{} `json:"badges"`
	Bio             string        `json:"bio"`
	BirthDate       time.Time     `json:"birth_date"`
	Blend           string        `json:"blend"`
	CanCreateSquad  bool          `json:"can_create_squad"`
	ConnectionCount int           `json:"connection_count"`
	CreateDate      time.Time     `json:"create_date"`
	Discoverable    bool          `json:"discoverable"`
	DistanceFilter  int           `json:"distance_filter"`
	FacebookID      string        `json:"facebook_id"`
	Gender          int           `json:"gender"`
	GenderFilter    int           `json:"gender_filter"`
	InterestedIn    []int         `json:"interested_in"`
	Interests       []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"interests"`
	Jobs []struct {
		Title struct {
			Displayed bool   `json:"displayed"`
			Name      string `json:"name"`
			ID        string `json:"id"`
		} `json:"title"`
		Company struct {
			Displayed bool   `json:"displayed"`
			Name      string `json:"name"`
			ID        string `json:"id"`
		} `json:"company"`
	} `json:"jobs"`
	Location interface{} `json:"location"`
	Name     string      `json:"name"`
	Photos   []struct {
		YoffsetPercent   int    `json:"yoffset_percent,omitempty"`
		ID               string `json:"id"`
		XdistancePercent int    `json:"xdistance_percent,omitempty"`
		Main             bool   `json:"main,omitempty"`
		YdistancePercent int    `json:"ydistance_percent,omitempty"`
		XoffsetPercent   int    `json:"xoffset_percent,omitempty"`
		FileName         string `json:"fileName"`
		FbID             string `json:"fbId"`
		Extension        string `json:"extension"`
		ProcessedFiles   []struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"processedFiles"`
		URL         string  `json:"url"`
		SuccessRate float64 `json:"successRate"`
		SelectRate  int     `json:"selectRate"`
		Shape       string  `json:"shape,omitempty"`
	} `json:"photos"`
	PhotoOptimizerEnabled   bool      `json:"photo_optimizer_enabled"`
	PhotoOptimizerHasResult bool      `json:"photo_optimizer_has_result"`
	PingTime                time.Time `json:"ping_time"`
	Pos                     struct {
		At  int64   `json:"at"`
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"pos"`
	PosInfo struct {
		City struct {
			Name   string `json:"name"`
			Bounds struct {
				Ne struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"ne"`
				Sw struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"sw"`
			} `json:"bounds"`
		} `json:"city"`
		Country struct {
			Name   string `json:"name"`
			Bounds struct {
				Ne struct {
					Lat float64 `json:"lat"`
					Lng int     `json:"lng"`
				} `json:"ne"`
				Sw struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"sw"`
			} `json:"bounds"`
			Cc string `json:"cc"`
		} `json:"country"`
	} `json:"pos_info"`
	Schools            []interface{} `json:"schools"`
	SquadsOnly         bool          `json:"squads_only"`
	SquadsDiscoverable bool          `json:"squads_discoverable"`
	SquadAdsShown      bool          `json:"squad_ads_shown"`
	Username           string        `json:"username"`
}

func (t *TinderGo) Profile() (Profile, error) {
	pfl := Profile{}
	url := "https://api.gotinder.com/profile"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return pfl, errs[0]
	}

	err := json.Unmarshal([]byte(b), &pfl)
	if err != nil {
		return pfl, err
	}

	return pfl, nil
}

package tindergo

import (
	"encoding/json"
	"strings"
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
		YoffsetPercent   float64 `json:"yoffset_percent,omitempty"`
		ID               string  `json:"id"`
		XdistancePercent float64 `json:"xdistance_percent,omitempty"`
		Main             bool    `json:"main,omitempty"`
		YdistancePercent float64 `json:"ydistance_percent,omitempty"`
		XoffsetPercent   float64 `json:"xoffset_percent,omitempty"`
		FileName         string  `json:"fileName"`
		Extension        string  `json:"extension"`
		ProcessedFiles   []struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"processedFiles"`
		URL         string  `json:"url"`
		SuccessRate float64 `json:"successRate"`
		SelectRate  float64 `json:"selectRate"`
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
					Lng float64 `json:"lng"`
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

	b = strings.Replace(b, "\"main\",", "true, ", -1)

	err := json.Unmarshal([]byte(b), &pfl)
	if err != nil {
		return pfl, err
	}

	return pfl, nil
}

type ProfileUpdateResponse struct {
	ID                string        `json:"_id"`
	ActiveTime        time.Time     `json:"active_time"`
	AgeFilterMax      int           `json:"age_filter_max"`
	AgeFilterMin      int           `json:"age_filter_min"`
	APIToken          string        `json:"api_token"`
	Badges            []interface{} `json:"badges"`
	Bio               string        `json:"bio"`
	BirthDate         time.Time     `json:"birth_date"`
	Blend             string        `json:"blend"`
	CanCreateSquad    bool          `json:"can_create_squad"`
	College           []interface{} `json:"college"`
	ConnectionCount   int           `json:"connection_count"`
	CreateDate        time.Time     `json:"create_date"`
	Discoverable      bool          `json:"discoverable"`
	DistanceFilter    int           `json:"distance_filter"`
	DistanceFilterMin int           `json:"distance_filter_min"`
	Friends           []string      `json:"friends"`
	FullName          string        `json:"full_name"`
	Gender            int           `json:"gender"`
	GenderFilter      int           `json:"gender_filter"`
	Groups            []interface{} `json:"groups"`
	HighSchool        []interface{} `json:"high_school"`
	InterestedIn      []int         `json:"interested_in"`
	Interests         []struct {
		CreatedTime string `json:"created_time"`
		ID          string `json:"id"`
		Name        string `json:"name"`
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
	LatestUpdateDate time.Time   `json:"latest_update_date"`
	Location         interface{} `json:"location"`
	Name             string      `json:"name"`
	Photos           []struct {
		YoffsetPercent   int    `json:"yoffset_percent,omitempty"`
		ID               string `json:"id"`
		XdistancePercent int    `json:"xdistance_percent,omitempty"`
		Main             bool   `json:"main,omitempty"`
		YdistancePercent int    `json:"ydistance_percent,omitempty"`
		XoffsetPercent   int    `json:"xoffset_percent,omitempty"`
		FileName         string `json:"fileName"`
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
	PosMajor struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
		At  int64   `json:"at"`
	} `json:"pos_major"`
	PromotedOutOfDate  bool          `json:"promoted_out_of_date"`
	Schools            []interface{} `json:"schools"`
	SquadsOnly         bool          `json:"squads_only"`
	SquadsDiscoverable bool          `json:"squads_discoverable"`
	SquadAdsShown      bool          `json:"squad_ads_shown"`
	Username           string        `json:"username"`
}

type ProfileUpdateRequest struct {
	DistanceFilter int `json:"distance_filter"`
}

func (t *TinderGo) UpdateDistance(miles int) (ProfileUpdateResponse, error) {
	pfl := ProfileUpdateResponse{}
	req := ProfileUpdateRequest{DistanceFilter: miles}
	bReq, err := json.Marshal(req)
	if err != nil {
		return pfl, err
	}

	url := "https://api.gotinder.com/profile"
	b, errs := t.requester.Post(url, string(bReq))
	if errs != nil {
		return pfl, errs[0]
	}

	b = strings.Replace(b, "\"main\",", "true, ", -1)

	err = json.Unmarshal([]byte(b), &pfl)
	if err != nil {
		return pfl, err
	}

	return pfl, nil
}

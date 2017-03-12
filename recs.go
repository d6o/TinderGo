package tindergo

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type RecsCoreResponse struct {
	Status  int `json:"status"`
	Results []struct {
		Type         string `json:"type"`
		GroupMatched bool   `json:"group_matched"`
		User         struct {
			DistanceMi        int           `json:"distance_mi"`
			CommonConnections []interface{} `json:"common_connections"`
			ConnectionCount   int           `json:"connection_count"`
			CommonLikes       []interface{} `json:"common_likes"`
			CommonInterests   []interface{} `json:"common_interests"`
			CommonFriends     []interface{} `json:"common_friends"`
			ContentHash       string        `json:"content_hash"`
			ID                string        `json:"_id"`
			Badges            []interface{} `json:"badges"`
			Bio               string        `json:"bio"`
			BirthDate         time.Time     `json:"birth_date"`
			Name              string        `json:"name"`
			PingTime          time.Time     `json:"ping_time"`
			Photos            []struct {
				ID             string `json:"id"`
				URL            string `json:"url"`
				ProcessedFiles []struct {
					Width  int    `json:"width"`
					Height int    `json:"height"`
					URL    string `json:"url"`
				} `json:"processedFiles"`
			} `json:"photos"`
			Jobs    []interface{} `json:"jobs"`
			Schools []interface{} `json:"schools"`
			Teaser  struct {
				String string `json:"string"`
			} `json:"teaser"`
			Teasers       []interface{} `json:"teasers"`
			SNumber       int           `json:"s_number"`
			Gender        int           `json:"gender"`
			BirthDateInfo string        `json:"birth_date_info"`
			GroupMatched  bool          `json:"group_matched"`
		} `json:"user"`
	} `json:"results"`
}

func (t *TinderGo) RecsCore() (RecsCoreResponse, error) {
	var recs RecsCoreResponse
	url := "https://api.gotinder.com/recs/core"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return recs, errs[0]
	}

	err := json.Unmarshal([]byte(b), &recs)
	if err != nil {
		return recs, err
	}

	fmt.Println(recs)

	if recs.Status != 200 {
		return recs, errors.New("Error getting Recs Core")
	}

	return recs, nil
}

type RecsSocial struct {
	Type  string `json:"type"`
	Group struct {
		ID    string `json:"id"`
		Owner struct {
			DistanceMi        int `json:"distance_mi"`
			CommonConnections []struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Photo struct {
					Small  string `json:"small"`
					Medium string `json:"medium"`
					Large  string `json:"large"`
				} `json:"photo"`
				Degree int `json:"degree"`
			} `json:"common_connections"`
			ConnectionCount   int           `json:"connection_count"`
			CommonLikes       []interface{} `json:"common_likes"`
			CommonInterests   []interface{} `json:"common_interests"`
			UncommonInterests []interface{} `json:"uncommon_interests"`
			CommonFriends     []interface{} `json:"common_friends"`
			ID                string        `json:"_id"`
			Badges            []interface{} `json:"badges"`
			Bio               string        `json:"bio"`
			BirthDate         time.Time     `json:"birth_date"`
			Name              string        `json:"name"`
			PingTime          time.Time     `json:"ping_time"`
			Photos            []struct {
				ID             string `json:"id"`
				URL            string `json:"url"`
				ProcessedFiles []struct {
					URL    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"processedFiles"`
			} `json:"photos"`
			Jobs    []interface{} `json:"jobs"`
			Schools []interface{} `json:"schools"`
			Teaser  struct {
				String string `json:"string"`
			} `json:"teaser"`
			Teasers       []interface{} `json:"teasers"`
			Gender        int           `json:"gender"`
			BirthDateInfo string        `json:"birth_date_info"`
		} `json:"owner"`
		CreatedDate int64 `json:"created_date"`
		Members     []struct {
			DistanceMi        int           `json:"distance_mi"`
			CommonConnections []interface{} `json:"common_connections"`
			ConnectionCount   int           `json:"connection_count"`
			CommonLikes       []interface{} `json:"common_likes"`
			CommonInterests   []interface{} `json:"common_interests"`
			CommonFriends     []interface{} `json:"common_friends"`
			ID                string        `json:"_id"`
			Badges            []interface{} `json:"badges"`
			Bio               string        `json:"bio"`
			BirthDate         time.Time     `json:"birth_date"`
			Name              string        `json:"name"`
			PingTime          time.Time     `json:"ping_time"`
			Photos            []struct {
				ID             string `json:"id"`
				URL            string `json:"url"`
				ProcessedFiles []struct {
					URL    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"processedFiles"`
			} `json:"photos"`
			Jobs    []interface{} `json:"jobs"`
			Schools []interface{} `json:"schools"`
			Teaser  struct {
				String string `json:"string"`
			} `json:"teaser"`
			Teasers       []interface{} `json:"teasers"`
			Gender        int           `json:"gender"`
			BirthDateInfo string        `json:"birth_date_info"`
		} `json:"members"`
		Expired bool   `json:"expired"`
		Status  string `json:"status"`
	} `json:"group"`
	GroupMatched bool `json:"group_matched"`
}

type RecsSocialResponse struct {
	Status  int          `json:"status"`
	Results []RecsSocial `json:"results"`
}

func (t *TinderGo) RecsSocial() ([]RecsSocial, error) {
	recs := RecsSocialResponse{}
	url := "https://api.gotinder.com/recs/social"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return recs.Results, errs[0]
	}

	err := json.Unmarshal([]byte(b), &recs)
	if err != nil {
		return recs.Results, err
	}

	if recs.Status != 200 {
		return recs.Results, errors.New("Error getting Recs Social.")
	}

	return recs.Results, nil
}

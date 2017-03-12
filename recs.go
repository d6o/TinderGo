package tindergo

import (
	"encoding/json"
	"errors"
	"time"
)

type RecsCoreUser struct {
	DistanceMi      int           `json:"distance_mi"`
	ConnectionCount int           `json:"connection_count"`
	CommonLikes     []interface{} `json:"common_likes"`
	CommonFriends   []interface{} `json:"common_friends"`
	ContentHash     string        `json:"content_hash"`
	ID              string        `json:"_id"`
	Badges          []interface{} `json:"badges,omitempty"`
	Bio             string        `json:"bio"`
	BirthDate       time.Time     `json:"birth_date"`
	Name            string        `json:"name"`
	PingTime        time.Time     `json:"ping_time"`
	Photos          []struct {
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
	Teasers           []interface{} `json:"teasers"`
	SNumber           int           `json:"s_number"`
	SpotifyThemeTrack interface{}   `json:"spotify_theme_track,omitempty"`
	Gender            int           `json:"gender"`
	BirthDateInfo     string        `json:"birth_date_info"`
	Instagram         struct {
		LastFetchTime         time.Time `json:"last_fetch_time"`
		CompletedInitialFetch bool      `json:"completed_initial_fetch"`
		Photos                []struct {
			Image     string `json:"image"`
			Thumbnail string `json:"thumbnail"`
			Ts        string `json:"ts"`
			Link      string `json:"link"`
		} `json:"photos"`
		MediaCount     int    `json:"media_count"`
		ProfilePicture string `json:"profile_picture"`
		Username       string `json:"username"`
	} `json:"instagram,omitempty"`
	CommonConnections []interface{} `json:"common_connections,omitempty"`
	CommonInterests   []interface{} `json:"common_interests,omitempty"`
	UncommonInterests []interface{} `json:"uncommon_interests,omitempty"`
	SpotifyTopArtists []struct {
		Selected bool   `json:"selected"`
		Name     string `json:"name"`
		ID       string `json:"id"`
		TopTrack struct {
			URI   string `json:"uri"`
			Album struct {
				Images []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"images"`
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"album"`
			Name       string `json:"name"`
			PreviewURL string `json:"preview_url"`
			Artists    []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"artists"`
			ID string `json:"id"`
		} `json:"top_track"`
	} `json:"spotify_top_artists,omitempty"`
	IsTraveling bool `json:"is_traveling,omitempty"`
}

type RecsCoreResponse struct {
	Status  int            `json:"status"`
	Results []RecsCoreUser `json:"results"`
}

func (t *TinderGo) RecsCore() ([]RecsCoreUser, error) {
	var recs RecsCoreResponse
	url := "https://api.gotinder.com/recs/core"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return recs.Results, errs[0]
	}

	err := json.Unmarshal([]byte(b), &recs)
	if err != nil {
		return recs.Results, err
	}

	if recs.Status != 200 {
		return recs.Results, errors.New("Error getting Recs Core")
	}

	return recs.Results, nil
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

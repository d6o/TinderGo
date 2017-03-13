package tindergo

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type User struct {
	ConnectionCount int           `json:"connection_count"`
	CommonLikes     []interface{} `json:"common_likes"`
	CommonFriends   []interface{} `json:"common_friends"`
	UID             string        `json:"_id"`
	Bio             string        `json:"bio"`
	BirthDate       time.Time     `json:"birth_date"`
	Name            string        `json:"name"`
	PingTime        time.Time     `json:"ping_time"`
	Photos          []UserPhoto   `json:"photos"`
	Jobs            []interface{} `json:"jobs"`
	Schools         []interface{} `json:"schools"`
	Teasers         []interface{} `json:"teasers"`
	Gender          int           `json:"gender"`
	BirthDateInfo   string        `json:"birth_date_info"`
	ID              string        `json:"id"`
	DistanceMi      int           `json:"distance_mi"`
}

type UserPhoto struct {
	URL            string `json:"url"`
	ProcessedFiles []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"processedFiles"`
	Extension string `json:"extension"`
	FileName  string `json:"fileName"`
	Shape     string `json:"shape"`
	Main      bool   `json:"main,omitempty"`
	ID        string `json:"id"`
}

type UserResponse struct {
	Status  int  `json:"status"`
	Results User `json:"results"`
}

func (t *TinderGo) User(userID string) (User, error) {
	user := UserResponse{}
	url := "https://api.gotinder.com/user/" + userID
	b, errs := t.requester.Get(url)
	if errs != nil {
		return user.Results, errs[0]
	}

	b = strings.Replace(b, "\"main\",", "true, ", -1)

	err := json.Unmarshal([]byte(b), &user)
	if err != nil {
		return user.Results, err
	}

	if user.Status != 200 {
		return user.Results, errors.New("Error getting user information.")
	}

	return user.Results, nil
}

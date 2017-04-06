package tindergo

import (
	"encoding/json"
	"time"
)

type Match struct {
	ID                string        `json:"_id"`
	Closed            bool          `json:"closed"`
	CommonFriendCount int           `json:"common_friend_count"`
	CommonLikeCount   int           `json:"common_like_count"`
	CreatedDate       time.Time     `json:"created_date"`
	Dead              bool          `json:"dead"`
	LastActivityDate  time.Time     `json:"last_activity_date"`
	MessageCount      int           `json:"message_count"`
	Messages          []interface{} `json:"messages"`
	Participants      []string      `json:"participants"`
	Pending           bool          `json:"pending"`
	NotFollowing      interface{}   `json:"not_following"`
	IsSuperLike       bool          `json:"is_super_like"`
	IsBoostMatch      bool          `json:"is_boost_match"`
	Person            struct {
		ID        string        `json:"_id"`
		Badges    []interface{} `json:"badges"`
		Bio       string        `json:"bio"`
		BirthDate time.Time     `json:"birth_date"`
		Gender    int           `json:"gender"`
		Name      string        `json:"name"`
		PingTime  time.Time     `json:"ping_time"`
		Photos    []struct {
			SelectRate     float64 `json:"selectRate"`
			SuccessRate    float64 `json:"successRate"`
			FileName       string  `json:"fileName"`
			ID             string  `json:"id"`
			Extension      string  `json:"extension"`
			ProcessedFiles []struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"processedFiles"`
			URL string `json:"url"`
		} `json:"photos"`
		Teasers []interface{} `json:"teasers"`
	} `json:"person"`
}

type MatchesResponse struct {
	Matches          []Match       `json:"matches"`
	Blocks           []string      `json:"blocks"`
	Lists            []interface{} `json:"lists"`
	DeletedLists     []interface{} `json:"deleted_lists"`
	LastActivityDate time.Time     `json:"last_activity_date"`
}

type MatchesRequest struct {
	LastActivityDate string `json:"last_activity_date"`
}

func (t *TinderGo) Matches() ([]Match, error) {
	mr := MatchesResponse{}
	req := MatchesRequest{}
	bReq, err := json.Marshal(req)
	if err != nil {
		return mr.Matches, err
	}

	url := "https://api.gotinder.com/updates"
	b, errs := t.requester.Post(url, string(bReq))
	if errs != nil {
		return mr.Matches, errs[0]
	}

	err = json.Unmarshal([]byte(b), &mr)
	if err != nil {
		return mr.Matches, err
	}

	return mr.Matches, nil
}

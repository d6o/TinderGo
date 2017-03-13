package tindergo

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LikeResponse struct {
	Match          bool `json:"match,omitempty"`
	LikesRemaining int  `json:"likes_remaining,omitempty"`
	Status         int  `json:"status,omitempty"`
}

func (t *TinderGo) Decide(user RecsCoreUser, action string) (LikeResponse, error) {
	like := LikeResponse{}
	url := "https://api.gotinder.com/%s/%s?content_hash=%s&s_number=%s"
	url = fmt.Sprintf(url, action, user.ID, user.ContentHash, user.SNumber)
	b, errs := t.requester.Get(url)
	if errs != nil {
		return like, errs[0]
	}

	err := json.Unmarshal([]byte(b), &like)
	if err != nil {
		return like, err
	}

	if like.Status != 200 && like.Status != 0 {
		return like, errors.New("Error saving like.")
	}

	return like, nil
}

func (t *TinderGo) Like(user RecsCoreUser) (LikeResponse, error) {
	return t.Decide(user, "like")
}

func (t *TinderGo) Pass(user RecsCoreUser) (LikeResponse, error) {
	return t.Decide(user, "pass")
}

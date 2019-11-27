package tindergo

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LikeResponse struct {
	Match          interface{} `json:"match,omitempty"`
	LikesRemaining int         `json:"likes_remaining,omitempty"`
	Status         int         `json:"status,omitempty"`
}

type SuperLikeResponse struct {
	LimitExceeded bool `json:"limit_exceeded"`
	Status        int  `json:"status"`
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

func (t *TinderGo) SuperLike(userID string, s_number string) (SuperLikeResponse, error) {
	like := SuperLikeResponse{}
	url := fmt.Sprintf("https://api.gotinder.com/like/%s/super?locale=ru", userID)
	var body string
	if s_number != "" {
		body = fmt.Sprintf(`{"s_number":"%s"}`, s_number)
	}
	b, errs := t.requester.Post(url, body)
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

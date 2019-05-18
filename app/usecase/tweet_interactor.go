package usecase

import "../../app/domain"

type TweetInteractor struct {
}

func (interactor *TweetInteractor) Tweets() (domain.Tweet, error) {
	var tweet domain.Tweet
	var err error
	return tweet, err
}

func (interactor *TweetInteractor) TweetById(id int) (domain.Tweet, error) {
	var tweet domain.Tweet
	var err error
	return tweet, err
}

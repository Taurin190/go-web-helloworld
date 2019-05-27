package usecase

import "../../app/domain"

type TweetRepository interface {
	FindById(int) (domain.Tweet, error)
	FindAll() (domain.Tweet, error)
}

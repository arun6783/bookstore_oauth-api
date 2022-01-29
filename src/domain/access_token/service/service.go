package service

import (
	"github.com/arun6783/bookstore_oauth-api/domain/access_token"
	"github.com/arun6783/bookstore_oauth-api/repository/db"
	"github.com/arun6783/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type service struct {
	repository db.DbRepository
}

func NewService(repo db.DbRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

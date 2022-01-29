package db

import (
	"github.com/arun6783/bookstore_oauth-api/domain/access_token"
	"github.com/arun6783/bookstore_oauth-api/utils/errors"
)

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Not implemented database coonnection")
}

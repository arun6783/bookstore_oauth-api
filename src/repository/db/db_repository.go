package db

import (
	"github.com/arun6783/bookstore_oauth-api/clients/cassandra"
	"github.com/arun6783/bookstore_oauth-api/domain/access_token"
	"github.com/arun6783/bookstore_oauth-api/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "Select access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateAccessToken = "Update access_tokens SET expires=? WHERE access_token=?;"
)

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

func (r *dbRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateAccessToken, accessToken.Expires, accessToken.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken, accessToken.AccessToken, accessToken.UserId, accessToken.ClientId, accessToken.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, accessTokenId).
		Scan(
			&result.AccessToken,
			&result.UserId,
			&result.ClientId,
			&result.Expires,
		); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError("db error occured when trying to get data")
	}

	return &result, nil
}

package http

import (
	"net/http"
	"strings"

	"github.com/arun6783/bookstore_oauth-api/domain/access_token/service"
	"github.com/gin-gonic/gin"
)

func NewHandler(service service.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service service.Service
}

func (h *accessTokenHandler) GetById(c *gin.Context) {

	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(int(err.Status), err)
	}

	c.JSON(http.StatusOK, accessToken.AccessToken)
}

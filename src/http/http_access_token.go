package http

import (
	"net/http"
	"strings"

	"github.com/arun6783/bookstore_oauth-api/domain/access_token"
	"github.com/arun6783/bookstore_oauth-api/domain/access_token/service"
	"github.com/arun6783/bookstore_oauth-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func NewHandler(service service.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service service.Service
}

func (h *accessTokenHandler) GetById(c *gin.Context) {

	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}

	c.JSON(http.StatusOK, accessToken.AccessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBind(&at); err != nil {
		var err = errors.NewBadResuestError("invalid json body")
		c.JSON(int(err.Status), err)
		return
	}

	err := h.service.Create(at)
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}

	c.JSON(http.StatusOK, at)
}

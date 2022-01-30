package app

import (
	"github.com/arun6783/bookstore_oauth-api/domain/access_token/service"
	"github.com/arun6783/bookstore_oauth-api/http"
	"github.com/arun6783/bookstore_oauth-api/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	dbRepository := db.NewRepository()

	atService := service.NewService(dbRepository)

	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")

}

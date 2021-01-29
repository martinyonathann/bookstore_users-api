package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinyonathann/bookstore_users-api/services"
	"github.com/martinyonathann/bookstore_users-api/utils/curl"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
)

const (
	urlCreateItem = "http://localhost:8081/items"
)

func CreateItem(c *gin.Context) {
	email, password, _ := c.Request.BasicAuth()

	countLogin, countErr := services.UsersService.LoginUser(email, password)

	if countErr != nil {
		c.JSON(http.StatusOK, countErr.Error)
		return
	}
	if countLogin < 1 {
		c.JSON(http.StatusOK, errors.NewNotFoundError("Account Not Found"))
		return
	}

	itemsDomain, err := curl.RequestToGateway(c.Request.Method, urlCreateItem, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, itemsDomain)
	return

	// logger.RequestLog("Request to "+url, zap.Any("data_request", itemsDomain))
}

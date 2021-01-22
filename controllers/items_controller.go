package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/martinyonathann/bookstore_users-api/utils/curl"
)

const (
	urlCreateItem = "http://localhost:8081/items"
)

func CreateItem(c *gin.Context) {

	// reqBody, _ := (ioutil.ReadAll(c.Request.Body))
	// json.Unmarshal([]byte(reqBody), &itemsDomain)

	// logger.RequestLog("Request", zap.Any("data_request", itemsDomain))

	itemsDomain, _ := curl.RequestToGateway(c.Request.Method, urlCreateItem, c.Request.Body)

	c.JSON(200, itemsDomain)

	// logger.RequestLog("Request to "+url, zap.Any("data_request", itemsDomain))
}

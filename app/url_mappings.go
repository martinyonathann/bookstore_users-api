package app

import (
	"github.com/martinyonathann/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.GetUser)

	router.POST("/users", controllers.CreateUser)
}

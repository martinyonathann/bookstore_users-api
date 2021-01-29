package app

import (
	"github.com/martinyonathann/bookstore_users-api/controllers"
)

func mapUrls() {
	//User API
	router.GET("/users/:user_id", controllers.Get)
	router.GET("/internal/users/search", controllers.Search)
	router.POST("/login", controllers.Login)
	router.POST("/users", controllers.Create)
	router.PUT("/users/:user_id", controllers.Update)
	router.PATCH("/users/:user_id", controllers.Update)
	router.DELETE("/users/:user_id", controllers.Delete)

	//Item API
	router.POST("items", controllers.CreateItem)
}

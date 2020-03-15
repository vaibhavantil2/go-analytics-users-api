package app

import (
	ping "github.com/nanoTitan/analytics-users-api/controllers/ping"
	users "github.com/nanoTitan/analytics-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users/:user_id", users.CreateUser)
	router.GET("/users/search/:user_id", users.SearchUser)
	router.GET("/users", users.GetUser)
}

package main

import (
	"myblog/app"
	"myblog/config"
	"myblog/router"
)

func main() {
	config.LoadConfig()
	app.SetupDB()
	port := config.Viper.GetString("server.port")
	//mode := config.Viper.GetString("server.mode")
	Router := router.NewRouter()
	Router.Run(":" + port)
}

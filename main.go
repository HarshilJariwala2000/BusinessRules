package main

import (
	"calculationengine/constants"
	"calculationengine/store"
	"calculationengine/router"
	
)

func main(){
	constants.Load()
	storage.Connect()
	// storage.AutoMigrate()

	router.Api()
	router.Router.Run("0.0.0.0:3000")

}
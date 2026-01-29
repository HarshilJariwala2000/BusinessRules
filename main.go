package main

import (
	"calculationengine/constants"
	"calculationengine/store"
	"calculationengine/router"
	"text/scanner"
	"strings"
	"fmt"
)

func main(){
	constants.Load()
	storage.Connect()
	// storage.AutoMigrate()
	var s scanner.Scanner
	s.Init(strings.NewReader(`IF map, mrp, ^ "Hello"`))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Println(scanner.TokenString(tok))
	}

	router.Api()
	router.Router.Run("0.0.0.0:3000")

}
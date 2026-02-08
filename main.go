package main

import (
	"calculationengine/constants"
	// "calculationengine/router"
	"calculationengine/service/interpreter"
	// "calculationengine/store"
	"fmt"
	    "encoding/json"

	// "strings"
	// "text/scanner"
)

func main(){
	constants.Load()
	// storage.Connect()
	// storage.AutoMigrate()
	// var s scanner.Scanner
	// s.Init(strings.NewReader(`IF map, mrp, ^ <> "Hello"`))
	// for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
	// 	fmt.Println(scanner.TokenString(tok))
	// }

	lexer := interpreter.NewLexer("IF(map = mrp * tax / 100, IF(A = B, C, D), b)")
	parser := interpreter.NewParser(lexer)
	program := parser.ParseProgram()
	fmt.Println(program)
	// fmt.Printf("%+v\n", program)
	b, err := json.MarshalIndent(program, "", "  ")
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Println(string(b))

	// actual := program.String()

	// fmt.Println(actual)
	// fmt.Println("'+'")
	// router.Api()
	// router.Router.Run("0.0.0.0:3000")

}
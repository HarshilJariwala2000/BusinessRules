package main

import (
	"calculationengine/constants"
	// "calculationengine/router"
	"calculationengine/service/evaluator"
	"calculationengine/service/parser"
	// "calculationengine/store"
	// "encoding/json"
	"fmt"
	"time"
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
	start := time.Now()
	lexer := parser.NewLexer("100 + 102 + 103")
	nparser := parser.NewParser(lexer)
	program := nparser.ParseProgram()
	eval := evaluator.Eval(program.Statements[0])
	fmt.Println(eval)

	allTime := time.Now()
	fmt.Printf("Total took: %s\n", allTime.Sub(start))
	// fmt.Println(program)
	// fmt.Printf("%+v\n", program)

	// actual := program.String()

	// fmt.Println(actual)
	// fmt.Println("'+'")
	// router.Api()
	// router.Router.Run("0.0.0.0:3000")

}
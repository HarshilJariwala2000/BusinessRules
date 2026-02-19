package evaluator

import (
	"calculationengine/service/parser"
	"fmt"
)

var (
	TRUE = &Boolean{Value:true}
	FALSE = &Boolean{Value:false}
)

func Eval(node parser.Node) Object {
	// fmt.Println()
	// fmt.Printf("The type is: %T\n", node)
	switch node := node.(type) {
	// case *parser.Program:
	// 	return evalStatements(parser.Statement)

		case *parser.ExpressionStatement:
			return Eval(node.Expression)

		case *parser.PrefixExpression:
			right := Eval(node.Right)
			return evalPrefixExpression(node.Operator, right)

		case *parser.IntegerLiteral:
			return &Integer{Value:int64(node.Value)}

		case *parser.Identifier:


		case *parser.Boolean:
			return nativeBoolToBooleanObject(node.Value)

		case *parser.InfixExpression:
			left := Eval(node.Left)
			right := Eval(node.Right)
			valuess:=evalInfixExpression(node.Operator, left, right)
			fmt.Println(valuess)
			return valuess//evalInfixExpression(node.Operator, left, right)

		case *parser.IfExpression:
			return evalIfExpression(node)
			
		
	}


	return nil
}

func evalIfExpression(ie *parser.IfExpression) Object {
	condition := Eval(ie.Condition)
	if(isTruthy(condition)){
		return Eval(ie.Consequence)
	}else if(ie.Alternative!=nil){
		return Eval(ie.Alternative)
	}else{
		return nil
	}
}

func isTruthy(object Object) bool {
	switch {
	case object == TRUE:
		return true
	case object == FALSE:
		return false
	case object.Type()==INTEGER_OBJ && object.(*Integer).Value==0:
		return false
	default:
		return true
	}
}

func evalInfixExpression(operator string, left Object, right Object) Object{
	switch {
		case left.Type()==INTEGER_OBJ && right.Type()==INTEGER_OBJ:
			return evalIntegerInfixExpression(operator, left, right)
		/*
			We’re using pointer comparison here to check for equality between booleans. That works
			because we’re always using pointers to our objects and in the case of booleans we only ever use
			two: TRUE and FALSE. So, if something has the same value as TRUE (the memory address that is)
			then it’s true.
		*/
		case operator=="=":
			return nativeBoolToBooleanObject(left == right)
		case operator=="<>":
			return nativeBoolToBooleanObject(left != right)
		case left.Type() != right.Type():
			return newError("Type mismatch: %s %s %s", left.Type(), operator, right.Type())
		default:
			return newError("Unknown Operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left Object, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
		case "+":
			return &Integer{Value:(leftVal+rightVal)}
		case "-":
			return &Integer{Value:(leftVal - rightVal)}
		case "*":
			return &Integer{Value:(leftVal * rightVal)}
		case "/":
			return &Integer{Value:(leftVal / rightVal)}
		case ">":
			return nativeBoolToBooleanObject(leftVal > rightVal)
		case "<":
			return nativeBoolToBooleanObject(leftVal < rightVal)
		case "<>":
			return nativeBoolToBooleanObject(leftVal != rightVal)
		case "=":
			return nativeBoolToBooleanObject(leftVal == rightVal)
		default:
			return newError("Unknown Operator: %s %s %s", left.Type(), operator, right.Type())
	}

}

func nativeBoolToBooleanObject(input bool) *Boolean{
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("Unknown Operator: %s%s", operator, right.Type())
	

	}
}

func evalMinusPrefixOperatorExpression(right Object) Object{
	if(right.Type()!=INTEGER_OBJ){
		return nil
	}
	value := right.(*Integer).Value
	return &Integer{Value:-value}
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message:fmt.Sprintf(format, a...)}
}

// func evalStatements(stmts []parser.Statement) Object {
// 	var result Object
// 	for _, statement := range stmts {
// 		result = Eval(statement)
// 	}

// 	return result
// }
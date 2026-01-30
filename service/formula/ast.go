package interpreter

import "bytes"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type ExpressionStatement struct {
	Token Token
	Expression Expression
}

func (es *ExpressionStatement) String() string { return "" }

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.TokenValue
}

type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string { return "" }

func (i *Identifier) TokenLiteral() string {
	return i.Token.TokenValue
}

type IntegerLiteral struct {
	Token Token
	Value int
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) String() string { return "" }

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.TokenValue
}

type PrefixExpression struct {
	Token Token
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) String() string { 
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String() 
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.TokenValue
}




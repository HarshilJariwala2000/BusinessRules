package interpreter

import (
	"fmt"
	"strconv"
)

// import (
// )

const (
	_ int = iota
	LOWEST
	EQUAL
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
)

type Parser struct {
	l *Lexer
	errors []string

	currentToken Token
	peekToken Token

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns map[TokenType]infixParseFn

}

func (p *Parser) registerPrefixFunction(tokenType TokenType, fn prefixParseFn){
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfixFunction(tokenType TokenType, fn infixParseFn){
	p.infixParseFns[tokenType] = fn
}

func NewParser(l *Lexer) *Parser{
	p := &Parser{l:l, errors: []string{}}
	p.prefixParseFns = make(map[TokenType]prefixParseFn)
	p.registerPrefixFunction(IDENT, p.parseIdentifier)
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.currentToken, Value: p.currentToken.TokenValue}
}

func (p *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: p.currentToken}
	value, err := strconv.Atoi(p.currentToken.TokenValue)
	if err!=nil {
		msg := fmt.Sprintf("Cannot parse %q as Integer", p.currentToken.TokenValue)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) nextToken(){
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() Program {
	return nil
}

func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.currentToken.TokenType == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.TokenType == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", t, p.peekToken.TokenType)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseStatement() Statement{
	switch p.currentToken.TokenType {
		default:
			return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{Token: p.currentToken}
	stmt.Expression = p.parseExpression(LOWEST)

	if(p.peekTokenIs(EOF)){
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.currentToken.TokenType]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) expectPeek(t TokenType) bool {
	if(p.peekTokenIs(t)){
		p.nextToken()
		return true
	}else {
		p.peekError(t)
		return false
	}
}

type (
	prefixParseFn func() Expression
	infixParseFn func(Expression) Expression 
)
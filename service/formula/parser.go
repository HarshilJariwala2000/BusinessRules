package interpreter

import "fmt"

// import (
// )

type Parser struct {
	l *Lexer

	currentToken Token
	peekToken Token

	errors []string
}

func NewParser(l *Lexer) *Parser{
	p := &Parser{l:l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return p
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

func (p *Parser) expectPeek(t TokenType) bool {
	if(p.peekTokenIs(t)){
		p.nextToken()
		return true
	}else {
		p.peekError(t)
		return false
	}
}
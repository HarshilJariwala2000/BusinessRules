package interpreter

import (
	"text/scanner"
	"strings"
)

type Lexer struct{
	input string
	position int
	readPosition int
	ch string
	scannerTokenType string
	s scanner.Scanner
}

func New(input string) *Lexer{
	l := &Lexer{input: input}
	var s scanner.Scanner
	s.Init(strings.NewReader(l.input))
	l.s = s
	l.readChar()
	return l
}

func (l *Lexer) readChar(){
	tok := l.s.Scan()
	l.ch = l.s.TokenText()
	l.scannerTokenType = scanner.TokenString(tok)
	l.position = l.s.Position.Offset
	l.readPosition = l.position + 1
}

func newToken(tokenType TokenType, ch string) Token {
	return Token{ TokenType: tokenType, TokenValue: ch }
}

func (l * Lexer) NextToken() Token {
	var tok Token

	switch l.scannerTokenType {
		case "(":
			tok = newToken(LPAREN, l.ch)
		case ")":
			tok = newToken(RPAREN, l.ch)
		case ",":
			tok = newToken(COMMA, l.ch)
		case "Ident":
			switch l.ch {
				case "IF":
					tok = newToken(IF, l.ch)
				default:
					tok = newToken(IDENT, l.ch)
			}
		default:
			tok = newToken(TokenType(l.scannerTokenType), l.ch)
	}
	l.readChar()
	return tok
}

package parser

import (
	"github.com/philmish/s-tree/kvdb/qla"
	"github.com/philmish/s-tree/kvdb/qla/ast"
)

type Parser struct {
	l *qla.Lexer

	currToken qla.Token
	peekToken qla.Token
}

func New(l *qla.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two times to set peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgramm() *ast.Program {
	//TODO Implement recursive parsing
	return nil
}

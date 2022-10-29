package parser

import (
	"fmt"

	"github.com/philmish/s-tree/kvdb/qla"
	"github.com/philmish/s-tree/kvdb/qla/ast"
)

type Parser struct {
	l *qla.Lexer

	currToken qla.Token
	peekToken qla.Token

	errors []string
}

func New(l *qla.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two times to set peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t qla.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t qla.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t qla.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t qla.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s got %s.", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgramm() *ast.Program {
	programm := &ast.Program{}
	programm.Statements = []ast.Statement{}

	for p.currToken.Type != qla.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			programm.Statements = append(programm.Statements, stmt)
		}

		p.nextToken()
	}
	return programm
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case qla.ADD:
		return p.parseAddStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAddStatement() *ast.AddStatement {
	stmt := &ast.AddStatement{Token: p.currToken}

	if !qla.IsTypeToken(p.peekToken.Type) {
		return nil
	}
	p.nextToken()
	stmt.KeyType = p.currToken.Type

	if !p.expectPeek(qla.IDENT) {
		return nil
	}
	stmt.Key = p.currToken.Literal

	if !qla.IsTypeToken(p.peekToken.Type) {
		return nil
	}
	p.nextToken()
	stmt.ValueType = p.currToken.Type

	if !p.expectPeek(qla.IDENT) {
		return nil
	}
	stmt.Value = p.currToken.Literal

	for !p.curTokenIs(qla.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

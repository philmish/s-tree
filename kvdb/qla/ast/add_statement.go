package ast

import (
	"github.com/philmish/s-tree/kvdb/qla"
)

type AddStatement struct {
	Token     qla.Token
	KeyType   *qla.TokenType
	Key       string
	ValueType qla.TokenType
	Value     string
}

// Satisfy StatementNode interface
func (as *AddStatement) statementNode()       {}
func (as *AddStatement) TokenLiteral() string { return as.Token.Literal }

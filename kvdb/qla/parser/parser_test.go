package parser

import (
	"testing"

	"github.com/philmish/s-tree/kvdb/qla"
	"github.com/philmish/s-tree/kvdb/qla/ast"
)

func TestAddStatement(t *testing.T) {
	input := `
    ADD STR Hello STR World;
    ADD STR Foo STR Bar;
    `
	l := qla.NewLexer(input)
	p := New(l)

	programm := p.ParseProgramm()
	if programm == nil {
		t.Fatalf("Parsing Programm returned nil")
	}

	if len(programm.Statements) != 2 {
		t.Fatalf("Expected 2 Statements, got %d", len(programm.Statements))
	}

	tests := []struct {
		key string
		val string
	}{
		{key: "Hello", val: "World"},
		{key: "Foo", val: "Bar"},
	}
	for i, tc := range tests {
		stmt := programm.Statements[i]
		if !testAddStatement(t, stmt, tc.key, tc.val) {
			return
		}
	}
}

func testAddStatement(t *testing.T, s ast.Statement, key, val string) bool {
	if s.TokenLiteral() != "ADD" {
		t.Errorf("Expected ADD as literal, found %q", s.TokenLiteral())
		return false
	}

	addStmt, ok := s.(*ast.AddStatement)
	if !ok {
		t.Errorf("statement is not a AddStatement, found %T", s)
		return false
	}

	if addStmt.Key != key || addStmt.Value != val {
		t.Errorf("Expected key value pair %s:%s, found %s:%s", key, val, addStmt.Key, addStmt.Value)
		return false
	}

	return true
}

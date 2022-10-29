package parser

import (
	"testing"

	"github.com/philmish/s-tree/kvdb/qla"
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
}

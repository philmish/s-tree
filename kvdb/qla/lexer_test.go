package qla_test

import (
	"github.com/philmish/s-tree/kvdb/qla"
	"testing"
)

type tCase struct {
	expectedT   qla.TokenType
	expectedLit string
}

func runLexerTest(l *qla.Lexer, cases []tCase, t *testing.T) {
	for i, tt := range cases {
		tok := l.NextToken()

		if tok.Type != tt.expectedT {
			t.Fatalf(
				"test case [%d] - wrong token type - expected %q go %q",
				i, tt.expectedT, tok.Type,
			)
		}

		if tok.Literal != tt.expectedLit {
			t.Fatalf(
				"test case [%d] - wrong token literal - expected %q go %q",
				i, tt.expectedLit, tok.Literal,
			)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := "=()[:];"

	testCases := []tCase{
		{qla.EQUALS, "="},
		{qla.BRACEL, "("},
		{qla.BRACER, ")"},
		{qla.BRACKETL, "["},
		{qla.COLON, ":"},
		{qla.BRACKETR, "]"},
		{qla.SEMICOLON, ";"},
	}

	l := qla.NewLexer(input)
	runLexerTest(l, testCases, t)
}
func TestKeywords(t *testing.T) {
	input := "ADD(STR Hello = STR World);GET(STR Hello)"

	testCases := []tCase{
		{qla.ADD, "ADD"},
		{qla.BRACEL, "("},
		{qla.STR, "STR"},
		{qla.IDENT, "Hello"},
		{qla.EQUALS, "="},
		{qla.STR, "STR"},
		{qla.IDENT, "World"},
		{qla.BRACER, ")"},
		{qla.SEMICOLON, ";"},
		{qla.GET, "GET"},
		{qla.BRACEL, "("},
		{qla.STR, "STR"},
		{qla.IDENT, "Hello"},
		{qla.BRACER, ")"},
	}

	l := qla.NewLexer(input)
	runLexerTest(l, testCases, t)
}

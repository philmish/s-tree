package qla_test

import (
	"github.com/philmish/s-tree/kvdb/qla"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=()[:];"

	testCases := []struct {
		expectedT   qla.TokenType
		expectedLit string
	}{
		{qla.EQUALS, "="},
		{qla.BRACEL, "("},
		{qla.BRACER, ")"},
		{qla.BRACKETL, "["},
		{qla.COLON, ":"},
		{qla.BRACKETR, "]"},
		{qla.SEMICOLON, ";"},
	}

	l := qla.NewLexer(input)

	for i, tt := range testCases {
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

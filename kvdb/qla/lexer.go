package qla

type Lexer struct {
	input  string
	cursor int
	prev   int
	ch     byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.prev >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.prev]
	}
	l.cursor = l.prev
	l.prev += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '(':
		tok = newToken(BRACEL, l.ch)
	case ')':
		tok = newToken(BRACER, l.ch)
	case '{':
		tok = newToken(CURLYL, l.ch)
	case '}':
		tok = newToken(CURLYR, l.ch)
	case '[':
		tok = newToken(BRACKETL, l.ch)
	case ']':
		tok = newToken(BRACKETR, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case ':':
		tok = newToken(COLON, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '=':
		tok = newToken(EQUALS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

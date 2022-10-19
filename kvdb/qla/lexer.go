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

	l.skipWhitespace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdentifier(tok.Literal)
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
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

func (l *Lexer) readIdentifier() string {
	position := l.cursor
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.cursor]
}

func (l *Lexer) readNumber() string {
	position := l.cursor
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.cursor]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

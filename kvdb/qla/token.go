package qla

// Syntax example for command: ADD(STR Hello = INT 123)
// This command would add a key node with the "Hello" with a value node 123

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Delimiter
	COLON     = ":"
	SEMICOLON = ";"
	COMMA     = ","
	BRACKETL  = "["
	BRACKETR  = "]"
	CURLYL    = "{"
	CURLYR    = "}"
	BRACEL    = "("
	BRACER    = ")"

	// Operator
	EQUALS = "="

	// Types
	STR   = "STR"
	INT   = "INT"
	BOOL  = "BOOL"
	STRS  = "STRS"
	INTS  = "INTS"
	BOOLS = "BOOLS"

	// Commands
	ADD    = "ADD"
	GET    = "GET"
	UPDATE = "UPDATE"
	UPSERT = "UPSERT"
	DELETE = "DELETE"
)

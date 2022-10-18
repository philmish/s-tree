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
	IDENT   = "IDENT"

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

var keywords = map[string]TokenType{
	"ADD":    ADD,
	"GET":    GET,
	"UPDATE": UPDATE,
	"UPSERT": UPSERT,
	"DELETE": DELETE,
	"STR":    STR,
	"INT":    INT,
	"BOOL":   BOOL,
	"STRS":   STRS,
	"INTS":   INTS,
	"BOOLS":  BOOLS,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

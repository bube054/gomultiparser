package json

type TokenType int

const (
	// JSON Value
	String TokenType = iota // ""
	Number                  // 123, 45.6
	Object                  // { ... }
	Array                   // [ ... ]
	True                    // true
	False                   // false
	Null                    // null

	// JSON Delimiter
	LeftCurlyBrace   // {
	RightCurlyBrace  // }
	LeftSquareBrace  // [
	RightSquareBrace // ]
	Comma            // ,
	SemiColon        // :

	// Misc
	EOF
	ILLEGAL
)

func (tok TokenType) String() string {
	tokenNames := map[TokenType]string{
		0:  "string",
		1:  "number",
		2:  "object",
		3:  "array",
		4:  "true",
		5:  "false",
		6:  "null",
		7:  "leftCurlyBrace",
		8:  "rightCurlyBrace",
		9:  "leftSquareBrace",
		10: "rightSquareBrace",
		11: "comma",
		12: "semiColon",
		13: "EOF",
		14: "Illegal",
	}

	name, ok := tokenNames[tok]

	if !ok {
		return "Unknown"
	}

	return name
}

func (tok TokenType) Index() int {
	return int(tok)
}

type Token struct {
	Type TokenType
	Char string
}

func NewToken(typ TokenType, char string) Token {
	return Token{
		Type: typ,
		Char: char,
	}
}

package json

import (
	"reflect"
	"testing"
)

func TestLexingNull(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tokens []Token
	}{
		{name: "Lex nothing", input: "", tokens: []Token{}},
		{
			name:  "Lex null",
			input: "null",
			tokens: []Token{
				NewToken(Null, Null.String()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			ans := l.generateTokens()

			if !reflect.DeepEqual(ans, tt.tokens) {
				t.Errorf("got %v, wanted %v", ans, tt.tokens)
			}
		})
	}
}

func TestLexingBool(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tokens []Token
	}{
		{
			name:  "Lex true",
			input: "true",
			tokens: []Token{
				NewToken(True, True.String()),
			},
		},
		{
			name:  "Lex false",
			input: "false",
			tokens: []Token{
				NewToken(False, False.String()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			ans := l.generateTokens()

			if !reflect.DeepEqual(ans, tt.tokens) {
				t.Errorf("got %v, wanted %v", ans, tt.tokens)
			}
		})
	}
}

func TestLexingString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tokens []Token
	}{
		{
			name:  "Lex simple word",
			input: `"json"`,
			tokens: []Token{
				NewToken(String, `"json"`),
			},
		},
		{
			name:  "Lex simple sentence",
			input: `"json is a data format"`,
			tokens: []Token{
				NewToken(String, `"json is a data format"`),
			},
		},
		{
			name:  "Lex simple sentence with spaces",
			input: `" This string contains spaces "`,
			tokens: []Token{
				NewToken(String, `" This string contains spaces "`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			ans := l.generateTokens()

			if !reflect.DeepEqual(ans, tt.tokens) {
				t.Errorf("got %v, wanted %v", ans, tt.tokens)
			}
		})
	}
}

func TestLexingNumber(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tokens []Token
	}{
		{
			name:  "Lex single digit 0",
			input: `0`,
			tokens: []Token{
				NewToken(Number, `0`),
			},
		},
		{
			name:  "Lex single digit 9",
			input: `9`,
			tokens: []Token{
				NewToken(Number, `9`),
			},
		},
		{
			name:  "Lex multiple digits 989",
			input: `989`,
			tokens: []Token{
				NewToken(Number, `989`),
			},
		},
		{
			name:  "Lex negative digits -54578",
			input: `-54578`,
			tokens: []Token{
				NewToken(Number, `-54578`),
			},
		},
		{
			name:  "Lex float 3.14",
			input: `3.14`,
			tokens: []Token{
				NewToken(Number, `3.14`),
			},
		},
		{
			name:  "Lex negative float -3.14",
			input: `-3.14`,
			tokens: []Token{
				NewToken(Number, `-3.14`),
			},
		},
		{
			name:  "Lex negative float -3.14",
			input: `-3.14`,
			tokens: []Token{
				NewToken(Number, `-3.14`),
			},
		},
		{
			name:  "Lex exp 1e10",
			input: `1e10`,
			tokens: []Token{
				NewToken(Number, `1e10`),
			},
		},
		{
			name:  "Lex exp -2.5e-3",
			input: `-2.5e-3`,
			tokens: []Token{
				NewToken(Number, `-2.5e-3`),
			},
		},
		{
			name:  "Lex exp 6.022e23",
			input: `6.022e23`,
			tokens: []Token{
				NewToken(Number, `6.022e23`),
			},
		},
		{
			name:  "Lex exp 1E+3",
			input: `1E+3`,
			tokens: []Token{
				NewToken(Number, `1E+3`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			ans := l.generateTokens()

			if !reflect.DeepEqual(ans, tt.tokens) {
				t.Errorf("got %v, wanted %v", ans, tt.tokens)
			}
		})
	}
}

func TestLexingArray(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tokens []Token
	}{
		// {
		// 	name:  "Lex simple word",
		// 	input: `"json"`,
		// 	tokens: []Token{
		// 		NewToken(String, `"json"`),
		// 	},
		// },
		// {
		// 	name:  "Lex simple sentence",
		// 	input: `"json is a data format"`,
		// 	tokens: []Token{
		// 		NewToken(String, `"json is a data format"`),
		// 	},
		// },
		// {
		// 	name:  "Lex simple sentence with spaces",
		// 	input: `" This string contains spaces "`,
		// 	tokens: []Token{
		// 		NewToken(String, `" This string contains spaces "`),
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			ans := l.generateTokens()

			if !reflect.DeepEqual(ans, tt.tokens) {
				t.Errorf("got %v, wanted %v", ans, tt.tokens)
			}
		})
	}
}

func TestIsString(t *testing.T) {
	tests := []struct {
		name            string
		cur, next, prev rune
		exp             bool
	}{
		// No " or \ or control characters
		{name: "Valid string (abc)", cur: 'a', next: 'b', prev: 'c', exp: true},
		{name: "Valid string (pqr)", cur: 'p', next: 'q', prev: 'r', exp: true},
		{name: "Valid string (xy)", cur: 'x', next: 0, prev: 'y', exp: true},

		// Starts with \ but ends with escapable char
		{name: "Valid string (\\\")", cur: '\\', next: '"', prev: 0, exp: true},
		{name: "Valid string (\\\\)", cur: '\\', next: '\\', prev: 0, exp: true},
		{name: "Valid string (\\/)", cur: '\\', next: '/', prev: 0, exp: true},
		{name: "Valid string (\\b)", cur: '\\', next: 'b', prev: 0, exp: true},
		{name: "Valid string (\\f)", cur: '\\', next: 'f', prev: 0, exp: true},
		{name: "Valid string (\\n)", cur: '\\', next: 'n', prev: 0, exp: true},
		{name: "Valid string (\\r)", cur: '\\', next: 'r', prev: 0, exp: true},
		{name: "Valid string (\\t)", cur: '\\', next: 't', prev: 0, exp: true},
		{name: "Valid string (\\u)", cur: '\\', next: 'u', prev: 0, exp: true},

		// Starts with \ but does not ends with escapable char
		{name: "Valid string (\\a)", cur: '\\', next: 'a', prev: 0, exp: false},
		{name: "Valid string (\\x)", cur: '\\', next: 'x', prev: 0, exp: false},
		{name: "Valid string (\\y)", cur: '\\', next: 'y', prev: 0, exp: false},

		// Is " or \ or control characters that has next character is not escapable
		{name: "Invalid string (\"yz), is \"", cur: '"', next: 'y', prev: 'z', exp: false},
		{name: "Invalid string (\\yz), is \\", cur: '\\', next: 'y', prev: 'z', exp: false},
		{name: "Invalid string (	yz), is 	", cur: '	', next: 'y', prev: 'z', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isString(tt.cur, tt.next, tt.prev)

			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsBackslash(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		{name: "Valid backslash", input: '\\', exp: true},
		{name: "Invalid backslash (/)", input: '/', exp: false},
		{name: "Invalid backslash (t)", input: 't', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isBackslash(tt.input)

			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsDoubleQuote(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		{name: "Valid double quote", input: '"', exp: true},
		{name: "Invalid double quote (')", input: '\'', exp: false},
		{name: "Invalid double quote (a)", input: 'a', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isDoubleQuote(tt.input)

			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsControlChar(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		// C0 control characters (U+0000–U+001F)
		{name: "NUL", input: '\u0000', exp: true},
		{name: "BEL", input: '\u0007', exp: true},
		{name: "BS", input: '\u0008', exp: true},
		{name: "TAB", input: '\u0009', exp: true},
		{name: "LF", input: '\u000A', exp: true},
		{name: "CR", input: '\u000D', exp: true},
		{name: "US", input: '\u001F', exp: true},

		// DEL (U+007F)
		{name: "DEL", input: '\u007F', exp: true},

		// C1 control characters (U+0080–U+009F)
		{name: "PAD", input: '\u0080', exp: true},
		{name: "NEL", input: '\u0085', exp: true},
		{name: "APC", input: '\u009F', exp: true},

		// Non-control characters (expected false)
		{name: "Space", input: ' ', exp: false},
		{name: "Digit", input: '5', exp: false},
		{name: "Letter", input: 'A', exp: false},
		{name: "Emoji", input: '😊', exp: false},
		{name: "Greek letter", input: 'Ω', exp: false},
		{name: "Hebrew letter", input: 'ש', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isControlChar(tt.input)
			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsWhiteSpace(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		{name: "Valid whitespace (' ')", input: ' ', exp: true},
		{name: "Valid whitespace ('\t')", input: '\t', exp: true},
		{name: "Valid whitespace ('\n')", input: '\n', exp: true},
		{name: "Valid whitespace ('\r')", input: '\r', exp: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isWhiteSpace(tt.input)
			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		{name: "Valid digit (0)", input: '0', exp: true},
		{name: "Valid digit (1)", input: '1', exp: true},
		{name: "Valid digit (2)", input: '2', exp: true},
		{name: "Valid digit (3)", input: '3', exp: true},
		{name: "Valid digit (4)", input: '4', exp: true},
		{name: "Valid digit (5)", input: '5', exp: true},
		{name: "Valid digit (6)", input: '6', exp: true},
		{name: "Valid digit (7)", input: '7', exp: true},
		{name: "Valid digit (8)", input: '8', exp: true},
		{name: "Valid digit (9)", input: '9', exp: true},

		{name: "Invalid digit (a)", input: 'a', exp: false},
		{name: "Invalid digit (a)", input: '@', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isDigit(tt.input)
			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsMinus(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		{name: "Valid minus (-)", input: '-', exp: true},
		{name: "Valid plus (+)", input: '+', exp: true},

		{name: "Invalid times (*)", input: '*', exp: false},
		{name: "Invalid divide (/)", input: '/', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isPlusOrMinus(tt.input)
			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

func TestIsE(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		exp   bool
	}{
		{name: "Valid e", input: 'e', exp: true},
		{name: "Valid E", input: 'E', exp: true},

		{name: "Invalid é", input: 'é', exp: false},
		{name: "Invalid È", input: 'È', exp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isE(tt.input)
			if ans != tt.exp {
				t.Errorf("got %v, wanted %v", ans, tt.exp)
			}
		})
	}
}

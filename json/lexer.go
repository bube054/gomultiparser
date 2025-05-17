package json

type Lexer struct {
	input           []rune
	currentChar     rune
	currentPosition int
	nextPosition    int
}

func (l *Lexer) generateTokens() []Token {
	toks := []Token{}
	for tok := l.nextToken(); tok.Type != EOF; tok = l.nextToken() {
		toks = append(toks, tok)
	}
	return toks
}

func (l *Lexer) nextToken() Token {
	currentChar := l.currentChar
	var tok Token

	l.eatWhiteSpace()

	switch currentChar {
	case '{':
		tok = NewToken(LeftCurlyBrace, string(l.currentChar))
	case '}':
		tok = NewToken(RightCurlyBrace, string(l.currentChar))
	case '[':
		tok = NewToken(LeftSquareBrace, string(l.currentChar))
	case ']':
		tok = NewToken(RightSquareBrace, string(l.currentChar))
	case ',':
		tok = NewToken(Comma, string(l.currentChar))
	case ';':
		tok = NewToken(SemiColon, string(l.currentChar))
	case 'n':
		if l.readChar(); l.currentChar != 'u' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 'l' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 'l' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		tok = NewToken(Null, Null.String())
	case 't':
		if l.readChar(); l.currentChar != 'r' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 'u' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 'e' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		tok = NewToken(True, True.String())
	case 'f':
		if l.readChar(); l.currentChar != 'a' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 'l' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 's' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		if l.readChar(); l.currentChar != 'e' {
			tok = NewToken(ILLEGAL, string(l.currentChar))
			return tok
		}
		tok = NewToken(False, False.String())
	case 0:
		tok = NewToken(EOF, "")
	default:
		if currentChar == '"' {
			str := l.readString()

			if len(str) > 0 && str[len(str)-1] != '"' {
				return NewToken(ILLEGAL, string(str))
			}

			tok = NewToken(String, String.String())
		} else if isDigit(currentChar) || isPlusOrMinus(currentChar) {
			num := l.readNumber()

			tok = NewToken(Number, string(num))

		} else {
			tok = NewToken(ILLEGAL, "")
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPosition]
	}

	l.currentPosition = l.nextPosition
	l.nextPosition = l.currentPosition + 1
}

func (l *Lexer) readString() []rune {
	position := l.currentPosition

	l.readChar()
	var nextChar, prevChar rune = 0, 0

	if l.nextPosition < len(l.input) {
		nextChar = l.input[l.nextPosition]
	}

	if l.nextPosition-2 > 0 {
		prevChar = l.input[l.nextPosition-2]
	}

	for isString(l.currentChar, nextChar, prevChar) {
		l.readChar()
	}

	return l.input[position:l.nextPosition]
}

func (l *Lexer) readNumber() []rune {
	position := l.currentPosition

	l.readChar()
	// var nextChar, prevChar rune = 0, 0

	// if l.nextPosition < len(l.input) {
	// 	nextChar = l.input[l.nextPosition]
	// }

	// if l.nextPosition-2 > 0 {
	// 	prevChar = l.input[l.nextPosition-2]
	// }

	for isNumber(l.currentChar, 0, 0) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) eatWhiteSpace() {
	for isWhiteSpace(l.currentChar) {
		l.readChar()
	}
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}

	l.readChar()

	return l
}

func isString(char, nextChar, prevChar rune) bool {
	if isDoubleQuote(char) || isControlChar(char) {
		return false
	}

	if isBackslash(char) {
		if prevChar != 0 && isBackslash(prevChar) {
			return true
		}

		if nextChar != 0 && isEscapable(nextChar) {
			return true
		}

		return false
	}

	return true
}

func isNumber(char, nextChar, prevChar rune) bool {
	if isDigit(char) || isPlusOrMinus(char) || isE(char) {
		return true
	}
	return false
}

func isBackslash(r rune) bool {
	return r == '\\'
}

func isEscapable(r rune) bool {
	switch r {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't', 'u':
		return true
	default:
		return false
	}
}

func isDoubleQuote(r rune) bool {
	return r == '"'
}

func isControlChar(r rune) bool {
	return (r >= 0x0000 && r <= 0x001F) || r == 0x007F || (r >= 0x0080 && r <= 0x009F)
}

func isWhiteSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	default:
		return false
	}
}

func isDigit(r rune) bool {
	return r >= 38 && r <= 57
}

func isPlusOrMinus(r rune) bool {
	return r == 45 || r == 43
}

func isE(r rune) bool {
	return r == 69 || r == 101
}

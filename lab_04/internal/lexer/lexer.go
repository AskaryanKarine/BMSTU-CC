package lexer

import "unicode"

type Lexer struct {
	input   string
	pos     int
	line    int
	column  int
	start   int
	startLn int
	startCl int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		line:   1,
		column: 1,
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	l.start = l.pos
	l.startLn = l.line
	l.startCl = l.column

	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF}
	}

	if typ, ok := l.tryOperator(); ok {
		return typ
	}

	ch := l.input[l.pos]

	if unicode.IsLetter(rune(ch)) {
		return l.readIdentifier()
	}

	if unicode.IsDigit(rune(ch)) {
		return l.readNumber()
	}

	l.pos++
	l.column++
	return Token{Type: TokenERROR, Literal: string(ch)}
}

func (l *Lexer) tryOperator() (Token, bool) {
	if l.pos+1 < len(l.input) {
		sub := l.input[l.pos : l.pos+2]
		if typ, ok := operators[sub]; ok {
			tok := Token{
				Type:    typ,
				Literal: sub,
			}
			l.pos += 2
			l.column += 2
			return tok, true
		}
	}

	sub := l.input[l.pos : l.pos+1]
	if typ, ok := operators[sub]; ok {
		tok := Token{
			Type:    typ,
			Literal: sub,
		}
		l.pos++
		l.column++
		return tok, true
	}

	return Token{}, false
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		if ch == '\n' {
			l.line++
			l.column = 1
			l.pos++
			continue
		}
		if unicode.IsSpace(rune(ch)) {
			l.pos++
			l.column++
			continue
		}
		break
	}
}

func (l *Lexer) readIdentifier() Token {
	start := l.pos
	for l.pos < len(l.input) {
		ch := rune(l.input[l.pos])
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
			break
		}
		l.pos++
		l.column++
	}

	literal := l.input[start:l.pos]

	if typ, ok := keywords[literal]; ok {
		return Token{Type: typ, Literal: literal}
	}
	return Token{Type: TokenIDENT, Literal: literal}
}

func (l *Lexer) readNumber() Token {
	start := l.pos
	for l.pos < len(l.input) {
		ch := rune(l.input[l.pos])
		if !unicode.IsDigit(ch) {
			break
		}
		l.pos++
		l.column++
	}

	literal := l.input[start:l.pos]
	return Token{Type: TokenNUMBER, Literal: literal}
}

package lexer

const (
	TokenEOF = iota
	TokenERROR
	TokenIDENT
	TokenNUMBER
	TokenBEGIN
	TokenEND
	TokenSEMICOLON
	TokenASSIGN
	TokenEQ
	TokenPLUS
	TokenMINUS
	TokenMULT
	TokenDIV
	TokenLT
	TokenLE
	TokenGT
	TokenGE
	TokenNE
	TokenLPAREN
	TokenRPAREN
)

var (
	keywords = map[string]int{
		"begin": TokenBEGIN,
		"end":   TokenEND,
	}

	operators = map[string]int{
		";":  TokenSEMICOLON,
		"=":  TokenASSIGN,
		"==": TokenEQ,
		"+":  TokenPLUS,
		"-":  TokenMINUS,
		"*":  TokenMULT,
		"/":  TokenDIV,
		"(":  TokenLPAREN,
		")":  TokenRPAREN,
		"<":  TokenLT,
		"<=": TokenLE,
		">":  TokenGT,
		">=": TokenGE,
		"<>": TokenNE,
	}
)

type Token struct {
	Type    int
	Literal string
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == TokenEOF || tok.Type == TokenERROR {
			break
		}
	}
	return tokens
}

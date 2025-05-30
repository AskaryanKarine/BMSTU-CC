package parser

import "github.com/AskaryanKarine/BMSTU-CC/lab_03/internal/lexer"

type ASTNode struct {
	Type     string
	Value    string
	Children []*ASTNode
	Token    lexer.Token
}

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) CurrentToken() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.TokenEOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) match(tokenType int) bool {
	if p.CurrentToken().Type == tokenType {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) expect(tokenType int) (lexer.Token, bool) {
	tok := p.CurrentToken()
	if tok.Type == tokenType {
		p.advance()
		return tok, true
	}
	return tok, false
}

// Правила грамматики
func (p *Parser) parseFactor() (*ASTNode, bool) {
	tok := p.CurrentToken()
	switch tok.Type {
	case lexer.TokenIDENT:
		p.advance()
		return &ASTNode{
			Type:  "Factor",
			Value: tok.Literal,
			Token: tok,
		}, true
	case lexer.TokenNUMBER:
		p.advance()
		return &ASTNode{
			Type:  "Factor",
			Value: tok.Literal,
			Token: tok,
		}, true
	case lexer.TokenLPAREN:
		p.advance()
		expr, ok := p.parseArithExpr()
		if !ok {
			return nil, false
		}
		if _, ok := p.expect(lexer.TokenRPAREN); !ok {
			return nil, false
		}
		return &ASTNode{
			Type:     "Factor",
			Children: []*ASTNode{expr},
			Token:    tok,
		}, true
	default:
		return nil, false
	}
}

func (p *Parser) parseTerm() (*ASTNode, bool) {
	left, ok := p.parseFactor()
	if !ok {
		return nil, false
	}

	for {
		tok := p.CurrentToken()
		if tok.Type == lexer.TokenMULT || tok.Type == lexer.TokenDIV {
			p.advance()
			right, ok := p.parseFactor()
			if !ok {
				return nil, false
			}
			left = &ASTNode{
				Type: "Term",
				Children: []*ASTNode{
					left,
					{Type: "MulOp", Value: tok.Literal, Token: tok},
					right,
				},
				Token: tok,
			}
		} else {
			break
		}
	}
	return left, true
}

func (p *Parser) parseArithExpr() (*ASTNode, bool) {
	left, ok := p.parseTerm()
	if !ok {
		return nil, false
	}

	for {
		tok := p.CurrentToken()
		if tok.Type == lexer.TokenPLUS || tok.Type == lexer.TokenMINUS {
			p.advance()
			right, ok := p.parseTerm()
			if !ok {
				return nil, false
			}
			left = &ASTNode{
				Type: "ArithExpr",
				Children: []*ASTNode{
					left,
					{Type: "AddOp", Value: tok.Literal, Token: tok},
					right,
				},
				Token: tok,
			}
		} else {
			break
		}
	}
	return left, true
}

func (p *Parser) parseExpression() (*ASTNode, bool) {
	left, ok := p.parseArithExpr()
	if !ok {
		return nil, false
	}

	tok := p.CurrentToken()
	if tok.Type == lexer.TokenLT || tok.Type == lexer.TokenLE ||
		tok.Type == lexer.TokenGT || tok.Type == lexer.TokenGE ||
		tok.Type == lexer.TokenEQ || tok.Type == lexer.TokenNE {
		p.advance()
		right, ok := p.parseArithExpr()
		if !ok {
			return nil, false
		}
		return &ASTNode{
			Type: "Expression",
			Children: []*ASTNode{
				left,
				{Type: "RelOp", Value: tok.Literal, Token: tok},
				right,
			},
			Token: tok,
		}, true
	}
	return &ASTNode{
		Type:     "Expression",
		Children: []*ASTNode{left},
		Token:    tok,
	}, true
}

func (p *Parser) parseOperator() (*ASTNode, bool) {
	idTok, ok := p.expect(lexer.TokenIDENT)
	if !ok {
		return nil, false
	}

	if _, ok := p.expect(lexer.TokenASSIGN); !ok {
		return nil, false
	}

	expr, ok := p.parseExpression()
	if !ok {
		return nil, false
	}

	return &ASTNode{
		Type: "Operator",
		Children: []*ASTNode{
			{Type: "Identifier", Value: idTok.Literal, Token: idTok},
			expr,
		},
		Token: idTok,
	}, true
}

func (p *Parser) parseOperatorTail() (*ASTNode, bool) {
	if !p.match(lexer.TokenSEMICOLON) {
		return &ASTNode{
			Type: "OperatorTail",
		}, true
	}

	op, ok := p.parseOperator()
	if !ok {
		return nil, false
	}

	tail, ok := p.parseOperatorTail()
	if !ok {
		return nil, false
	}

	return &ASTNode{
		Type:     "OperatorTail",
		Children: []*ASTNode{op, tail},
	}, true
}

func (p *Parser) parseStmtList() (*ASTNode, bool) {
	first, ok := p.parseOperator()
	if !ok {
		return nil, false
	}

	tail, ok := p.parseOperatorTail()
	if !ok {
		return nil, false
	}

	return &ASTNode{
		Type:     "StmtList",
		Children: []*ASTNode{first, tail},
	}, true
}

func (p *Parser) parseBlock() (*ASTNode, bool) {
	if !p.match(lexer.TokenBEGIN) {
		return nil, false
	}

	stmts, ok := p.parseStmtList()
	if !ok {
		return nil, false
	}

	if !p.match(lexer.TokenEND) {
		return nil, false
	}

	return &ASTNode{
		Type:     "Block",
		Children: []*ASTNode{stmts},
	}, true
}

func (p *Parser) parseProgram() (*ASTNode, bool) {
	block, ok := p.parseBlock()
	if !ok {
		return nil, false
	}

	return &ASTNode{
		Type:     "Program",
		Children: []*ASTNode{block},
	}, true
}

func (p *Parser) Parse() (*ASTNode, bool) {
	return p.parseProgram()
}

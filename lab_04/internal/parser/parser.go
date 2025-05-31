package parser

import "github.com/AskaryanKarine/BMSTU-CC/lab_04/internal/lexer"

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

func (p *Parser) parseProgram() (*ASTNode, bool) {
	expr, ok := p.parseExpression()
	if !ok {
		return nil, false
	}

	if p.CurrentToken().Type != lexer.TokenEOF {
		return nil, false
	}
	return expr, true
}

func (p *Parser) Parse() (*ASTNode, bool) {
	return p.parseProgram()
}

func ASTToRPN(node *ASTNode) []string {
	if node == nil {
		return nil
	}

	var rpn []string

	if len(node.Children) == 3 &&
		(node.Type == "Expression" || node.Type == "ArithExpr" || node.Type == "Term") {
		left := ASTToRPN(node.Children[0])
		right := ASTToRPN(node.Children[2])
		op := ASTToRPN(node.Children[1])

		rpn = append(rpn, left...)
		rpn = append(rpn, right...)
		rpn = append(rpn, op...)
		return rpn
	}

	if node.Type == "Factor" {
		if len(node.Children) == 3 {
			return ASTToRPN(node.Children[1])
		} else if len(node.Children) == 1 {
			return ASTToRPN(node.Children[0])
		}
	}

	if node.Value != "" {
		return []string{node.Value}
	}

	for _, child := range node.Children {
		rpn = append(rpn, ASTToRPN(child)...)
	}

	return rpn
}

func ASTToRPN1(node *ASTNode) ([]string, *ASTNode) {
	if node == nil {
		return nil, nil
	}

	switch node.Type {
	case "Expression", "ArithExpr", "Term":
		var rpn []string
		var rpnNode *ASTNode
		var childrenNodes []*ASTNode

		for _, child := range node.Children {
			childRPN, childNode := ASTToRPN1(child)
			rpn = append(rpn, childRPN...)
			if childNode != nil {
				childrenNodes = append(childrenNodes, childNode)
			}
		}

		// Для RPN дерева создаем узел-оператор с операндами в качестве детей
		if len(childrenNodes) > 0 && node.Value != "" {
			rpnNode = &ASTNode{
				Type:  node.Type + "_RPN",
				Value: node.Value,
			}
			for _, child := range childrenNodes {
				rpnNode.Children = append(rpnNode.Children, child)
			}
		} else if len(childrenNodes) > 0 {
			rpnNode = &ASTNode{
				Type:     node.Type + "_RPN",
				Children: childrenNodes,
			}
		}

		return rpn, rpnNode

	case "Factor", "Identifier", "Number":
		return []string{node.Value}, &ASTNode{
			Type:  node.Type + "_RPN",
			Value: node.Value,
			Token: node.Token,
		}

	case "AddOp", "MulOp", "RelOp":
		return []string{node.Value}, &ASTNode{
			Type:  node.Type + "_RPN",
			Value: node.Value,
			Token: node.Token,
		}

	default:
		return nil, nil
	}
}

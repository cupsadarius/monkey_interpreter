package parser

import (
	"fmt"
	"strconv"

	"github.com/cupsadarius/monkey_interpreter/ast"
	"github.com/cupsadarius/monkey_interpreter/token"
	"github.com/cupsadarius/monkey_interpreter/utils"
)

type (
	prefixParserFn func() ast.Expression
	infixParserFn  func(ast.Expression) ast.Expression
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // + -
	PRODUCT     // * /
	PREFIX      // -x or !x
	CALL        // myFunction(x)
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.LT_EQ:    LESSGREATER,
	token.GT_EQ:    LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) parseIdentifier() ast.Expression {
	defer utils.UnTrace(utils.Trace("parseIdentifier"))

	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer utils.UnTrace(utils.Trace("parseIntegerLiteral"))
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer at %d: %d", p.curToken.Literal, p.curToken.Line, p.curToken.Column)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	defer utils.UnTrace(utils.Trace("parseFloatLiteral"))

	lit := &ast.FloatLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float at %d: %d", p.curToken.Literal, p.curToken.Line, p.curToken.Column)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	defer utils.UnTrace(utils.Trace("parseBooleanLiteral"))
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer utils.UnTrace(utils.Trace("parseExpression"))

	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return nil
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	defer utils.UnTrace(utils.Trace("parsePrefixExpression"))

	expression := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Literal}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer utils.UnTrace(utils.Trace("parseInfixExpression"))

	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

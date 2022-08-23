package parser

import (
	"fmt"
	"strconv"

	"github.com/cupsadarius/monkey_interpreter/ast"
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

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
    p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
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
  lit := &ast.FloatLiteral{Token: p.curToken}

  value, err := strconv.ParseFloat(p.curToken.Literal,  64)

  if err != nil {
    msg := fmt.Sprintf("could not parse %q as float at %d: %d", p.curToken.Literal, p.curToken.Line, p.curToken.Column)
    p.errors = append(p.errors, msg)
    return nil
  }

  lit.Value = value 

  return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
  expression := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Literal}

  p.nextToken()

  expression.Right = p.parseExpression(PREFIX)

  return expression
}

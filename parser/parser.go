package parser

import (
	"fmt"

	"github.com/cupsadarius/monkey_interpreter/ast"
	"github.com/cupsadarius/monkey_interpreter/lexer"
	"github.com/cupsadarius/monkey_interpreter/token"
	"github.com/cupsadarius/monkey_interpreter/utils"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParserFn
	infixParseFns  map[token.TokenType]infixParserFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens so curToken and peekToken are populated
	p.nextToken()
	p.nextToken()

  p.prefixParseFns = make(map[token.TokenType]prefixParserFn)

  p.registerPrefix(token.IDENT, p.parseIdentifier)
  p.registerPrefix(token.INT, p.parseIntegerLiteral)
  p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
  p.registerPrefix(token.BANG, p.parsePrefixExpression)
  p.registerPrefix(token.MINUS, p.parsePrefixExpression)
  p.registerPrefix(token.TRUE, p.parseBooleanLiteral)
  p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
  p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
  p.registerPrefix(token.IF, p.parseIfExpression)
  p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
  
  p.infixParseFns = make(map[token.TokenType]infixParserFn)
  p.registerInfix(token.PLUS, p.parseInfixExpression)
  p.registerInfix(token.MINUS, p.parseInfixExpression)
  p.registerInfix(token.SLASH, p.parseInfixExpression)
  p.registerInfix(token.ASTERISK, p.parseInfixExpression)
  p.registerInfix(token.EQ, p.parseInfixExpression)
  p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
  p.registerInfix(token.LT, p.parseInfixExpression)
  p.registerInfix(token.LT_EQ, p.parseInfixExpression)
  p.registerInfix(token.GT, p.parseInfixExpression)
  p.registerInfix(token.GT_EQ, p.parseInfixExpression)
  p.registerInfix(token.LPAREN, p.parseCallExpression)

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead at line %d, column %d", t, p.peekToken.Type, p.peekToken.Line, p.peekToken.Column)

	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
  msg := fmt.Sprintf("no prefix parse function for %s found", t)

  p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
  defer utils.UnTrace(utils.Trace("parseLetStatement"))

	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

  p.nextToken()

  stmt.Value = p.parseExpression(LOWEST)


	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
  defer utils.UnTrace(utils.Trace("parseReturnStatement"))

	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

  stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParserFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParserFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

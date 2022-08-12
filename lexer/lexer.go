package lexer

import (
	"github.com/cupsadarius/monkey_interpreter/token"
)

type Lexer struct {
	input         string
	position      int  // current position in the input
	readPosition  int  // current reading position in input
	ch            byte // current char under examination
	currentLine   int  // current line in input
	currentColumn int  // current postion on the line
}

func newToken(tokenType token.TokenType, ch byte, line, col int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line, Column: col}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func New(input string) *Lexer {
	l := &Lexer{input: input, currentLine: 1, currentColumn: 0}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	switch l.ch {
	case '\n':
		l.currentLine += 1
		l.currentColumn = 0
	default:
		l.currentColumn += 1
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isPartOfNumber(ch byte) bool {
	ahead := l.peakAhead()
	behind := l.peakBack()

	return isDigit(ch) || ch == '.' && (ahead != 0 && isDigit(ahead) || behind != 0 && isDigit(behind))
}

func (l *Lexer) readNumber() string {
	position := l.position
	for l.isPartOfNumber(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peakAhead() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) peakBack() byte {
	if l.position-1 < 0 {
		return 0
	}
	return l.input[l.position-1]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.currentLine, l.currentColumn)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.currentLine, l.currentColumn)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.currentLine, l.currentColumn)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.currentLine, l.currentColumn)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.currentLine, l.currentColumn)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.currentLine, l.currentColumn)
	case '.':
		if l.isPartOfNumber(l.ch) {
			tok = token.Token{}
			tok.Literal = l.readNumber()
			tok.Type = token.LookupNumericIdentifier(tok.Literal)
			tok.Line = l.currentLine
			tok.Column = l.currentColumn - 1

			return tok
		} else {
			tok = newToken(token.DOT, l.ch, l.currentLine, l.currentColumn)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, l.currentLine, l.currentColumn)
	case '=':
		if l.peakAhead() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal, Line: l.currentLine, Column: l.currentColumn}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.currentLine, l.currentColumn)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch, l.currentLine, l.currentColumn)
	case '!':
		if l.peakAhead() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, Line: l.currentLine, Column: l.currentColumn}
		} else {
			tok = newToken(token.BANG, l.ch, l.currentLine, l.currentColumn)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.currentLine, l.currentColumn)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.currentLine, l.currentColumn)
	case '<':
		if l.peakAhead() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LT_EQ, Literal: literal, Line: l.currentLine, Column: l.currentColumn}
		} else {
			tok = newToken(token.LT, l.ch, l.currentLine, l.currentColumn)
		}
	case '>':
		if l.peakAhead() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GT_EQ, Literal: literal, Line: l.currentLine, Column: l.currentColumn}
		} else {
			tok = newToken(token.GT, l.ch, l.currentLine, l.currentColumn)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			tok.Line = l.currentLine
			tok.Column = l.currentColumn - 1

			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.LookupNumericIdentifier(tok.Literal)
			tok.Line = l.currentLine
			tok.Column = l.currentColumn - 1

			return tok
		} else {
			tok = newToken(token.ILEGAL, l.ch, l.currentLine, l.currentColumn)
		}
	}

	l.readChar()

	return tok
}

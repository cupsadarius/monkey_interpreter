package lexer

import (
	"testing"

	"github.com/cupsadarius/monkey_interpreter/token"
)

func TestSingleCharSymbols(t *testing.T) {
	input := `=+(){},;.!-/*<>&`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.ASSIGN, "=", 1, 1},
		{token.PLUS, "+", 1, 2},
		{token.LPAREN, "(", 1, 3},
		{token.RPAREN, ")", 1, 4},
		{token.LBRACE, "{", 1, 5},
		{token.RBRACE, "}", 1, 6},
		{token.COMMA, ",", 1, 7},
		{token.SEMICOLON, ";", 1, 8},
		{token.DOT, ".", 1, 9},
		{token.BANG, "!", 1, 10},
		{token.MINUS, "-", 1, 11},
		{token.SLASH, "/", 1, 12},
		{token.ASTERISK, "*", 1, 13},
		{token.LT, "<", 1, 14},
		{token.GT, ">", 1, 15},
		{token.ILEGAL, "&", 1, 16},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - tokenLine wrong. expected=%d, got=%d", i, tt.expectedLine, tok.Line)
		}
		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - tokenColumn wrong. expected=%d, got=%d", i, tt.expectedColumn, tok.Column)
		}
	}
}

func TestFloats(t *testing.T) {
	input := `.023; 1.23; 1.;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.FLOAT, ".023", 1, 4},
		{token.SEMICOLON, ";", 1, 5},
		{token.FLOAT, "1.23", 1, 10},
		{token.SEMICOLON, ";", 1, 11},
		{token.FLOAT, "1.", 1, 14},
		{token.SEMICOLON, ";", 1, 15},
		{token.EOF, "", 1, 16},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - tokenLine wrong. expected=%d, got=%d", i, tt.expectedLine, tok.Line)
		}
		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - tokenColumn wrong. expected=%d, got=%d", i, tt.expectedColumn, tok.Column)
		}
	}
}

func TestSimpleSyntax(t *testing.T) {
	input := `
  let five = 5;
  let ten = 10;

  let add = fn(x, y) {
    return x + y;
  };

  let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.LET, "let", 2, 5},
		{token.IDENT, "five", 2, 10},
		{token.ASSIGN, "=", 2, 12},
		{token.INT, "5", 2, 14},
		{token.SEMICOLON, ";", 2, 15},
		{token.LET, "let", 3, 5},
		{token.IDENT, "ten", 3, 9},
		{token.ASSIGN, "=", 3, 11},
		{token.INT, "10", 3, 14},
		{token.SEMICOLON, ";", 3, 15},
		{token.LET, "let", 5, 5},
		{token.IDENT, "add", 5, 9},
		{token.ASSIGN, "=", 5, 11},
		{token.FUNCTION, "fn", 5, 14},
		{token.LPAREN, "(", 5, 15},
		{token.IDENT, "x", 5, 16},
		{token.COMMA, ",", 5, 17},
		{token.IDENT, "y", 5, 19},
		{token.RPAREN, ")", 5, 20},
		{token.LBRACE, "{", 5, 22},
		{token.RETURN, "return", 6, 10},
		{token.IDENT, "x", 6, 12},
		{token.PLUS, "+", 6, 14},
		{token.IDENT, "y", 6, 16},
		{token.SEMICOLON, ";", 6, 17},
		{token.RBRACE, "}", 7, 3},
		{token.SEMICOLON, ";", 7, 4},
		{token.LET, "let", 9, 5},
		{token.IDENT, "result", 9, 12},
		{token.ASSIGN, "=", 9, 14},
		{token.IDENT, "add", 9, 18},
		{token.LPAREN, "(", 9, 19},
		{token.IDENT, "five", 9, 23},
		{token.COMMA, ",", 9, 24},
		{token.IDENT, "ten", 9, 28},
		{token.RPAREN, ")", 9, 29},
		{token.SEMICOLON, ";", 9, 30},
		{token.EOF, "", 9, 31},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - tokenLine wrong. expected=%d, got=%d", i, tt.expectedLine, tok.Line)
		}
		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - tokenColumn wrong. expected=%d, got=%d", i, tt.expectedColumn, tok.Column)
		}
	}
}

func TestComplexSyntax(t *testing.T) {
	input := `let five = 5;
  let ten = 10;

  let add = fn(x, y) {
    x + y;
  };

  let result = add(five, ten);
  !-/*5;
  5 < 10 > 5;

  if (5 < 10) {
    return true;
  } else {
    return false;
  }

  10 == 10;
  10 != 9;
  10 <= 11;
  11 >= 10;
	"foobar";
	"foo bar";
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.LET, "let", 1, 3},
		{token.IDENT, "five", 1, 8},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},
		{token.LET, "let", 2, 5},
		{token.IDENT, "ten", 2, 9},
		{token.ASSIGN, "=", 2, 11},
		{token.INT, "10", 2, 14},
		{token.SEMICOLON, ";", 2, 15},
		{token.LET, "let", 4, 5},
		{token.IDENT, "add", 4, 9},
		{token.ASSIGN, "=", 4, 11},
		{token.FUNCTION, "fn", 4, 14},
		{token.LPAREN, "(", 4, 15},
		{token.IDENT, "x", 4, 16},
		{token.COMMA, ",", 4, 17},
		{token.IDENT, "y", 4, 19},
		{token.RPAREN, ")", 4, 20},
		{token.LBRACE, "{", 4, 22},
		{token.IDENT, "x", 5, 5},
		{token.PLUS, "+", 5, 7},
		{token.IDENT, "y", 5, 9},
		{token.SEMICOLON, ";", 5, 10},
		{token.RBRACE, "}", 6, 3},
		{token.SEMICOLON, ";", 6, 4},
		{token.LET, "let", 8, 5},
		{token.IDENT, "result", 8, 12},
		{token.ASSIGN, "=", 8, 14},
		{token.IDENT, "add", 8, 18},
		{token.LPAREN, "(", 8, 19},
		{token.IDENT, "five", 8, 23},
		{token.COMMA, ",", 8, 24},
		{token.IDENT, "ten", 8, 28},
		{token.RPAREN, ")", 8, 29},
		{token.SEMICOLON, ";", 8, 30},
		{token.BANG, "!", 9, 3},
		{token.MINUS, "-", 9, 4},
		{token.SLASH, "/", 9, 5},
		{token.ASTERISK, "*", 9, 6},
		{token.INT, "5", 9, 7},
		{token.SEMICOLON, ";", 9, 8},
		{token.INT, "5", 10, 3},
		{token.LT, "<", 10, 5},
		{token.INT, "10", 10, 8},
		{token.GT, ">", 10, 10},
		{token.INT, "5", 10, 12},
		{token.SEMICOLON, ";", 10, 13},
		{token.IF, "if", 12, 4},
		{token.LPAREN, "(", 12, 6},
		{token.INT, "5", 12, 7},
		{token.LT, "<", 12, 9},
		{token.INT, "10", 12, 12},
		{token.RPAREN, ")", 12, 13},
		{token.LBRACE, "{", 12, 15},
		{token.RETURN, "return", 13, 10},
		{token.TRUE, "true", 13, 15},
		{token.SEMICOLON, ";", 13, 16},
		{token.RBRACE, "}", 14, 3},
		{token.ELSE, "else", 14, 8},
		{token.LBRACE, "{", 14, 10},
		{token.RETURN, "return", 15, 10},
		{token.FALSE, "false", 15, 16},
		{token.SEMICOLON, ";", 15, 17},
		{token.RBRACE, "}", 16, 3},
		{token.INT, "10", 18, 4},
		{token.EQ, "==", 18, 7},
		{token.INT, "10", 18, 10},
		{token.SEMICOLON, ";", 18, 11},
		{token.INT, "10", 19, 4},
		{token.NOT_EQ, "!=", 19, 7},
		{token.INT, "9", 19, 9},
		{token.SEMICOLON, ";", 19, 10},
		{token.INT, "10", 20, 4},
		{token.LT_EQ, "<=", 20, 7},
		{token.INT, "11", 20, 10},
		{token.SEMICOLON, ";", 20, 11},
		{token.INT, "11", 21, 4},
		{token.GT_EQ, ">=", 21, 7},
		{token.INT, "10", 21, 10},
		{token.SEMICOLON, ";", 21, 11},
		{token.STRING, "foobar", 22, 4},
		{token.SEMICOLON, ";", 22, 10},
		{token.STRING, "foo bar", 23, 4},
		{token.SEMICOLON, ";", 23, 11},
		{token.EOF, "", 24, 3},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - tokenLine wrong. expected=%d, got=%d", i, tt.expectedLine, tok.Line)
		}
		if tok.Column != tt.expectedColumn {
			t.Fatalf("tests[%d] - tokenColumn wrong. expected=%d, got=%d", i, tt.expectedColumn, tok.Column)
		}
	}
}

package parser

import (
	"fmt"
	"testing"

	"github.com/cupsadarius/monkey_interpreter/ast"
	"github.com/cupsadarius/monkey_interpreter/lexer"
)

func TestIdentifierExpression(t *testing.T) {

	input := "foobar;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program,Statements[0] is not an ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not ast.Identifier, got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not %s, got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral is not %s, got=%s", "foobar", ident.TokenLiteral())
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	ident, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp is not ast.IntegerLiteral, got=%T", il)
		return false
	}

	if ident.Value != value {
		t.Fatalf("ident.Value not %d, got=%d", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Fatalf("ident.TokenLiteral is not %s, got=%s", "5", ident.TokenLiteral())
		return false
	}

	return true
}

func TestIntegerLiteral(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program,Statements[0] is not an ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	if !testIntegerLiteral(t, stmt.Expression, 5) {
		return
	}
}

func testFloatLiteral(t *testing.T, fl ast.Expression, value float64) bool {
	ident, ok := fl.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("exp is not ast.FloatLiteral, got=%T", fl)
		return false
	}

	if ident.Value != 5.1 {
		t.Fatalf("ident.Value not %f, got=%f", 5.1, ident.Value)
		return false
	}

	if ident.TokenLiteral() != "5.1" {
		t.Fatalf("ident.TokenLiteral is not %s, got=%s", "5", ident.TokenLiteral())
		return false
	}

	return true
}

func TestFloatLiteral(t *testing.T) {
	input := "5.1;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program,Statements[0] is not an ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	if !testFloatLiteral(t, stmt.Expression, 5.1) {
		return
	}

}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program,Statements[0] is not an ast.ExpressionStatement, got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not an ast.PrefixExpression, got=%T", stmt.Expression)
		}

    if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, got=%s", tt.operator, exp.Operator)
    }

    if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
      return 
    }
	}
}


func TestParsingInfixExpression(t *testing.T) {
  infixTests := []struct {
  input string
    leftValue int64
    operator string
    rightValue int64
  } {
    {"5 + 5", 5, "+", 5},
    {"5 - 5", 5, "-", 5},
    {"5 * 5", 5, "*", 5},
    {"5 / 5", 5, "/", 5},
    {"5 > 5", 5, ">", 5},
    {"5 < 5", 5, "<", 5},
    {"5 == 5", 5, "==", 5},
    {"5 != 5", 5, "!=", 5},
  }

  for _, tt := range infixTests {
  
    l := lexer.New(tt.input)
    p := New(l)

    program := p.ParseProgram()

    checkParserErrors(t, p)

    if len(program.Statements) != 1 {
			t.Fatalf("program does not contain %d statements, got=%d", 1,  len(program.Statements))
    }

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program,Statements[0] is not an ast.ExpressionStatement, got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not an ast.InfixExpression, got=%T", stmt.Expression)
		}

    if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
      return 
    }
    if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, got=%s", tt.operator, exp.Operator)
    }

    if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
      return 
    }
  }
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}


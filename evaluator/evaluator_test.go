package evaluator

import (
	"testing"

	"github.com/cupsadarius/monkey_interpreter/lexer"
	"github.com/cupsadarius/monkey_interpreter/object"
	"github.com/cupsadarius/monkey_interpreter/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerExpression(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value, got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testFloatExpression(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value, got=%f, want=%f", result.Value, expected)
		return false
	}

	return true
}

func testBooleanExpression(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value, got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestEvalIntegerExpresson(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerExpression(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpresson(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.1", 5.1},
		{"10.01", 10.01},
		{"-5.1", -5.1},
		{"-10.01", -10.01},
		{"5.0 + 5.0 + 5.0 + 5.0 - 10.0", 10.0},
		{"2 * 2 * 2 * 2 * 2.5", 40},
		{"5.0 * 2.5 + 10.0", 22.5},
		{"5.5 + 2.0 * 10.0", 25.5},
		{"20.0 + 2.0 * -10.0", 0.0},
		{"50.0 / 2.0 * 2.0 + 10.0", 60.0},
		{"2.0 * (5.0 + 10.0)", 30.0},
		{"3.0 * 3.0 * 3.0 + 10.0", 37.0},
		{"3.0 * (3.0 * 3.0) + 10.0", 37.0},
		{"(5.0 + 10.0 * 2.0 + 15.0 / 3.0) * 2.0 + -10.0", 50.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatExpression(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpresson(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanExpression(t, evaluated, tt.expected)
	}
}

func TestEvalBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanExpression(t, evaluated, tt.expected)
	}
}

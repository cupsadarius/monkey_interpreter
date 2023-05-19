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
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
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

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
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

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
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

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not Null, got=%T (%+v)", obj, obj)
		return false
	}

	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value, got=%s, want=%s", result.Value, expected)
		return false
	}

	return true
}

func testFunctionObject(t *testing.T, obj object.Object, params []string, body string) bool {
	fn, ok := obj.(*object.Function)

	if !ok {
		t.Fatalf("object is not Function, got=%T (%+v)", obj, obj)
	}

	if len(fn.Parameters) != len(params) {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}
	for index, param := range fn.Parameters {
		if param.String() != params[index] {
			t.Fatalf("parameter at %d is not %q. got=%q", index, params[index], param.String())
		}
	}

	if fn.Body.String() != body {
		t.Fatalf("body is not %q. got=%q", body, fn.Body.String())
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
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalStringExpresson(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"hello world"`, "hello world"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected)
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
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpresson(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"true == false", false},
		{"false == false", true},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 > 2) == true", false},
		{"(1 < 2) == false", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
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
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10.2 }", 10.2},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if v, ok := tt.expected.(int); ok {
			testIntegerObject(t, evaluated, int64(v))
		} else if v, ok := tt.expected.(float64); ok {
			testFloatObject(t, evaluated, float64(v))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"return 10;", 10},
		{"return 1.1;", 1.1},
		{"return 10; 9;", 10},
		{"return 2 * 5;", 10},
		{"return 2 * 5.1;", 10.2},
		{"9; return 2 * 5; 9;", 10},
		{`
      if (10 > 1) {
        if (10 > 1) {
          return 10;
        }

        return 1;
      }
    `, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if v, ok := tt.expected.(int); ok {
			testIntegerObject(t, evaluated, int64(v))
		} else if v, ok := tt.expected.(float64); ok {
			testFloatObject(t, evaluated, float64(v))
		}
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (10 > 1) { true + false; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{`
      if (10 > 1) {
        if (10 > 1) {
          return true + false;
        }

        return 1;
      }
      `, "unknown operator: BOOLEAN + BOOLEAN",
		},
		{"foobar", "identifier not found: foobar"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
		{"let a = 5.1; a;", 5.1},
		{"let a = 5 * 5.1; a;", 25.5},
		{"let a = 5.1; let b = a; b;", 5.1},
		{"let a = 5; let b = a; let c = a + b + 5.1; c;", 15.1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if v, ok := tt.expected.(int); ok {
			testIntegerObject(t, evaluated, int64(v))
		} else if v, ok := tt.expected.(float64); ok {
			testFloatObject(t, evaluated, float64(v))
		}
	}
}

func TestFunctionStatement(t *testing.T) {
	input := "fn(x) { x + 2; };"
	evaluated := testEval(input)

	testFunctionObject(t, evaluated, []string{"x"}, "(x + 2)")
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) {x * 2;}; double(5);", 10},
		{"let add = fn(x, y) {x + y;}; add(5, 5);", 10},
		{"let add = fn(x, y) {x + y;}; add(5 + 5, add(5, 5));", 20},
		{"let identity = fn(x) { x; }; identity(5.1);", 5.1},
		{"let identity = fn(x) { return x; }; identity(5.1);", 5.1},
		{"let double = fn(x) {x * 2;}; double(5.1);", 10.2},
		{"let add = fn(x, y) {x + y;}; add(5.1, 5.1);", 10.2},
		{"let add = fn(x, y) {x + y;}; add(5.1 + 5.1, add(5.1, 5.1));", 20.4},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if v, ok := tt.expected.(int); ok {
			testIntegerObject(t, evaluated, int64(v))
		} else if v, ok := tt.expected.(float64); ok {
			testFloatObject(t, evaluated, float64(v))
		}
	}
}

func TestClosure(t *testing.T) {
	input := `
    let newAdder = fn(x) {
      fn(y) { x + y; };
    };

    let addTwo = newAdder(2);
    addTwo(2);
  `
	testIntegerObject(t, testEval(input), 4)
}

package object

import (
	"fmt"
	"math/big"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	FLOAT_OBJ   = "FLOAT"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Float struct {
	Value float64
}

func (f *Float) Inspect() string  { return fmt.Sprintf("%f", f.Value) }
func (f *Float) Type() ObjectType { return FLOAT_OBJ }

func (f *Float) Add(val *Float) big.Rat {
  a := new(big.Rat).SetFloat64(f.Value)
  b := new(big.Rat).SetFloat64(val.Value)

  result := new(big.Rat)
  return *result.Add(a, b)
}

func (f *Float) Sub(val *Float) big.Rat {
  a := new(big.Rat).SetFloat64(f.Value)
  b := new(big.Rat).SetFloat64(val.Value)

  result := new(big.Rat)
  return *result.Sub(a, b)
}

func (f *Float) Mul(val *Float) big.Rat {
  a := new(big.Rat).SetFloat64(f.Value)
  b := new(big.Rat).SetFloat64(val.Value)

  result := new(big.Rat)
  return *result.Mul(a, b)
}

func (f *Float) Quo(val *Float) big.Rat {
  a := new(big.Rat).SetFloat64(f.Value)
  b := new(big.Rat).SetFloat64(val.Value)

  result := new(big.Rat)
  return *result.Quo(a, b)
}

func FloatFromInteger(obj Object) *Float {
  intVal := obj.(*Integer).Value
	return &Float{Value: float64(intVal)}
}


type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

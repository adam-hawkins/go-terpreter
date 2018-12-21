package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/literallystan/go-terpreter/ast"
)

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

//ObjectType string representation of the object's type
type ObjectType string

//Object provide an interface for all Types to abide by
type Object interface {
	Type() ObjectType
	Inspect() string
}

//Integer represents int type for interpreter
type Integer struct {
	Value int64
}

//Inspect returns the literal value as a string
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

//Type returns the object's Type
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

//Boolean represents bool type for interpreter
type Boolean struct {
	Value bool
}

//Inspect returns the literal value as a string
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

//Type returns the object's Type
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

//Null represents null for interpreter
type Null struct{}

//Inspect returns the literal value as a string
func (n *Null) Inspect() string { return "" }

//Type returns the object's Type
func (n *Null) Type() ObjectType { return NULL_OBJ }

//ReturnValue contains values from return statements
type ReturnValue struct {
	Value Object
}

//Type returns the object's Type
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

//Inspect returns the literal value as a string
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

//Error handles error messages to throw from the interpretor
type Error struct {
	Message string
}

//Type returns the object's Type
func (e *Error) Type() ObjectType { return ERROR_OBJ }

//Inspect returns the literal value as a string
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type Builtin struct {
	Fn BuiltinFunction
}

type BuiltinFunction func(args ...Object) Object

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

//TODO: Cache HashKey() return methods

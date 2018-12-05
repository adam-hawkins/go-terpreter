package object

import (
	"fmt"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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
func (n *Null) Inspect() string { return "null" }

//Type returns the object's Type
func (n *Null) Type() ObjectType { return NULL_OBJ }

package object

//Environment - a map of variable name to values
type Environment struct {
	store map[string]Object
}

//Get gets the Environment's value
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

//Set sets the Environment's value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

//NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

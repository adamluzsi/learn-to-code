package xinterface

// Exercise 1: Define Interface with Method (Basic Interface Definition)
type Describer interface {
	Describe() string
}

// Exercise 2: Implement Interface with Struct (Interface Implementation)
// Return a description of the T.

type T struct {
	Desc string
}

// Exercise 3: Use Interface in Function (Interface as Parameter)
func UseInterfaceMethod(d Describer) string {
	// Call the Describe method on the interface and return its result.
	return ""
}

// Exercise 4: Test Interface Implementation (Interface Validation)
func UseAsAnInterfaceImplementation() Describer {
	// Return a T struct that implements the Describer interface.
	return nil
}

// Exercise 5: Type assert the interface value (Interface Type Assertion)
func InterfaceTypeAssertion(d Describer) bool {
	// Return true if the `d` variable contains a T value type.
	return false
}

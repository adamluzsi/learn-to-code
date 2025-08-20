package xstruct

type Person struct {
	Name string
	Age  int
}

// Exercise 1: Create Struct with Given Fields (Basic Struct Creation)
func CreateStruct() Person {
	// Create a Person struct with name "Alice" and age 30.
	return Person{}
}

// Exercise 2: Create Struct with Given Fields (Basic Struct Creation)
func NewStruct() *Person {
	// Create a &Person struct with name "Zoe" and age 24.
	// and do it in one go by returning a pointer to it "&Person{...}"
	return nil
}

// Exercise 3: Add a new field called "MyField" on Person struct.
// The type of this field is up to you.

// Exercise 4: Modify Age via Value Receiver (Value Receiver Method)
// Expected method name is "WithAge" with age int as its only argument.
// It should modify the age of the Person by its argument and return the copy of the receiver's Person value.

// Exercise 5: Modify Age via Pointer Receiver (Pointer Receiver Method)
// Expected method name is "UpdateAge" with age int as its only argument.

// Exercise 6: Create your very own struct using the `type` and `struct` keywords,
// and return a value of it in this function.
func ShowYourOwnType() any {
	return nil
}

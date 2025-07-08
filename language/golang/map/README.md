# Map Exercise

## What is a Map?

A map, also known as an associative array or dictionary, is a built-in Go data type that stores key-value pairs. It allows you to look up values by their corresponding keys.

Think of a map like a phonebook. You have names (keys) and phone numbers (values). When you want to find someone's phone number, you look them up in the phonebook using their name as the key.

**How does it work?** (simplified)

1. **Keys**: Each entry in a map has a unique key. Keys are used to identify and retrieve values from the map.
2. **Values**: Each key is associated with a value. Values can be any type of data (e.g., strings, integers, structs).
3. **Lookups**: When you want to find a value, you use its corresponding key to look it up in the map.

**Real-world Example**

Let's say we're building an online shopping cart application. We need to store information about each product, such as its name, price, and quantity. A map is perfect for this because we can use the product ID (a unique string) as the key and store the corresponding product information as the value.

example:
```go
m := make(map[string]string)

// Add some entries to the map
m["foo"] = "Lorem ipsum dolor sit amet"
// -> map[string]string{"foo":"Lorem ipsum dolor sit amet"}

m["bar"] = "consetetur sadipscing elitr"
// -> map[string]string{"foo":"Lorem ipsum dolor sit amet", "bar":"consetetur sadipscing elitr"}

m["baz"] = "sed diam nonumy eirmod tempor" 
// -> map[string]string{"foo":"Lorem ipsum dolor sit amet", "bar":"consetetur sadipscing elitr", "baz":"sed diam nonumy eirmod tempor"}
```

In this example, we use the keys `"foo"`, `"bar"`, and `"baz"` to store the corresponding text values. We can then look up the value for a given key using `m[key]`.

For example:
```go
fmt.Println(m["foo"]) // Output: Lorem ipsum dolor sit amet
```

## Hands-on Practice

In this exercise, you'll practice working with maps in Go.
Open the `map.go` file and complete the exercises. 

Replace each occurrence of `panic("implement-me")` with a proper implementation.

Then, run `go test` to verify your work using the tests in `map_test.go`.

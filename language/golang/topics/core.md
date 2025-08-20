# Core Language Concepts of Go

## Variables

Go uses the `var` keyword or short declaration `:=` to create variables.

```go
var x int           // declares x of type int, initialized to 0
var s string = "hi" // declares s with explicit initializer
y := 42             // short declaration; type inferred as int
```

You can declare multiple at once:

```go
var a, b int = 1, 2
c, d := "foo", 3.14
```

### Constants

Constants are declared with `const`. Their value must be computable at compile time.

```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

### Primitive Types

- **Booleans**: `bool` (`true`, `false`)
- **Integers**: `int`, `int8`, `int16`, `int32`, `int64` (and unsigned variants: `uint`, `byte`, `rune`)
- **Floating-point**: `float32`, `float64`
- **Complex**: `complex64`, `complex128`
- **Strings**: immutable sequences of bytes (`string`)

### Basic Operations

- **Arithmetic**: `+`, `-`, `*`, `/`, `%`
- **Comparison**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Logical**: `&&`, `||`, `!`
- **Bitwise**: `&`, `|`, `^`, `<<`, `>>`
- **String concatenation**: `+`
- **Zero values**: Variables default to zero value (e.g., `0` for numbers, `""` for strings, `false` for booleans).

```go
a := 10
b := 3
sum := a + b         // 13
equal := (a == b)    // false
notEqual := a != b   // true
str := "Hello, " + "Go"
```

## Control Structures

### `if-else`

```go
if n%2 == 0 {
    fmt.Println("even")
} else if n%3 == 0 {
    fmt.Println("divisible by 3")
} else {
    fmt.Println("odd")
}
```

**Optional short statement**:

```go
if err := doSomething(); err != nil {
    log.Fatal(err)
}
```

### Loops

Go has only one loop construct, `for`.

```go
// Classic
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// As while
sum := 1
for sum < 1000 {
    sum += sum
}

// Infinite
for {
    // break to exit
}

// For with multiple variables
for a, b := 0, 1; a < 10; a, b = b, a+b {
    fmt.Println(a)
}
```

### `switch`

```go
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println("Weekend")
case time.Monday:
    fmt.Println("Start of week")
default:
    fmt.Println("Midweek")
}

// Expressionless switch (like if-else chain)
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}
```

### `range`

Iterate over slices, arrays, maps, strings, channels.

```go
// Slice
nums := []int{2, 4, 6}
for i, v := range nums {
    fmt.Printf("%d: %d\n", i, v)
}

// Map
m := map[string]int{"a": 1, "b": 2}
for key, val := range m {
    fmt.Println(key, val)
}

// String (rune-wise)
for i, r := range "Go" {
    fmt.Println(i, r)
}

// Only value
for _, v := range nums {
    fmt.Println(v)
}
```

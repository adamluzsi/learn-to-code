
Ensure `go version` prints your installed Go version and `GOPATH`/`GOROOT` are set.

- **Workspace & Modules**  
   - Create a project directory anywhere outside `$GOPATH/src`.  
   - Initialize a module:  
     ```bash
     mkdir ~/projects/myapp
     cd ~/projects/myapp
     go mod init example.com/myapp
     ```
     This creates `go.mod` and enables dependency management[1].

- **Editor Configuration**  
   - Use VS Code with the **Go** extension or GoLand.  
   - Enable auto-format (`gofmt`), linting (`golangci-lint`), and autocomplete.

## 2. Go Syntax & Basics

### 2.1 File Structure  
Every Go program consists of:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
- `package main`: entry point package.  
- `import`: brings in standard or third-party packages.  
- `func main()`: program’s entry function[2].

### 2.2 Variables & Types  
```go
var age int = 30
name := "Alice"       // shorthand
const Pi = 3.14
```
- Static typing; shorthand `:=` infers type.  
- Common types: `int`, `float64`, `string`, `bool`.

### 2.3 Statements & Formatting  
- Statements end implicitly at line break; use braces `{}` on the same line for blocks.  
- Always run `gofmt` to enforce canonical style[2].

## 3. Control Structures

Go has three primary constructs: `if`, `for`, and `switch`.

### 3.1 `if` / `else`  
```go
if score := 85; score >= 60 {
    fmt.Println("Pass")
} else {
    fmt.Println("Fail")
}
```
- Optional initializer before condition.  
- No parentheses around conditions[3].

### 3.2 `for` (the only loop)  
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

for _, v := range []int{10, 20, 30} {
    fmt.Println(v)
}

for condition {
    // like while
}
```
- Classic, range-based, and condition-only forms.

### 3.3 `switch`  
```go
switch day {
case "Mon":
    fmt.Println("Monday")
case "Tue", "Wed":
    fmt.Println("Midweek")
default:
    fmt.Println("Weekend")
}
```
- No implicit fallthrough; use `fallthrough` explicitly[3].

## 4. Functions

Functions encapsulate reusable logic.

### 4.1 Declaration & Invocation  
```go
func greet(name string) string {
    return "Hello, " + name
}

func main() {
    fmt.Println(greet("Bob"))
}
```
- Begin with `func`, name, parameters, return types, and body in `{}`[4].

### 4.2 Multiple Returns & Error Handling  
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divide by zero")
    }
    return a / b, nil
}
```
- Return values and errors explicitly.

## 5. Core Data Structures

### 5.1 Arrays  
```go
var arr [3]int = [3]int{1, 2, 3}
```
- Fixed size.

### 5.2 Slices  
```go
s := []int{10, 20, 30}
s = append(s, 40)         // dynamic growth
fmt.Println(len(s), cap(s))
sub := s[1:3]             // slicing
```
- Built atop arrays; dynamic length and capacity[5].

### 5.3 Maps  
```go
m := map[string]int{"a": 1, "b": 2}
m["c"] = 3
delete(m, "b")
for k, v := range m {
    fmt.Println(k, v)
}
```
- Unordered key-value store.

## 6. Packages & Modules

Organize code across files and projects.

- **Packages**  
   - Each directory defines a package.  
   - Export identifiers by capitalizing names.
- **Modules**  
   - Rooted by `go.mod`.  
   - Manage dependencies via `go get`, `go mod tidy`[1].

Use `import "example.com/myapp/utils"` to include sibling packages in your module.

## 7. Hands-On Practice

- **“Hello World” Variation**  
   - Read a name from `os.Args` and greet the user.
- **Mini CLI Tool**  
   - Implement a simple calculator: parse flags for operations (`add`, `sub`, etc.).
- **Slice & Map Exercises**  
   - Write functions to count word frequencies in a string (use map).  
   - Implement a dynamic buffer using slices.
- **Module Project**  
   - Split the calculator into `main` and a `math` package.  
   - Add unit tests using Go’s `testing` package.

## 8. Recommended Resources

- The Go Tour (tour.golang.org) for interactive basics.  
- “Effective Go” for idiomatic patterns[6].  
- Go by Example (gobyexample.com) for annotated examples[7].

**Key Takeaway:** Focus on writing small, self-contained programs to internalize syntax, control flow, and modular code organization. Practice daily by building CLI tools, exploring slices/maps, and structuring projects into packages and modules.

Sources
[1] An introduction to Packages, Imports and Modules in Go https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules
[2] Go Syntax - W3Schools https://www.w3schools.com/go/go_syntax.php
[3] Control Structures in Go | Developer Portal https://tutorials.cosmos.network/tutorials/4-golang-intro/4-control.html
[4] Go Functions - W3Schools https://www.w3schools.com/go/go_functions.php
[5] Slices - KodeKloud Notes https://notes.kodekloud.com/docs/Golang/Arrays-Slices-and-Maps/Slices
[6] Effective Go - The Go Programming Language https://go.dev/doc/effective_go
[7] Go by Example https://gobyexample.com
[8] An Introduction to the Go Programming Language: Setup, Syntax ... https://dev.to/ankitakanchan/an-introduction-to-the-go-programming-language-setup-syntax-and-basic-concepts-53gh
[9] Understanding Go's control structures: if, for, and switch - Reintech https://reintech.io/blog/understanding-gos-control-structures-if-for-switch
[10] Go Functions (With Examples) - Programiz https://www.programiz.com/golang/function
[11] Golang Tutorial for Beginners | Full Go Course - YouTube https://www.youtube.com/watch?v=yyUHQIec83I
[12] Mastering Go's Control Flow and Functions: Your Path to Gopher ... https://nikhilakki.in/mastering-gos-control-flow-and-functions-your-path-to-gopher-greatness
[13] Mastering Functions in Go: A Beginner's Guide to Advanced Topics https://blog.stackademic.com/mastering-functions-in-go-a-beginners-guide-to-advanced-topics-5561c4432878
[14] How to Learn Golang – A Beginner's Guide to the Basics https://www.freecodecamp.org/news/golang-for-beginners/
[15] Introduction to Control Structures in Go | CodeSignal Learn https://codesignal.com/learn/courses/mastering-control-structures-in-go/lessons/introduction-to-control-structures-in-go
[16] Learn Functions in Go with Examples | golangbot.com https://golangbot.com/functions/
[17] Tutorials - The Go Programming Language https://go.dev/doc/tutorial/
[18] Mastering Control Structures in Golang: A Comprehensive Guide https://www.linkedin.com/pulse/mastering-control-structures-golang-comprehensive-guide-singh
[19] Go Functions and Methods | Golang 101 Series | Go Tutorial https://www.youtube.com/watch?v=GEvnSEUvp2w
[20] Tutorial: Get started with Go - The Go Programming Language https://go.dev/doc/tutorial/getting-started
[21] Control Statements - Control Your Go - DEV Community https://dev.to/jankaritech/control-statements-control-your-go-1om7
[22] Functions in Go (Golang) - Udacity https://www.udacity.com/blog/2023/05/functions-in-go-golang.html
[23] Go Modules: A Beginner's Guide. - DEV Community https://dev.to/kingkunte_/go-modules-beginners-guide-4a7p
[24] How to Iterate Over Arrays, Slices and Maps in Go - Stackademic https://blog.stackademic.com/how-to-iterate-over-arrays-slices-and-maps-in-go-b4f09a5b2158
[25] Go packages, Go modules and init functions tutorial | golangbot.com https://golangbot.com/go-packages/
[26] Mastering Collections in Go: A Comprehensive Guide to Arrays ... https://dev.to/avinashtechlvr/mastering-collections-in-go-a-comprehensive-guide-to-arrays-slices-and-maps-39m
[27] Introduction To Golang: Arrays, Maps, Slices And Make https://www.openmymind.net/Introduction-To-Go-Arrays-Maps-Slices-And-Make/
[28] Tutorial: Create a Go module - The Go Programming Language https://go.dev/doc/tutorial/create-module
[29] Go Beginners Series: Collection Types, Arrays, Slices, and Maps https://olodocoder.hashnode.dev/go-beginners-series-collection-types-arrays-slices-maps
[30] Using Go Modules - The Go Programming Language https://go.dev/blog/using-go-modules
[31] Advance Data Types in Go: Arrays, Slices, Maps, Functions https://codewithintellect.hashnode.dev/advance-data-types-in-go-arrays-slices-maps-functions
[32] How to Use Go Modules - DigitalOcean https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
[33] Arrays, Slices and Maps in Go - Go 101 https://go101.org/article/container.html
[34] Go modules and packages — which one to use and when - YouTube https://www.youtube.com/watch?v=EdV1rx5613g
[35] Arrays, Slices, Maps - Go Programming Tutorial #3 - YouTube https://www.youtube.com/watch?v=JfT7-JbjfVo
[36] Understanding Go Modules for Beginners - TiDB https://www.pingcap.com/article/understanding-go-modules-for-beginners/
[37] How do packages and modules work? : r/golang - Reddit https://www.reddit.com/r/golang/comments/wtp10l/how_do_packages_and_modules_work/

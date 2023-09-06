# A. Language Track

- [Start](README.md)
- [Next](B_architecture_track.md)

---

## Intro

| Topic | Link |
| ----------- | ----------- |
| Hello World | [gobyexample.com/hello-world](https://gobyexample.com/hello-world) |
| Values |  [gobyexample.com/values](https://gobyexample.com/values) |
| Variables |  [gobyexample.com/variables](https://gobyexample.com/variables) |
| Constants |  [gobyexample.com/constants](https://gobyexample.com/constants) |
| Structs | [gobyexample.com/structs](https://gobyexample.com/structs) |
| Pointers |  [gobyexample.com/pointers](https://gobyexample.com/pointers) |

Please understand each of the links above before proceeding.

Each section of code in gobyexample has a copy button, you can run it on Go playground: [go.dev/play](https://go.dev/play)

## Case has some meaning in Go

Exported variables and functions within a package have an initial **uppercase** letter, making them accessible from outside the package.

### Private to package

In Go, objects (variables, types, and functions) with lowercase initial letters are considered private to the package they are defined in and cannot be accessed from outside that package.

```go
package firstPackage

// Objects with lowercase initial letters are private to this package.
var myInt int
type myStruct string
func myFunc() { ... }
```

```go
package secondPackage

// It is not possible to reference lowercase objects from another package.
// Example attempts that will not work:
// a := firstPackage.myInt
// b := firstPackage.myStruct{}
// firstPackage.myFunc()
```

### Exported from package

In Go, objects (variables, types, and functions) with uppercase initial letters are considered exported and can be accessed from other packages.

```go
package firstPackage

// Objects with uppercase initial letters are exported and accessible from other packages.
var MyInt int
type MyStruct struct { ... }
func MyFunc() { ... }
```

```go
package secondPackage

// This means we can instantiate objects and call functions from the other package.
a := firstPackage.MyInt
b := firstPackage.MyStruct{}
firstPackage.MyFunc()
```

## Examples

### Hello World

- See [hello_world](A/01_hello_world.go)
- Go playground: [https://go.dev/play/p/eGReLQib29N](https://go.dev/play/p/eGReLQib29N)
- ```go run A/01_hello_world.go```

### Constants, Values & Variables

- code: [constants_values_and_variables](А/02_constants_values_and_variables.go)
- go playground: [https://go.dev/play/p/I_Oil4wQOJK](https://go.dev/play/p/I_Oil4wQOJK)
- ```go run A/02_constants_values_and_variables.go```

### Structs and pointers

- code: [structs_and_pointers](A/03_structs_and_pointers.go)
- go playground: [https://go.dev/play/p/5s7VBFnQupB](https://go.dev/play/p/5s7VBFnQupB)
- ```go run A/03_structs_and_pointers.go```

---

- [Start](README.md)
- [Next](B_architecture_track.md)
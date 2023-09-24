# Language Track

- [Start](README.md)
- [Next](B_architecture_track_02.md)

---


## Intro

| Topic | Link |
| ----------- | ----------- |
| For | [gobyexample.com/for](https://gobyexample.com/for) |
| If/Else | [gobyexample.com/if-else](https://gobyexample.com/if-else) |
| Switch | [gobyexample.com/switch](https://gobyexample.com/switch) |
| Arrays | [gobyexample.com/arrays](https://gobyexample.com/arrays) |
| Slices | [gobyexample.com/slices](https://gobyexample.com/slices) |
| Maps | [gobyexample.com/maps](https://gobyexample.com/maps) |
| Range | [gobyexample.com/range](https://gobyexample.com/range) |
| Context | [gobyexample.com/context](https://gobyexample.com/context) |

Please understand each of the links above before proceeding.

Each section of code in gobyexample has a copy button, you can run it on Go playground: [go.dev/play](https://go.dev/play)

## Examples

### Flow control

- Code: [A/01_looping/loop.go](A/01_looping/loop.go)
- Go playground: [https://go.dev/play/p/RqzPJlx6yUq](https://go.dev/play/p/RqzPJlx6yUq)
- ```go run A/01_looping/loop.go```


### Conditionals

- Code: [A/02_conditionals/conditionals.go](A/02_conditionals/conditionals.go)
- Go playground: [https://go.dev/play/p/N9mEeJwcH3B](https://go.dev/play/p/N9mEeJwcH3B)
- ```go run A/02_conditionals/conditionals.go```

### Arrays, Slices, Maps and Range

- Code: [A/03_arrays_slices_maps/arrays_slices_maps.go](A/03_arrays_slices_maps/arrays_slices_maps.go)
- Go playground: [https://go.dev/play/p/oh7WK19pznj](https://go.dev/play/p/oh7WK19pznj)
- ``go run A/03_arrays_slices_maps/arrays_slices_maps.go```

## Context

Use of [Context](https://go.dev/src/context/context.go) is key to Go programming.  Context is just a special type of variable, mainly used for two use cases

- WithValue allows the wrapped transmission of constants/variables
- WithCancel, WithDeadline, and WithTimeout functions support the instant termination of invoked functions or goroutines

Contexts also help us avoid the scalability trap of global package variables. For example, passing a pointer to your configuration/settings via Context is a great mechanism that avoids hardwiring global variables througout your code.

---

- [Start](README.md)
- [Next](B_architecture_track.md)
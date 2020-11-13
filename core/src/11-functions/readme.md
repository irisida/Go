![](/assets/gologo.png)

# The Go programming Language - Functions

A function is a block of code that can be executed and/or re-executed and given a set of inputs will yield the same outputs each time. in Go functions names are typically done in a camelCaseStyle and within a package function names must be unique. One feature of the Go language is that functions can return multiple values. There is no support for function overloading and the `main()` and `init()` functions are part of the language genetics and are therefore predefined.

Syntactically it follows:

```go
func (r receiver) functionName(parameters) (return types) {
    // code of the function body
}
```

## Some basic functions and function signatures

![](/assets/core/11/1101-funcs.png)

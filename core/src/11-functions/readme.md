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

Let's see some basic function signatures, from the trivial no parameters procedure type call, to parameterized functions returning a value, or multiple values.

![](/assets/core/11/1101-funcs.png)

Next we can have named returns, notice the named return in the function signature and that allows us simplify the actual returning of values in the function down to a simple `return` statement. Not only is this is this elegant, it's syntactically beautiful in my opinion.

![](/assets/core/11/1102-named-returns.png)

Variadic functions take a variable number of parameters by using the ellipsis parameter controller `...`, this allows a function to take zero, one or more arguments. However, where a function takes multiple parameters or varied types only the last parameter is permitted to be variadic. Given this can seem quite loosely controlled and the inevitable `what is the usecase?` for them we can typically say the usecases are:

1. where the number of parameters are unknown
2. where you don't want to create a temporary slice just to pass it to a function.
3. readability - it is infinitely easier to read code with one parameter than with 3, 5, 10 and so on...

![](/assets/core/11/1103-variadic.png)

Where variadic functions mutate a slice we can see that the values are passed by reference and not by value as the original structure outside of the variadic function will reflect the mutation thereafter.

![](/assets/core/11/1104-more-variadic.png)

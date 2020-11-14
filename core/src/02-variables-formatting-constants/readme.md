![](/assets/gologo.png)

# The Go programming Language - Variables, formatting & constants

Firstly we'll see examples of static typing where we demonstrate that types cannot leap to other types and that once a type is set it is fixed, reassigning to a new value that would infer a different type typically will keep the original underlying type. We see that with assigning what looks to be an int value to a `float64` variable.

![](/core/src/02-variables-formatting-constants/assets/201-static-typing.png)

It's important to understand the difference between inference and type assignation. Where we declare a variable and type, a value is useful but not a necessity because memory for the declared type can be set aside and updated later with a value of that type safely. Where we don't provide values Go assigns the default empty, or zero value. We can see here that we are declaring the types but no values. If we then use these variables without assigning an actual value they will hold the default zero value associated to their type.

![](/core/src/02-variables-formatting-constants/assets/202-zero-values.png)

When dealing with types, especially numeric types it's useful to be clear that formatting types is not the same thing as changing types. We see in the below sample that we have a static type that is able to be formatted for display in many ways.

![](/core/src/02-variables-formatting-constants/assets/203-formatting-types.png)

We should also understand the subtle differences in Go regarding constants. It can be said that all basic literals are constants, 5 will always be 5, a literal string value "Hi there" cannot be changed. We can therefore say that all literal are just unnamed constants. In Go a constant is not allowed to change during the running of the program. They are immutable.

1. You cannot change a constant
2. You cannot initiate a constant at runtime
3. You cannot use a variable to initialise a constant

![](/core/src/02-variables-formatting-constants/assets/204-constants.png)

Next we can show the use of a way to use a predefined integer sequence for enumerating constants. The `iota`. This is a zero indexed capability that lets us assign a sequence value to constant values that are group. redeclaring the `const` keyword in your program will reset the index of the iota and it will start a zero again.

![](/core/src/02-variables-formatting-constants/assets/205-iota.png)

A somewhat real world example could be enumerating days of the week, or indeed skipping closed days. Assume we had two requirements:

1. show all days of the week enumerated.
2. show open days enumerated, but is closed on Sunday and Wednesday

![](/core/src/02-variables-formatting-constants/assets/205-iota-part-two.png)

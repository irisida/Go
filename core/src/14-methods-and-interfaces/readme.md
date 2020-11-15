![](/assets/gologo.png)

# The Go programming Language - Methods & Interfaces

## Receiver functions

A receiver function is akin to a class method in a true OOP language, in Go its a function that takes an extra parameter to attach/align itself to a type. We can see built-in receiver functions in the standard library with the time package.

![](/core/src/14-methods-and-interfaces/assets/1401-builtin.png)

With programmer defined type we add to the function signature an additional parameter before the function name to show which type it is operating on. The sig would then look something like: `func (name type)functionName(parameters type) return type {}`

![](/core/src/14-methods-and-interfaces/assets/1402-receivers.png)

## Receiver functions with pointers

To be able to make changes to an an example of a type or struct we would use pointers as the receiver type. If we dont we have the same pass by value, local copy and changes lost issue we seen in the pointers examples earlier.

![](/core/src/14-methods-and-interfaces/assets/1403-pointer-receivers.png)

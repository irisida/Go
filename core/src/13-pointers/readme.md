![](/assets/gologo.png)

# The Go programming Language - Pointers

Pointers are a traditionally a "biggie"! People get stuck in languages like C, C++ and subsequently Go because pointers are a tougher concept at face value and therefore have a reputation of complexity. The reality is that the concept is reasonably simple enough and the while it's true that the syntax looks a little funkier it's usually that which throws people or puts a barrier to understanding.

- The RAM, computer memory, can and should be thought of as a series of boxes, cells, locations, pigeonholes, cupboards... whichever terms makes it most easy to conceptualise for you.
- Each cell, box, whatever.. is labelled with a unique number which increments sequentially, this is the memory's _location_...
- These addresses aren't 1,2,3 etc, they're typically longer hexadecimal(base 16) values and therefore they look ugly and frightening.

So we're going to think of that as memory locations as `street addresses` and the CPU is a `Taxi/Uber driver`. Each house holds a person(value) that is stored(sleeps) at that location(address). Follow? In our programs a variable is just a name/nickname for a location, or street address.

**What's any of this got to do with pointers?** well, a pointer is a variable where the value is an address of another piece of memory, or an address you visit to find out another address elsewhere.

![](/core/src/13-pointers/assets/1301-pointers.png)

There are multiple ways you can declare a pointer including uninitialised techniques which result in a `nil` pointer, the `new()` method which again will be uninitalised and nil. Let's see some of these in action. We can also see the use of the `&` address-of and the `*` dereferencing operator in action.

![](/core/src/13-pointers/assets/1302-declare-pointers.png)

WAIT! we've seen `var ptr1 *float64` and later `*p` as having different meanings so what does it mean?

- When you see the `*` before the type it means what is being declared is a pointer to that type, so in the case of `var ptr1 *float64` we're saying that ptr1 is a pointer something of the type float64, therefore whe can assign that to hold the address of a variable that actually is of type float64, no other type.
- When we see the `*` before the pointer variable it means it is the dereferencing of that pointer, in otherwords the underlying value held at the address that the pointer points to.

**These are important.** So let's see that compounded with another example.

![](/core/src/13-pointers/assets/1303-pointers.png)

## A pointer to a pointer

We may also have a pointer that points to a pointer, effectively it is two steps removed and its value will be the memory address of the pointer to the source. To dereference we can use the `**` operator, so we are dereferencing the dereference. See below for the creation and mutation of source values using a pointer to a pointer.

![](/core/src/13-pointers/assets/1304-p2p.png)

An important note for comparing pointers is that they can be compared with a boolean comparison and will only return true if they are pointed to the same memory location, and therefore the same value. Two identical values at different memory location will return false in a comparison.

## Using pointers with functions

Next we need to see how we handle pointers with functions. Go is a pass by value language and that means local copies are made and function scope for values is the typical. We can bypass that with pointers to be able to **appear to** pass by reference. This is not strictly the case as there is no _pass by reference_ in Go but Let's see a direct example that shows both how values are localised by default and how a pointer can be passed to a function that performs a mutation and the changes are reflected after the function has completed and returned.

![](/core/src/13-pointers/assets/1305-pointer-funcs.png)

## Pass by value Vs pass by pointer

As we have noted above, in a lot of programming languages when a value is passed by a pointer it is documented as pass by reference. There is no _pass by reference_ in Go and each time you pass an argument to a function, even a pointer, it will be copied to a different memory location. To understand it better let's have a look at an example of each data type when passed to a function that tries to modify the arguments inside its body.

![](/core/src/13-pointers/assets/1306-pointer-vs-values.png)

![](/core/src/13-pointers/assets/1307-pointers-vs-value-struct.png)

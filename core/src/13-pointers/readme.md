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

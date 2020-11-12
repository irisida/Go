![](/assets/gologo.png)

# The Go programming Language - Slices

A slice is a Go specific feature and works pretty much like a dynamic array in other languages. Where an array in Go is of fixed length and determined at compile time. A slice is dynamic, it can shrink or grow, length is not part of its type and belongs to runtime. An uninitialized slice is equal to `nil`. The zero value is nil.

Where it has commonality with an array is that both constructs demand all elements are of the same type. Both construct can use the `keyed` mechanism which allows us to specify a key as its index position.

![](/assets/core/06/06-601-slice-basics.png)

## Comparing slices in Go

Slices can be compared to nil in a simple comparison operation, `data != nil`, or `data == nil`. Comparing slices with value sets is not so trivial or syntactically easy. any direct comparison operation will generate a compile time error. Instead we must compare element by element and be mindful of lengths between structures.

![](/assets/core/06/06-602-compare-slices.png)

## Copying with slices

Copying slices can be done as an iteration of with the copy function, but the copy function does not demand the length of the source and destination be equal, in fact it's entirely possible to successfully transfer nothing to a zero length destination. So the copy operator requires a little care to be sure the results are intended.

![](/assets/core/06/06-603-copy-slices.png)

## Slice of a slice

We can create a new slice from an existing slice, or we can have a slice of a slice as a slice. I was actually going to try to work tge word slice in another couple of times there to be a smart-arse but you get the point.

![](/assets/core/06/06-604-slice-of-slice.png)

So now we've seen a slice what is it? Well a slice has some under the hood magic going on. Underneath, Go creates what is called a `backing array` and it's the backing array that stores the elements, not the slice. Go implements a slice as a data structure called a `slice header`. A slice header has three fields:

1. `The address` of the backing array (pointer)
2. `the length` of the slice. returned by `len()`
3. `the capacity` of the slice. returned by `cap()`

- The slice header is a runtime representation of slice.
- A nil slice does not have a backing array while it remains a nil slice.
- creating a slice from a slice uses the underlying backing array as the source for both as an built-in optimisation. mutation in one will reflect on the other. The slice header will be different and the address of the slice header will be different. The underlying array will be the same.

## Cap & len demo

The underlying array and how `cap()` and `len()` interact with it show some inefficiencies in Go, or optimisations depending on how you look at it.

![](/assets/core/06/06-605-burst-cap.png)

The output of this is as follows:

```shell
len: 0  cap: 0
len: 2  cap: 2
len: 3  cap: 4
len: 4  cap: 4
len: 5  cap: 8
```

Note that the first time it bursts the capacity we create a new underlying backing array with a higher capacity, but now we see we have a length of 3 and a capacity of 4 so we have allocated one space more than we need, this is an inefficiency. On the other hand creating that array and allocating it is more expensive than assigning a unit of excess space. The second time we burst the capacity it has again over-allocated based on the pattern of usage it can see. Where you stand on this space inefficiency vs operational efficiency is entirely up to you, nonetheless it remains an impressive feature of slices in Go.

![](/assets/gologo.png)

# The Go programming Language - Maps

- A map is a collection, much like an array or slice, but is constructed from `key:value` pairs.
- performance advantage of maps is that add, delete and get tale constant time.
- all keys and values are statically typed.
- keys must be unique, values may repeat.
- Any comparable type can be used as a key (facilitates comparison with double equals operator).
- can only compare a whole map to nil, not to another map.
- they are unordered.

Let's see some basics and a quick coverage of the `comma ok idiom` in Go.

![](/assets/core/08/08-801-map-basics.png)

To compare a map with another map we must create a string representation of it as natively the only comparison supported is to nil, armed with the strings we can check equality and operate conditionally.

![](/assets/core/08/08-802-map-compare.png)

To clone, or deep copy a map we should use the make() to create the map type we need and loop across the source to be copied element by element otherwise we will have the side effect from a direct copy where the underlying memory location and structure are the same meaning a mutation in one is reflected in the other.

![](/assets/core/08/08-803-clone-map.png)

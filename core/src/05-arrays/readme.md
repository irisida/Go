![](/assets/gologo.png)

# The Go programming Language - Arrays

An array is a composite type that is indexed. It stores a collection of elements which must be of the same type. Internally Go store arrays in contiguous memory so it is considered to be a very memory efficient structure. An array is defined by both its length and element type, these must be known at compile time.

We can declare arrays in multiple ways but the essence of compile time fixed sizing remains.

![](/assets/core/05/05-501-array-basics.png)

We can see standard ops on an array such as `len()` functions to use the length.

![](/assets/core/05/05-502-array-ops.png)

It is also possible for there to be multiple dimensions to an array. We should remember though that when copying an array it is a copy by value and not by reference. There is an implicit deep copy with arrays that is not quite the case with `slices` as we will see in future examples. Slices implicitly use a shallow copy, or copy by reference.

![](/assets/core/05/05-503-array-dims.png)

We can see below it is also possible to have the programmer define the indices with a mechanism known as keyed values.

![](/assets/core/05/05-504-keyval.png)

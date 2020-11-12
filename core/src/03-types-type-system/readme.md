![](/assets/gologo.png)

# The Go programming Language - Types and the Type system

## Built-in types

A type determines a set of values together with operations and methods specific to those values. In Go there are predeclared types, introduced types by means of programmer defined type declarations as well as composite type such as: array, slice, map, struct, pointer, function, interface and channel.

Go will gracefully handle the event of a type overflow. At runtime where a types boundaries are burst type will not crash out. If you break the boundary of the max value, it will boomerang to the min value, if you break the min it will bounce to the max.

![](/assets/core/03/03-301-overflow.png)

Next we can see new types `rune` and `byte`, in Go these are aliases for int32 in the case of rune and uint8 where it is byte.

![](/assets/core/03/03-302-rune-byte.png)

Boolean types, or `bool` are standard as with almost every programming language and have exclusively true/false values. The `string` type is a sequence of unicode characters enclosed i double quotes. Arrays, slices and maps will be covered in greater detail but a summary of what they are could be:

- `array` - numbered sequence of elements of a single type where the size is fixed.
- `slice` - a Go specific feature implementation of dynamic arrays that can shrink and grow in size.
- `map` - unordered group of elements of one type indexed by unique keys of another type. Most similar mainstream construct would be a dictionary in python.
- `struct` - a sequence of named elements, or fields. Each field must have a name and a type. Closest comparable construct would be a class in a pure OOP.
- `pointer` - A pointer is a variable that stores the memory address of another variable. The value of an uninitlalised pointer is nil.
- `func` - a function type.

Conversion isn't so much as quirky, as it is formal. Some languages will accept a loose or lossy conversion. Even for something simple like printing as a string Go likes to be quite specific. If you try to basically convert and int to a string by wrapping itu will return a rune (unicode character value for that number). You should lean towards the `Sprintf` of the fmt package to achieve these conversions.

![](/assets/core/03/03-303-conversion.png)

We can also convert from strings to numeric types although we must be quite specific about how. We can use the `strconv` methods for floats, bools, ints and uints.

![](/assets/core/03/03-304-conversion-pt2.png)

## Named types

A defined type is also called a `named type`. A named type is where we have a new type that is created by the programmer from an existing source, or base, type. These new types must have a name, it possible for them to have 0, 1 or many methods associated with them. What is interesting in comparison to `rune` and `byte` is that when we check the type of a created type it os not shown as a representation of its underlying type, it is show as of the type created. If you check if the new types _type_ and the underlying type they are not the same. Where you want to perform operation of one on the there there must be an explicit conversion to the required type.

lastly, but quite importantly, there is no type hierarchy in go, a type is a type and some might even say type is life....

![](/assets/core/03/03-305-named-types.png)

We can also attach methods to named types as mentioned above, therefore we can extend functionality and create a Go-centric simplicity and readability around our types and the operations we can do on them.

![](/assets/core/03/03-306-named-type-methods.png)

finally we have aliasing, where instead of declaring a type itself we assign it a named type, this creates an exact same type but with a new name. Aliased types are exactly what `rune` and `byte`, mentioned earlier, are. They are int32 and uint8 types underneath.

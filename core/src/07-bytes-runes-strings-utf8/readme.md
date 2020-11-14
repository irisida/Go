![](/assets/gologo.png)

# The Go programming Language - Bytes, runes, strings & utf-8

## Strings

In Go strings by default are utf-8 encoded, what can we say except that one of the creators of Go was also responsible for utf-8. Strings are enclosed din double quotes.

![](/core/src/07-bytes-runes-strings-utf8/assets/701-strings.png)

strings are synonymous with char types in many programming languages. Go does not have a cha type, instead Go uses Bytes and Runes to represent characters. Characters are often called `rune literals` and are expressed in Go by enclosing them in single quotes. Rune literals are represented using Unicode Code Points. A code point is a numeric value that represents a rune literal.

- ASCII is a subset of Unicode is made up of 128 code points.
- In Go a string is a series of byte values.
- A string is a slice of byte and any byte slice can be encoded to a string
- A code point in Go terminology is a rune.

## Runes and bytes

Note that we have runes that take more than a single byte.

![](/core/src/07-bytes-runes-strings-utf8/assets/702-runes.png)

If you want to take a slice of runes you need a two step process to convert from slice of bytes to slice of runes.

![](/core/src/07-bytes-runes-strings-utf8/assets/703-rune-slice.png)

## strings package

covering some typical methods of the strings package for splitting, replacing and general tinkering with strings.

![](/core/src/07-bytes-runes-strings-utf8/assets/704-pkg.png)

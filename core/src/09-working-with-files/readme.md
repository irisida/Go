![](/assets/gologo.png)

# The Go programming Language - Working with files

the common way to work with files in Go is wth the `os` package. It provides a platform independent way to handle operating system functionality in relation to the file system. The `os` interface being unix like allows us a single way to approach file operation on any platform. There are additional packages `io`, `ioutil` and `bufio` but the basic operations can be fully achieved with just the `os` package.

## Simple file writing operations

Basic file operations using the `os` package and the `log` package for idiomatic error logging

![](/core/src/09-working-with-files/assets/901-file-basics.png)

The `reader` and `writer` interfaces are abstraction used to read and write bytes to files without being bogged down by how we read or where we write the data. Let's see this with the `os` package and the truncation and `write` modes.

![](/core/src/09-working-with-files/assets/902-write.png)

Here we can see an `ioutil` package version of the simple write operation. This poses the question that if the os packages does this already why are we duplicating it? The answer is eah import has a cost and increases the statically linked binary size in the end.

![](/core/src/09-working-with-files/assets/903-write.png)

We can also use the `bufio` package to achieve a buffered and optimised writing strategy using a 4096 default buffer.

![](/core/src/09-working-with-files/assets/904-bufio.png)

## Reading files operations

We can use both the `io` and `ioutil` packages for reading, see the common approach below.

![](/core/src/09-working-with-files/assets/905-buffered-reader.png)

We can use `bufio` to have a buffered reader and read a file line by line, of by split/configured separator. Note here we're showing that you can read by Bytes() and Text() too.

![](/core/src/09-working-with-files/assets/906-line-reader.png)

Next we're seeing how to read from the stdin until a quit condition is met.

![](/core/src/09-working-with-files/assets/907-stdin.png)

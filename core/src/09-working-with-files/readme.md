![](/assets/gologo.png)

# The Go programming Language - Working with files

the common way to work with files in Go is wth the `os` package. It provides a platform independent way to handle operating system functionality in relation to the file system. The `os` interface being unix like allows us a single way to approach file operation on any platform. There are additional packages `io`, `ioutil` and `bufio` but the basic operations can be fully achieved with just the `os` package.

## Simple file operations

Basic file operations using the `os` package and the `log` package for idiomatic error logging

![](/assets/core/09/09-901-file-basics.png)

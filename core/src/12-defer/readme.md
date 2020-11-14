![](/assets/gologo.png)

# The Go programming Language - The defer statement

The `defer` statement allows us to make a particular action the last one,or last thing before a function exits. The scope is for the function it is enclosed in.

![](/assets/core/12/1201-defer.png)

A second example shows us deferring the primary operation of a function, calling other function and only then completing it action before exiting.

![](/assets/core/12/1202-scope-defer.png)

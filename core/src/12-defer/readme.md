![](/assets/gologo.png)

# The Go programming Language - The defer statement

The `defer` statement allows us to make a particular action the last one,or last thing before a function exits. The scope is for the function it is enclosed in.

![](/core/src/12-defer/assets/1201-defer.png)

A second example shows us deferring the primary operation of a function, calling other function and only then completing it action before exiting.

![](/core/src/12-defer/assets/1202-scope-defer.png)

The primary example of defer is where we want to handle a file, we open the file, `f` and call the f.close() immediately thereafter with a `defer` and this maintains a high readability, we don't have to scan lot of program code to find a matching close of the file, it's contextually complete and yet the action is deferred until the last action as would normally be the programmatic flow, again introducing a cognitive easing and syntactic beauty to the process.

![](/assets/gologo.png)

# The Go programming language - Concurrency

Go is the first major programming language with concurrency built in. Intel announced the first multi-core CPU in 2006 and the Go team formed in 2007 and released first public version in 2009, reaching V1.0 in 2012.

Firstly, let's be clear that `concurrency is not parallelism`. Concurrency means dealing with multiple things at once, while parallelism is the simultaneous execution of processes and requires multiple core CPUs. Still confused?

- `concurrency` - means loading more `goroutines` at a time. If one goroutine blocks another is picked up and started. On a single core CPU you can run ONLY concurrent applications but that are **not** run in parallel.
- `parallelism` - multiple `goroutines` executed at the same time.

## Go routines

A Go routine is a lightweight thread of execution. The building blocks of Go's concurrency. So what is it? It's a function capable of running concurrently with other functions. To create a go routine we use the keyword `go` followed by the function invocation. In physical terms a go routine is smaller than a thread too.

- goroutines take around 2kb of stack space to initialise compared to a thread which takes around 2mb.
- OS thread stack is fixed size but goroutines shrink and grow as needed.
- scheduling a goroutine is cheaper than scheduling a thread in resources.
- OS threads are scheduled by the OS kernel, goroutines are scheduled by the go scheduler.
- goroutines have no identity.

![](/core/src/16-concurrency/assets/1601-goroutine.png)

To coordniate with groutines we can use waitgroups instead of explicit waits which would be a wasteful and cumbersome way to control how our routimes are handled. See here an example of how to declare a waitgroup and how to have the function wait for the waitgroup to be done.

![](/core/src/16-concurrency/assets/1602-waitgroups.png)

![](/assets/gologo.png)

# The Go programming language - Concurrency

Go is the first major programming language with concurrency built in. Intel announced the first multiccore CPU in 2006 and the Go team formed in 2007 and released first public version in 2009, reaching V1.0 in 2012.

Firstly, let's be clear that `concurrency is not parallelism`. Concurrency means dealing with multiple things at once, while parallelism is the simultaneous execution of processes and requires multiple core CPUs. Still confused?
- `concurrency` - means loading more `goroutines` at a time. If one goroutine blocks another is picked up and started. On a single core  CPU you can run ONLY concurrent applications but that are **not** run in parallel.
- `parallelism` - multiple `goroutines` executed at the same time.
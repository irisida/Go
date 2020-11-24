![](/assets/gologo.png)

# The Go programming language - Channels

A channel provides a connection between two go-routines allowing them to communicate. Channels take a type and the data communicated must always be of that type.
An initialised channel holds an address and is much like a pointer and can be passed like a pointer.

Let's see the channel operator. We cannot run this sample as we will have a deadlock because no goroutine was created, but we can see the syntactic introduction that is made with channels

![](/core/src/17-channels/assets/17-channel-syntax.png)

## Channels with go-routines

Here we see an example of a single pass/return communication as well as a multi pass/return in a loop.

![](/core/src/17-channels/assets/1702-channel-routine.png)

## Channels, go-routines and function literals (anoymous functions)

It is common practice in Go to use a function literal (anonymous function) to spawn a new go-routine. Here we see that the function literal is called with the `go` keyword to create the call as a go-routine. The function is then immediately invoked, failure to invoke will result in a compilation error.

![](/core/src/17-channels/assets/1703-function-literal.png)

## Updated webserver response checker app with channels and go-routines

Next we will update the previous section project, the URL service checker to refactor and use channels and go-routines. It is important to remember if checking the number of goroutines an application has that the main process will count as one too, so if we're checking 3 URls we expect 4 goroutines in total, 5 URls meas 6 goroutines in total and so on.

![](/core/src/17-channels/assets/1704-url_checker.png)

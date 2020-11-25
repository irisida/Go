![](/assets/gologo.png)

# The Go programming language - Channels

A channel provides a connection between two go-routines allowing them to communicate. Channels take a type and the data communicated must always be of that type.
An initialised channel holds an address and is much like a pointer and can be passed like a pointer.

Let's see the channel operator. We cannot run this sample as we will have a deadlock because no goroutine was created, but we can see the syntactic introduction that is made with channels

![](/core/src/17-channels/assets/17-channel-syntax.png)

## Channels with go-routines

Here we see an example of a single pass/return communication as well as a multi pass/return in a loop.

![](/core/src/17-channels/assets/1702-channel-routine.png)

## Channels, go-routines and function literals (anonymous functions)

It is common practice in Go to use a function literal (anonymous function) to spawn a new go-routine. Here we see that the function literal is called with the `go` keyword to create the call as a go-routine. The function is then immediately invoked, failure to invoke will result in a compilation error.

![](/core/src/17-channels/assets/1703-function-literal.png)

## Updated webserver response checker app with channels and go-routines

Next we will update the previous section project, the URL service checker to refactor and use channels and go-routines. It is important to remember if checking the number of goroutines an application has that the main process will count as one too, so if we're checking 3 URls we expect 4 goroutines in total, 5 URls meas 6 goroutines in total and so on.

![](/core/src/17-channels/assets/1704-url-checker.png)

## Running process continuously

Now we'll strip the file-writing out the project as the key here is to demonstrate the ability to cycle continuously across a channel by putting he ouput of a channel back into it. We have 3 posibe syntax options here to iterate over a range of urls continuously using the channels communication.

![](/core/src/17-channels/assets/1705-background-channels.png)

## Buffered and unbuffered channels

When channels are declared without a capacity they are unbuffered, or what is often called synhronised/synchronous channels. A buffered channel has a capcity and even a closed channe will hold that capacity. Proceeding with printing the output of the channel for a closed channel will yeild a zero value for thet type of channel it is. If we try to send/pass more data into it we will have a `panic`.


Interms of benefits the unbuffered option has stronger synchronisation. Buffered options are more decoupled. If the upper limit of operations is known n advance buffered is a good option because we can send all the values befre the first is even received. Where that is unknown and/or a coupling is required the unbuffered one is a better choice.

#### An unbuffered example
![](/core/src/17-channels/assets/1706-unbuffered.png)

#### A buffered example
![](/core/src/17-channels/assets/1707-buffered.png)

## The select statement

The select statement lets a go routine wait on multiple communications operations. It blocks until one of its cases can be run and then it executes that cases. In the event the unblocking means that multiple applicable cases can now be run then Go will handle the multiple cases by selecting one at random for the next processing.

This is a powerful feature of `Go` in that it allows us to wait on multiple channel operations. Send a recieve operations on a nil channel block forever ad this can be used to disable a channel in a select statement. the `select` is only used with, only applicable to, channels in Go.

![](/core/src/17-channels/assets/1708-select.png)

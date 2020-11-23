![](/assets/gologo.png)

# The Go programming language - Channels

A channel provides a connection between two go-routines allowing them to communicate. Channels take a type and the data communicated must always be of that type.
An initialised channel holds an address and is much like a pointer and can be passed like a pointer.

Let's see the channel operator. We cannot run this sample as we will have a deadlock because no goroutine was created, but we can see the syntactic introduction that is made with channels

![](/core/sr/17-channels/assets/17-channel-syntax.png)
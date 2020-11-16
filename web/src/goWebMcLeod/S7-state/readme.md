![](/assets/gologo.png)

# Web development with Go - Section 7 - State

What is state? A point in time reflection of values and conditions of a computer users account or transaction. To interact with state we must be able to pass data back and forth and that means we have two types of possibility:

1. through the `body` with `post`
2. through the `url` with `get`

## Passing data in the url with a get request

As detailed above, when using the url technique we are passing data by get request.

![](/web/src/goWebMcLeod/S7-state/assets/701-pass-in-url.png)

## Pass data in body with a post request

The browser default method is a `get` request so we have to specify post request where we would want to make a request with the `post` method. In doing so that shifts the data transferance to the body, or form.

![](/web/src/goWebMcLeod/S7-state/assets/701-pass-in-body.png)

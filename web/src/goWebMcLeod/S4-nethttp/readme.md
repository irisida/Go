![](/assets/gologo.png)

# Web development with Go - Section 4 - Understanding the net/http package

- [Simple HTTP](/web/src/goWebMcLeod/S4-netHttp/01-simpleHttp)
- [Exploring http.Request](/web/src/goWebMcLeod/S4-netHttp/02-request)

## Key takeaways

The entry point for understanding the http package is the `handler` interface. This interface implements the `ServeHTTP(ResponseWriter, *Request)` method. Where this is implemented in other packages it implicitly satisfies the handler interface requirements.

![](/assets/web/goWebMcLeod/S4/createHandler.png)

When we explore the `http.Request` aspect of the http package, we can see the available methods such as `ParseForm`. In our sample the handler is fired when the address is reached and a POST request with the form is parsed. This passes the data to the template and the template injects the data values passed for the value in an input field. Clicking the button will fiure the form.

![](/assets/web/goWebMcLeod/S4/04-402-request-demo.png)

![](/assets/web/goWebMcLeod/S4/04-403-request.png)

We can also see the response writer in action using trivial response object. Note we can add custom headers as the headers take a key-va;ue pairing.

![](/assets/web/goWebMcLeod/S4/04-404-response-writer.png)

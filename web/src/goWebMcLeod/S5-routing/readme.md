![](/assets/gologo.png)

# Web development with Go - Section 5 - Understanding Routing

Terms synonymous with routing are `mux` and `multiplexing`, `servemux` and `server`. In electronics a multiplexer, or mux, is a device that selects one of several input signals and forwards it to single line, very much alike the concept of a server listening to requests of a given route and calling the function that is intended for calls to that route.

## mux concept basics

let's see a trivial concept model in action.

![](/web/src/goWebMcLeod/S5-routing/assets/501-servemux-basics.png)

Above is an inelegant solution and it is typical for a restful server to handle routes like types and point to typehnadling code blocks to process the request. Typically in a modern server we would manage that by method type and route. In the meantime here is a simple handler implementation and types per route

![](/web/src/goWebMcLeod/S5-routing/assets/502-servemux-routes.png)

## The default ServeMux

The next stage would be to move from a declared ServeMux to the default one supplied by Go. Here we still use the handlers as before but we are using the defauly mux of the http package.

![](/web/src/goWebMcLeod/S5-routing/assets/503-default-mux.png)

## The HandlerFunc

HandleFunc registers the handling function for a given pattern/route in the default serveMux. `func HandleFunc(pattern string, func(ResponseWriter, *Request))`

Additional notw where we see `ListenAndServe(address, nil)` the nil denotes that we are using the default serveMux and when this is not nil there will need to be a programmer defined and initialised servemux

![](/web/src/goWebMcLeod/S5-routing/assets/504-handlefunc.png)

![](/assets/gologo.png)

# Web Development with Go

## Section 2 - Templates and rendering templates

[Here](/web/src/goWebMcLeod/S2-templates/readme.md) I'm covering:

- Basic templating and creating html templates.
- Interpolating data from our source code into a template generation.
- How templates work with Go, static complication and the one chance to submit data.
- How to deal with complex data structures as the data passed into a template
- functions, methods, composition as part of a template.

## Section 3 - Creating your own Servers

[Here](/web/src/goWebMcLeod/S3-servers) I'm covering:

- creating a basic TCP server
- How HTTP sits on top of TCP
- reading and writing with buffers

## Section 4 - Understanding the net/http package

[Here](/web/src/goWebMcLeod/S4-nethttp) I'm covering

- implementing the handler interface
- exploring http.Request
- exploring responses

## Section 5 - Understanding routing, mux, servemux, multiplexer

[Here](/web/src/goWebMcLeod/S5-routing) I'm covering

- ServeMux basics
- ServeMux handlers
- ServeMux - the default multiplexer
- ServeMux and the HandleFunc, handlerFunc types

## Section 6 - Serving files

[Here](/web/src/goWebMcLeod/S6-serving-files) I'm covering

- serving files with io.copy
- serving local files with io.copy
- serveContent and serveFile
- fileservers
- static file servers
- notFoundHandler and favicons

![](/assets/gologo.png)

# Web development with Go - Section 6 - Serving files

## Serving files with io.copy

We can use of server to read and write out with io.copy to relay an internet image.

![](/web/src/goWebMcLeod/S6-serving-files/assets/601-io-copy.png)

## Using io.copy to serve local files

We can also use the io.copy method to serve locally hosted files on our server.

![](/web/src/goWebMcLeod/S6-serving-files/assets/602-io-copy-local-file.png)

Net we have ServeFile and ServeContent both part of the `http` package.

![](/web/src/goWebMcLeod/S6-serving-files/assets/603-servefile.png)

We can also strip prefixes for a server structure setup. As demoed here...

![](/web/src/goWebMcLeod/S6-serving-files/assets/604-strip-prefix.png)

When using a static fileserver with `http.FileServer` we have a special case where in the presence of an index.html file being found it will prevent a directly listing presenting the source files as well as other file assets.

![](/web/src/goWebMcLeod/S6-serving-files/assets/605-default-index.png)

![](/web/src/goWebMcLeod/S6-serving-files/assets/605-markup.png)

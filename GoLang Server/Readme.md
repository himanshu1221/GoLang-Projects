# Golang Server

## What is this project about ?

- This is a simple GoLang project which consist of three routes `/`,`/form`,`hello` which uses `GET` & `POST` method along with basics to handle `http requests`

## Following are the topics I have learned along with resources

- Http -> Package http provides HTTP client and server implementations.
Get, Head, Post, and PostForm make HTTP (or HTTPS) requests

[source](https://pkg.go.dev/net/http)

- http.Dir -> When you call Open on an http.Dir it will open the file you named relative to the original path used to create the http.Dir

[source](https://forum.golangbridge.org/t/how-does-http-dir-work/9203/2)

- http.FileServer -> it serves up an entire directory path. and all its children.

(In this project static folder contains all the static files)

[source](https://stackoverflow.com/questions/28793619/golang-what-to-use-http-servefile-or-http-fileserver)

- http.Handle -> it is a function that is used to register a handler for a specific URL pattern with the HTTP server.

[source](https://stackoverflow.com/questions/21957455/difference-between-http-handle-and-http-handlefunc)

- http.HandleFunc ->  is a function that is used to register a handler function for a specific URL pattern with the HTTP server.

[source](https://stackoverflow.com/questions/21957455/difference-between-http-handle-and-http-handlefunc)

- Diffrence between HandleFunc & handle ?
```
http.Handle expects a value that implements the http.Handler interface, 
which means that it must have a ServeHTTP method with the signature ServeHTTP(http.ResponseWriter, *http.Request). 
This method is responsible for handling the incoming HTTP request.

On the other hand, http.HandleFunc expects a function with the signature func(http.ResponseWriter, *http.Request). 
This function is automatically converted into an http.Handler that can be registered with the server using http.Handle.
```
- w http.ResponseWriter -> This is an interface that provides methods for writing the response back to the client. You can use this to write the HTTP response headers and the body of the response.

- r *http.Request -> This is a pointer to a http.Request struct, which contains information about the incoming HTTP request, such as the request headers, query parameters, and body.

- r.ParseForm() -> it is a function that parses the form data in the body of an HTTP POST request or the query string of an HTTP GET request and stores the data in a map called r.Form. The map contains the values of the form fields, indexed by their names.

- URL.Path -> it is used to check the path of the URL

- r.Method -> It is used to check the method of the API

- http.Error -> This function is used to throw the error

- r.FormValue function that returns the first value associated with the given form key.

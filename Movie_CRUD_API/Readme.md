# Movie CRUD API

## What is this project about ?
In this project we will be working with the CRUD operations (Create Read Update Delete) where we can get all the movies,get them by an unique id,create a new movie,update new movie & Delete a movie 
```We are not using any Database we are using structs and slices to perform operations```

## Dependencies
- `encoding/json` : This is used to eencode data into JSON when we send it to any API testing 
- `fmt` : This is used to print out stuff
- `log` : This is used to log out the error
- `math/rand` : It used to generate random number to give id to the movies 
- `net/http` : It is used to create a server
- `strconv` : It is used to convert numbers to strings
- `github.com/gorilla/mux` : External library provides functionalities for matching routes, serving static files, serving single-page applications, middleware, handling CORS requests, and testing handlers

## Topics I have Learnt

- `Struct` : It is like an Object in Javascript/Java 
- `Pointer` : pointer is a variable that stores the address of an object stored in memory
- `Slice` : Slices are similar to arrays, but are more powerful and flexible
- `mux.NewRouter()` : It allows you to create a router that will apply a set of matches to all routes you register with it
- `http.ListenAndServe` : Is is used to start the server at a given port using all the routes that have been declared
- `Log.Fatal` : It is used to Log out the error if any
- `w.Header().Set()` : It's a method used to set HTTP headers in an HTTP response
- `mux.Vars()` : It extracts the variables from the current request URL
- `json.NewEncoder().Encode()` : It is a function that encodes the specified value to JSON format and writes it to the specified io.Writer
- `json.NewDecoder().Decode()` : It is a common way to decode JSON data from an HTTP request body.
- `...` : is the ellipsis operator, which unpacks the elements of a slice as individual arguments.
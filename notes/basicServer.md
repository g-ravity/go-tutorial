You can directly serve static files in Go using the http.FileServer function

```go
http.Handle("/", http.FileServer(http.Dir("./static")))
```

This way we can directly serve static files like index.html or other files.

```go
http.ListenAndServe(":8080", nil)
```

This is used to actually create and listen to the server. ListenAndServe returns an error, which can be handled in a single line like this:

```go
if err := http.ListenAndServe(":8080", nil); err != nil {
	log.Fatal(err)
}
```

What log.Fatal does is that, it logs the error and then exits the program using os.Exit(1)

In the ListenAndServe function, the second argument is the handler. If we pass nil, it will use the default handler, which is the DefaultServeMux. Otherwise, we can pass our own handler which we can create using newServeMux. Mux is short for multiplexer, which can take multiple inputs and route them to different handlers.

The main difference between the DefaultServeMux and the newServeMux is that the DefaultServeMux is the global handler, which is used by all the requests. The newServeMux is a local handler, which is used by only the requests that are handled by it. In general, it is a good practice to use the newServeMux as it is more secure and easy to test and we can use multiple routers in the same application.

This is how we can use newServeMux

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
})
```

The handlers have two arguments, the standard response writer and the request. The response writer is used to write the response to the client. The request is used to get the request from the client. It is generally received as a pointer.

Now, if we are handling the router logic ourselves, then we can add checks for URL path and HTTP method as done in the project, otherwise these things will be handled itself if we use some framework like gin/gorilla mux.

Fprintf is different from simple printf. Fprintf writes to the writer interface provided (in this case it is the response writer, but any output stream which implements the writer interface can be used here), while printf writes to the standard output (console).

To receive form values from POST request

```go
if err := r.ParseForm(); err != nil {
    fmt.Fprintf(w, "ParseForm() err: %v", err)
}

name := r.FormValue("name")
```

Standard error handling. We can use r.FormValue("fieldName") to get the value of the field in the variable

Note this type:

```go
map[string]interface{}
```

This is a very interesting type because it is a map of string to interface. This means that the value can be any type. This is useful because we can use it to store any type of data in a map.

Generally, with maps, we can only store values of the same type. But with this type, we can store values of different types.

This is useful if we need to store, process and return json values in Go.

To return json values, we can use the json package, and specifically json.Marshal function

```go
userData := map[string]interface{}{
    "name":   firstName + " " + lastName,
    "mobile": mobile,
}

w.Header().Set("Content-Type", "application/json")

json, err := json.Marshal(userData)
if err != nil {
    http.Error(w, "Error in JSON response", http.StatusInternalServerError)
    return
}

w.Write(json)
```

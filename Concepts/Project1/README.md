# GoLang's built in HTTP Serve vs Gorilla Mux!

1. Go's `net/http` package has already provided a decent number of API's for writing
    HTTP applications. However, there is one thing that it does not do very well, which is complex
    request routing like segmenting a request URL into single parameters. 
   
2. Fortunately, there is a very popular community developed package that handles 
    that responsibilities named gorilla/mux.
   
3. With gorilla/mux, we can declare complex routes with variables, constrain routes 
    with methods etc.
   
4. We declared routes with variables, we should be able to capture the valuesof those
    variables in Handlers.
   
5. The net/http implementation is already powerful, By using gorilla/mux, we are able to get 
    an even simpler interface for dealing with complex routing and managing requests.
   
# Firebase

1. It is a NoSQL Database by Google
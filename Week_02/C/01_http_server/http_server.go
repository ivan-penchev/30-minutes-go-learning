package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/ivan-penchev/30-minutes-go-learning/Week_02/C/01_http_server/myOtel"
)

// Routes are a collection of defined api endpoints
// []route is an array of the route struct
type routes []route

// A Route defines the parameters for an api endpoint
type route struct {
	Name        string
	Auth        bool
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// global constant
const APPNAME = "http_server"

// Build the routing table and start the HTTP server
func main() {

	var thisRouter *mux.Router
	var err error

	// port for the http server
	port := "8080"

	// #############################################################
	// NEW CODE FOR OPENTELEMETRY:
	ctx := context.TODO()
	var shutdownTracer myOtel.ShutdownTracerFn
	ctx, shutdownTracer, err = myOtel.ConfigureStdoutTracing(ctx, APPNAME)
	if err != nil {
		defer shutdownTracer()
	}
	// #############################################################

	//
	routes := routes{}

	// add multiple routes as required:
	// append is a built in function and adds a new "route" to the "routes" array
	routes = append(routes, route{Name: "NewCustomer", Auth: false, Method: strings.ToUpper("GET"), Pattern: "/hello" + "/{ID}", HandlerFunc: hello})

	// build the routing table from the routes
	thisRouter, err = newRouter(routes)

	if err == nil {
		fmt.Printf("listening on port=%s\n", port)
		log.Fatal(http.ListenAndServe(":"+port, thisRouter))
	} else {
		fmt.Println(err.Error())
	}

}

// newRouter creates a new router
func newRouter(routes routes) (router *mux.Router, err error) {

	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		// NOTE:
		// middleware handlers intercept the request to do special processing
		// you can chain multiple middleware handlers
		// they are executed in reverse order of addition

		// call the assigned function
		handler = route.HandlerFunc

		if route.Auth {
			// TODO: put your auth middleware here
		}

		// TODO: add a cors handler
		// handler = CorsWrapper.Handler(handler)

		// add a logging handler
		handler = HttpLogger(handler, route.Name)

		// #############################################################
		// NEW CODE FOR OPENTELEMETRY:
		handler = myOtel.NewTracedHttpHandler(context.TODO(), "http_server", handler)
		// #############################################################

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		fmt.Printf("Route=%s Path=%s Method=%s\n", route.Name, route.Pattern, route.Method)
	}
	return router, nil
}

// handlers return functions with a defined signature
//
//	type Handler interface {
//	    ServeHTTP(ResponseWriter, *Request)
//	}
func HttpLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// start timer
		fmt.Printf("HTTP REQ %s %s Verb:%s Timer=start\n", r.Method, r.URL.Path, name)

		inner.ServeHTTP(w, r)

		// Log request
		fmt.Printf("HTTP %s %s Verb:%s Timer=%s\n", r.Method, r.URL.Path, name, time.Since(start))
	})
}

// The helloValue struct is used by the hello function
type helloValue struct {
	// must be uppercase for json to work
	Hello string `json:"hello"`
}

// This is the function we call when the GET /hello/<string> request is received
func hello(w http.ResponseWriter, r *http.Request) {

	var err error
	ctx := context.TODO()

	// #############################################################
	// NEW CODE FOR OPENTELEMETRY:
	// start the root trace span
	var span myOtel.Span
	ctx, span = myOtel.TraceStart(ctx, APPNAME, "hello")
	defer myOtel.TraceEnd(span)
	// #############################################################

	// get the parameters
	params := mux.Vars(r)
	hv := helloValue{Hello: params["ID"]}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	err = enc.Encode(hv)
	if err != nil {

		// #############################################################
		// NEW CODE FOR OPENTELEMETRY:
		// write the open telemetry trace
		myOtel.TraceError(span, err)
		// #############################################################
		fmt.Println(err.Error())

	}
}

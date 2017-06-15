/*
HTTP routes definitions and registrations.

Each route can be defined by instantiating a route type and populating its properties,
routes should be appended under 'routes' variable of type 'urlRoutes'.
 */

package api

import (
	"net/http"
)

const basePath = "/v1/" // todo move to config file
var schemes = []string {"HTTP", "HTTPS"} // todo move to config file
var dummyHandler = func(http.ResponseWriter, *http.Request){}  // todo to be deleted

// HTTP methods
const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DELETE = "DELETE"
)

// httpMethodsHandler represents the mapping between routes and their handlers function.
type httpMethodsHandler map[string]func(http.ResponseWriter, *http.Request)
// A route contains each HTTP route properties
type route struct{
	name    string
	methods httpMethodsHandler
	path    string
}
// Container type of route types
type urlRoutes []route

// routes definitions go here
var routes urlRoutes = urlRoutes{
	route{
		name: "info",
		path: "/info",
		methods: httpMethodsHandler{
			GET: dummyHandler,
		},
	},
	route{
		name: "words",
		path: "/words",
		methods: httpMethodsHandler{
			GET: dummyHandler,
			POST: dummyHandler,
		},
	},
	route{
		name: "word",
		path: "/words/{id}",
		methods: httpMethodsHandler{
			GET: dummyHandler,
			PUT: dummyHandler,
			DELETE: dummyHandler,
		},
	},
}

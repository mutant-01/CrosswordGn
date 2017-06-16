/*
HTTP routes definitions and registrations.

Each route can be defined by instantiating a route type and populating its properties,
routes should be appended under 'routes' variable of type 'urlRoutes'.
 */

package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

const basePath = "/v1/" // todo move to config file
var schemes = []string {"HTTP", "HTTPS"} // todo move to config file
var dummyHandler = func(http.ResponseWriter, *http.Request){}  // todo to be deleted

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

// generateRouter generates the API router and maps handlers accordingly based on 'routes' container.
func generateRouter() *mux.Router{
	baseRouter := mux.NewRouter().Schemes(schemes...).PathPrefix(basePath).Subrouter()

	for _, r := range routes{
		for method, handler := range r.methods{
			baseRouter.HandleFunc(r.path, handler).Methods(method)
		}
	}

	return baseRouter
}

// routes definitions go here
var routes urlRoutes = urlRoutes{
	route{
		name: "info",
		path: "/info",
		methods: httpMethodsHandler{
			http.MethodGet: dummyHandler,
		},
	},
	route{
		name: "words",
		path: "/words",
		methods: httpMethodsHandler{
			http.MethodGet: dummyHandler,
			http.MethodPost: dummyHandler,
		},
	},
	route{
		name: "word",
		path: "/words/{id}",
		methods: httpMethodsHandler{
			http.MethodGet: dummyHandler,
			http.MethodPut: dummyHandler,
			http.MethodDelete: dummyHandler,
		},
	},
}

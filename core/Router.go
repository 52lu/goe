package core

import "net/http"

type RouteConfig struct {
}

func (receiver RouteConfig) RegisteredRoute()  {
}

func (receiver *RouteConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return 
}

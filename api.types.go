type middleware func(http.HandlerFunc) http.HandlerFunc

type route struct {
  middlewares []middleware
  handler http.HandlerFunc
}

type routeCollection struct {
  middlewares []middleware
  routes map[string]route
}
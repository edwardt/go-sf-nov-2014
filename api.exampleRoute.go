func init() {
  var r routeCollection
  register("/v1/users", &r)
  
  r.add("GET /", route{
    middlewares: []middleware{
      onlyAdmin,
    },
    
    handler: function(w http.ResponseWriter, r *http.Request) {
      // fetch all users
    },
  })
}
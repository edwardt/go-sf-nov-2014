func onlyAdmin(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // send an error if the user who sent the request isn't an admin
    
    next(w, r)
  }
}
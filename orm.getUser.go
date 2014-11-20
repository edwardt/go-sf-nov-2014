var user models.User
found, err := db.GetByID(&user, 4) // try to find user 4
if err != nil {
  fmt.Println("failed to get user:", err)
} else if !found
  fmt.Println("no such user")
}
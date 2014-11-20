package main

func main() {
	newUser := models.User{
		Name:     "Benjamin Franklin",
		Email:    "lightningman1709@foundingfathers.com",
		Password: "italianwoah",
	}
	if err := db.Create(&newUser); err != nil {
		fmt.Println("failed to create new user:", err)
	}
}

package models

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ID        int64     `json:"ID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Virtual field that stores plaintext password in memory
	Password string `json:"password,omitempty"`

	// Hashed password persisted to the database
	PasswordHash string `json:"-"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) Columns() map[string]interface{} {
	return map[string]interface{}{
		"name":          &u.Name,
		"email":         &u.Email,
		"password_hash": &u.PasswordHash,
		"id":            &u.ID,
		"created_at":    &u.CreatedAt,
		"updated_at":    &u.UpdatedAt,
	}
}

func (u *User) InsertColumns() []string {
	return []string{
		"name", "email", "password_hash",
	}
}

func (u *User) UpdateColumns() []string {
	return []string{
		"name", "email", "password_hash",
	}
}

func (u *User) PreInsert() error {
	u.PasswordHash = hash(u.Password)
}

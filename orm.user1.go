package models

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ID        int64     `json:"ID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) Columns() map[string]interface{} {
	return map[string]interface{}{
		"name":       &u.Name,
		"email":      &u.Email,
		"id":         &u.ID,
		"created_at": &u.CreatedAt,
		"updated_at": &u.UpdatedAt,
	}
}

func (u *User) InsertColumns() []string {
	return []string{
		"name", "email",
	}
}

func (u *User) UpdateColumns() []string {
	return []string{
		"name", "email",
	}
}

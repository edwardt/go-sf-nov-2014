// START1 OMIT
package models

type Comment struct {
	ID        int64     `json:"ID"`
	UserID    int64     `json:"userID"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *Comment) Table() string {
	return "comments"
}

func (c *Comment) Columns() map[string]interface{} {
	return map[string]interface{}{
		"id":         &c.ID,
		"user_id":    &c.UserID,
		"contents":   &c.Contents,
		"created_at": &c.CreatedAt,
		"updated_at": &c.UpdatedAt,
	}
}
// STOP1 OMIT

// START2 OMIT
func (c *Comment) InsertColumns() []string {
	return []string{
		"user_id",
		"contents",
	}
}

func (c *Comment) UpdateColumns() []string {
	return []string{
		"contents",
	}
}

// Then, let's mark Comment as a Dependent of User

func (u *User) Dependents() map[db.Model]db.Join {
	return map[db.Model]db.Join{
		new(Comment): {
			SelfColumn: "id",
			JoinColumn: "user_id",
		},
	}
}
// STOP2 OMIT
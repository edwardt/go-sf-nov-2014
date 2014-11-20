var models []db.Model
err := db.GetAll(
  &models, new(models.Comment),
  db.WhereMap{ "user_id": 4 },
)
var comments = make([]models.Comment, len(models))
for i, comment := range comments {
  comments[i] := comment.(*models.Comment) 
}
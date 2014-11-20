package db

type Model interface {
	Table() string
	Columns() map[string]interface{}
	InsertColumns() []string
	UpdateColumns() []string
}

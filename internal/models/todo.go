package models

type Todo struct {
	ID     int
	Title  string
	Body   string
	Labels []string
}

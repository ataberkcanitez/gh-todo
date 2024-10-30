package services

import (
	"github.com/ataberkcanitez/gh-todo/internal/api"
)

type TodoService struct {
	client *api.GithubClient
}

func NewTodoService(client *api.GithubClient) *TodoService {
	return &TodoService{client: client}
}

func (s *TodoService) AddTodo(title, body string) error {
	return s.client.CreateIssue(title, body, "todo")

}

func (s *TodoService) ListTodos() (string, error) {
	return s.client.ListIssuesWithLabel("todo")
}

func (s *TodoService) DeleteTodoById(issueId string) error {
	return s.client.CloseIssue(issueId)
}

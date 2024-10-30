package cmd

import (
	"fmt"
	"github.com/ataberkcanitez/gh-todo/internal/services"
	"github.com/spf13/cobra"
)

func NewListCommand(service *services.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all todos",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := service.ListTodos()
			if err != nil {
				fmt.Println("error: ", err)
			} else {
				fmt.Println("todos:\n", todos)
			}
		},
	}
}

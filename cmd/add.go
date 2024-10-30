package cmd

import (
	"fmt"
	"github.com/ataberkcanitez/gh-todo/internal/services"
	"github.com/spf13/cobra"
)

func NewAddCommand(service *services.TodoService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new todo",
		Run: func(cmd *cobra.Command, args []string) {
			title, _ := cmd.Flags().GetString("title")
			body, _ := cmd.Flags().GetString("body")

			if title == "" || body == "" {
				fmt.Println("Please provide a title and a body")
				return
			}
			if err := service.AddTodo(title, body); err != nil {
				fmt.Println("error: ", err)
			} else {
				fmt.Println("todo added successfully")
			}
		},
	}

	cmd.Flags().String("title", "", "Title of the todo")
	cmd.Flags().String("body", "", "Body of the todo")
	return cmd
}

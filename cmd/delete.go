package cmd

import (
	"fmt"
	"github.com/ataberkcanitez/gh-todo/internal/services"
	"github.com/spf13/cobra"
)

func NewDeleteCommand(service *services.TodoService) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a todo by ID",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("please provide an ID")
				return
			}
			id := args[0]
			if err := service.DeleteTodoById(id); err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("todo deleted succesfully")
			}
		},
	}
}

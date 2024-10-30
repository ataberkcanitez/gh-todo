package main

import (
	"github.com/ataberkcanitez/gh-todo/cmd"
	"github.com/ataberkcanitez/gh-todo/internal/api"
	"github.com/ataberkcanitez/gh-todo/internal/services"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "todo"}
	githubClient := api.NewGithubClient()
	todoService := services.NewTodoService(githubClient)

	rootCmd.AddCommand(cmd.NewAddCommand(todoService))
	rootCmd.AddCommand(cmd.NewListCommand(todoService))
	rootCmd.AddCommand(cmd.NewDeleteCommand(todoService))

	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
	}
}

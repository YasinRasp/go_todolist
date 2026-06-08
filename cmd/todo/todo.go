package todo

import (
	"todolist/internal/commands"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "A CLI tool for managing tasks and todo lists",
	Long: `task is a CLI task manager that helps you organize your work.

Store tasks locally, mark them complete, and keep
track of your productivity over time.`,
Args: cobra.ArbitraryArgs,
}

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "adds a task to the list",
	Run:   commands.AddTask,
	Args:  cobra.MinimumNArgs(1),
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "adds a task to our list",
	Run:   commands.ListTask,
}

var compCmd = &cobra.Command{
	Use:   "complete [task id]",
	Short: "adds a task to our list",
	Run:   commands.CompTask,
	Args:  cobra.ExactArgs(1),
}

var delCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "adds a task to our list",
	Run:   commands.DelTask,
	Args:  cobra.ExactArgs(1),
}

var (
	all bool
)

func init() {
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(compCmd)
	RootCmd.AddCommand(delCmd)
	RootCmd.RemoveCommand()
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "showing all tasks, completed and uncomplete ones.")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

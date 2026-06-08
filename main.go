package main
import (
	"fmt"
	"os"
	"todolist/cmd/todo"
)


func main() {
	if err := todo.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

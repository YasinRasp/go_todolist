package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"todolist/internal/tools"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func AddTask(cmd *cobra.Command, args []string) {
	var text string

	if len(args) > 0 {
		text = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
			if scanner.Err() != nil {
				panic("internal error")
			}
		}
		text = strings.Join(lines, "\n")
	}
	if text == "" {
		panic("no task input specified. enter a valid input")
	}

	data, lid := tools.CheckCsv()

	var task tools.Task

	task.ID = lid + 1
	task.Description = text
	task.DateCreated = time.Now()
	task.Done = false

	var ntask = []string{strconv.Itoa(task.ID), task.Description, task.DateCreated.Format(time.RFC3339), strconv.FormatBool(task.Done)}

	data = append(data, ntask)

	tools.SaveCsv(data)
}
func ListTask(cmd *cobra.Command, args []string) {

	data, _ := tools.CheckCsv()
	var tasks []tools.Task = tools.ConvertStT(data)
	if len(data) == 0 {
		fmt.Println("No tasks in your schedule, add one with the add command.")
		return
	}
	var w = tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)
	defer w.Flush()
	a := cmd.Flag("all").Value.String()
	fmt.Fprintln(w, "ID\tDescription\tDate Created\tStatus")
	all, _ := strconv.ParseBool(a)
	var done string
	for _, row := range tasks {
		if row.Done {
			done = "✔️"
		} else {
			done = "❌"
		}
		if all {

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", row.ID, row.Description, timediff.TimeDiff(row.DateCreated), done)
		} else if !row.Done {

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", row.ID, row.Description, timediff.TimeDiff(row.DateCreated), done)
		}
	}
}

func CompTask(cmd *cobra.Command, args []string) {

	var arg string

	if len(args) > 0 {
		arg = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
			if scanner.Err() != nil {
				panic("internal error")
			}
		}
		arg = strings.Join(lines, "\n")
	}
	var cid int
	var err error
	if cid, err = strconv.Atoi(arg); err != nil {
		panic("no task id specified. enter a valid id")
	}

	data, lid := tools.CheckCsv()
	if cid > lid {
		fmt.Println("specified id is out of bound, the last id is : ", lid)
		os.Exit(1)
	}
	var tasks []tools.Task = tools.ConvertStT(data)
	for irow, task := range tasks {
		if task.ID == cid {
			if task.Done {
				fmt.Println("specified task is already done! with Description : ", task.Description)
			} else {
				task.Done = true
				tasks[irow] = task
			}
		}
	}
	data = tools.ConvertTtS(tasks)
	tools.SaveCsv(data)
}
func DelTask(cmd *cobra.Command, args []string) {

	var arg string

	if len(args) > 0 {
		arg = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
			if scanner.Err() != nil {
				panic("internal error")
			}
		}
		arg = strings.Join(lines, "\n")
	}
	var cid int
	var err error
	if cid, err = strconv.Atoi(arg); err != nil {
		panic("no task id specified. enter a valid id")
	}

	data, lid := tools.CheckCsv()
	if cid > lid {
		fmt.Println("specified id is out of bound, the last id is : ", lid)
		os.Exit(1)
	}
	var tasks []tools.Task = tools.ConvertStT(data)
	for irow, task := range tasks {
		if task.ID == cid {
			tasks = append(tasks[:irow], tasks[irow+1:]...)
		}
	}
	data = tools.ConvertTtS(tasks)
	tools.SaveCsv(data)
}

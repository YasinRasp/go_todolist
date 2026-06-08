package tools

import (
	"encoding/csv"
	"errors"

	//"fmt"
	"os"
	"strconv"
	"time"
	//"github.com/spf13/cobra"
)

type Task struct {
	ID          int       `csv:"id"`
	Description string    `csv:"description"`
	DateCreated time.Time `csv:"date_created"`
	Done        bool      `csv:"done"`
}

func CheckFile(name string) bool {
	_, err := os.Stat(name)
	return !errors.Is(err, os.ErrNotExist)
}

func CheckCsv() ([][]string, int) {
	var data [][]string
	var lastID int
	if !CheckFile("file.csv") {
		return data, 0
	}
	file, err := os.Open("file.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var rd = csv.NewReader(file)
	data, err = rd.ReadAll()
	if err != nil {
		panic("csv read error.")
	}

	lastID, _ = strconv.Atoi(data[len(data)-1][0])

	return data, lastID
}

func SaveCsv(data [][]string) {

	var header []string
	var headis bool = true
	if !CheckFile("file.csv") {
		headis = false
		header = []string{"id", "description", "date_created", "done"}
	}

	file, err := os.Create("file.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	if !headis {
		if err := writer.Write(header); err != nil {
			panic(err)
		}
	}
	for _, record := range data {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}
}

func ConvertStT(data [][]string) []Task {
	var tasks []Task
	var task Task
	for i, record := range data {
		if i == 0 {
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic("string conversion error. check the .csv file.")
		}
		task.ID = id
		task.Description = record[1]
		tm, _ := time.Parse(time.RFC3339, record[2])
		task.DateCreated = tm
		done, err := strconv.ParseBool(record[3])
		if err != nil {
			panic("string conversion error. check the .csv file.")
		}
		task.Done = done
		tasks = append(tasks, task)
	}
	return tasks
}

func ConvertTtS(tasks []Task) [][]string {
	var data = [][]string{{"id", "description", "date_created", "done"}}
	var d []string
	for _, record := range tasks {
		id := strconv.Itoa(record.ID)
		description := record.Description
		date := record.DateCreated.Format(time.RFC3339)
		done := strconv.FormatBool(record.Done)

		d = []string{id, description, date, done}
		data = append(data, d)
	}
	return data
}

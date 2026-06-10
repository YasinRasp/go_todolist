package tools

import (
	"encoding/csv"
	"errors"

	"fmt"
	"os"
	"strconv"
	"time"
	"syscall"
	//"github.com/spf13/cobra"
)

type Task struct {
	ID          int
	Description string
	DateCreated time.Time
	Done        bool
}

func CheckFile(name string) bool {
	_, err := os.Stat(name)
	return !errors.Is(err, os.ErrNotExist)
}

func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func CheckCsv() ([][]string, int) {
	var data [][]string
	var lastID int
	if !CheckFile("file.csv") {
		return data, 0
	}
	file, err := loadFile("file.csv")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)
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

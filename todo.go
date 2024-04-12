package todo_cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

// item represents a task item in the todo list.
type item struct {
	Task        string    // Task represents the description of the task.
	Done        bool      // Done indicates whether the task is completed or not.
	CreatedAt   time.Time // CreatedAt represents the timestamp when the task was created.
	CompletedAt time.Time // CompletedAt represents the timestamp when the task was completed.
}

// List represents a collection of items.
type List []item

// AddTask adds a new task to the list.
// It takes a string parameter `task` which represents the task to be added.
// The task is appended to the list with the current timestamp as the creation time.
func (l *List) AddTask(task string) {
	item := item{Task: task, CreatedAt: time.Now()}
	*l = append(*l, item)
}

// CompleteTask marks a task as completed by setting the 'Done' field to true
// and updating the 'CompletedAt' field with the current time.
// It takes an index parameter representing the position of the task in the list.
// Returns an error if the index is invalid.
func (l *List) CompleteTask(index int) error {
	ls := *l
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].Done = true
	ls[index-1].CompletedAt = time.Now()

	return nil
}

// DeleteTask deletes a task from the list at the specified index.
// It returns an error if the index is invalid.
func (l *List) DeleteTask(index int) error {
	ls := *l
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*l = append(ls[:index-1], ls[index:]...)
	return nil
}

// Load reads the contents of a file and populates the List with the data.
// If the file does not exist or is empty, it returns nil.
// If there is an error reading or unmarshaling the file, it returns the error.
func (l *List) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, &l)
	if err != nil {
		return err
	}

	return nil
}

// Save saves the current state of the List to a file in JSON format.
// It takes a filename as a parameter and returns an error if any occurred.
func (l *List) Save(filename string) error {
	file, err := json.MarshalIndent(&l, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ListTasks returns a copy of the current list of tasks.
func (l *List) ListTasks() List {
	return *l
}

// ListCompletedTasks returns a new List containing only the completed tasks from the original List.
func (l *List) ListCompletedTasks() List {
	var completed List
	for _, item := range *l {
		if item.Done {
			completed = append(completed, item)
		}
	}
	return completed
}

// ListPendingTasks returns a new List containing all the pending tasks from the current List.
func (l *List) ListPendingTasks() List {
	var pending List
	for _, item := range *l {
		if !item.Done {
			pending = append(pending, item)
		}
	}
	return pending
}

// CountPendingTasks returns the number of pending tasks in the list.
func (l *List) CountPendingTasks() int {
	var count int
	for _, item := range *l {
		if !item.Done {
			count++
		}
	}
	return count
}

// PrintTable prints the list of tasks in a table format.
// It creates a table using the simpletable package and populates it with the task details.
// The table includes columns for task number, task name, task status, creation timestamp, and completion timestamp.
// Completed tasks are marked with a checkmark symbol (✓) and are displayed in green.
// Pending tasks are marked with a cross symbol (✗) and are displayed in red.
// The table also includes a footer that shows the count of pending tasks.
func (l *List) PrintTable() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: yellow("#")},
			{Align: simpletable.AlignCenter, Text: cyan("Task")},
			{Align: simpletable.AlignCenter, Text: cyan("Status")},
			{Align: simpletable.AlignCenter, Text: cyan("Created At")},
			{Align: simpletable.AlignCenter, Text: cyan("Completed At")},
		},
	}

	var cells [][]*simpletable.Cell

	for i, item := range *l {
		i++
		task := blue(item.Task)
		if item.Done {
			task = green(fmt.Sprintf("\u2713 %s", item.Task))
		}
		cells = append(cells, []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: yellow(strconv.Itoa(i + 1))},
			{Text: task},
			{Text: func() string {
				if item.Done {
					return green("Completed")
				}
				return red("Pending")
			}()},
			{Text: magenta(item.CreatedAt.Format("2006-01-02 15:04:05"))},
			{Text: magenta(item.CompletedAt.Format("2006-01-02 15:04:05"))},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignRight, Span: 5, Text: red("You have " + strconv.Itoa(l.CountPendingTasks()) + " pending tasks")},
	}}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

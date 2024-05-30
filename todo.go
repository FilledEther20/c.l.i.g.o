package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string    //task description
	Status      bool      //status of task ie done(True) or not(False)
	CreatedAt   time.Time //time of creation
	CompletedAt time.Time //time of completion
}

//DS to store the tasks would be a slice

type Todos []item

// Add method in
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Status:      false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

// Completed
func (t *Todos) Complete(index int) error {
	ls := *t // we would have to check whether the requested index exists or not
	if index <= 0 || index > len(ls) {
		return errors.New("index out of range")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Status = true
	return nil
}

// Deletion
func (t *Todos) Deletion(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("index out of range")
	}
	*t = append((*t)[:index-1], (*t)[index:]...)
	return nil
}

// Load Function
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

// Store function
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal((t))
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Print functionality
func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		Status := blue("no")
		if item.Status {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			Status = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: Status},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Status {
			total++
		}
	}

	return total
}

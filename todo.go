package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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
	if index <= 0 || index >= len(ls) {
		return errors.New("index out of range")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Status = true
	return nil
}

// Deletion
func (t *Todos) Deletion(index int) error {
	if index <= 0 || index >= len(*t) {
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

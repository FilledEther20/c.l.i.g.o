package main

import (
	"errors"
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

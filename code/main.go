package main

import "time"

type item struct {
	Task        string    //task description
	Status      bool      //status of task ie done(True) or not(False)
	CreatedAt   time.Time //time of creation
	CompletedAt time.Time //time of completion
}

//DS to store the tasks would be a slice

type Todos []item

//Methods (Functionalities offered)

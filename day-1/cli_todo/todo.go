package main

import (
	"errors"
	"os"
	"strconv"
	"time"
	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CompletedAt *time.Time
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	newTodo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, newTodo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of bounds")
		println(err)
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	err := t.ValidateIndex(index)
	if err != nil {
		return err
	} else {
		*todos = append(t[:index], t[index+1:]...)
	}
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	err := t.ValidateIndex(index)
	if err != nil {
		return err
	}
	if t[index].Completed {
		return errors.New("already completed")
	}
	completionTime := time.Now()
	(*todos)[index].CompletedAt = &completionTime
	(*todos)[index].Completed = true
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	err := t.ValidateIndex(index)
	if err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Index", "Title", "Completed", "Completed At", "Created At")
	for index, t := range *todos{
		Completed := "❌"
		CompletedAt := ""
		if t.Completed {
			Completed = "✅"
			if t.CompletedAt != nil {
                CompletedAt = t.CompletedAt.Format(time.RFC1123)
            }
		}
		table.AddRow(strconv.Itoa(index), t.Title, Completed, t.CreatedAt.Format(time.RFC1123), CompletedAt)		
	}	
	table.Render()
}
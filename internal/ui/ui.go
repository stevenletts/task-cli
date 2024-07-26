package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TodoItem struct {
	title       string
	description string
	created     int
	due         int
}

type Model struct {
	ToDos     []TodoItem
	choices   []string
	cursor    int
	ViewState int
}

const (
	ViewSelection int = iota
	ViewAdd
	ViewDo
	ViewList
)

// system level config options should be entry view ie on run come to add or list
// configure options for displaing in the todo item as in order of deadline, created, description, title or reverse

func (m Model) Init() tea.Cmd {
	// create the database and the table
	// connect app to database
	// fetch titles and deadline and display in view list
	// on enter into todo item we can display text created
	//
	//
	// docker run --name todo-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d --rm postgres:13.0
	// below will persist
	// docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d -v pgdata:/var/lib/postgresql/data postgres:13.0
	//
	// docs for the driver i have installed already
	// https://github.com/jackc/pgx

	return nil
}

func InitialModel() Model {
	return Model{
		choices: []string{"add", "do", "list"},
	}
}

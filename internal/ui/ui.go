package ui

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackc/pgx/v5"
)

type TodoItem struct {
	id          int
	title       string
	description string
	created     time.Time
	due         string
}

type Model struct {
	ToDos       []TodoItem
	choices     []string
	cursor      int
	ViewState   int
	CurrentTodo TodoItem
	conn        *pgx.Conn
	AddModel    AddTodoModel
}

type AddTodoModel struct {
	Title       string
	Description string
	DueDate     string
	cursor      int
	conn        *pgx.Conn
}

const (
	ViewSelection int = iota
	ViewAdd
	ViewList
	ViewTodo
)

// system level config options should be entry view ie on run come to add or list
// configure options for displaing in the todo item as in order of deadline, created, description, title or reverse

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Close() {
	if m.conn != nil {
		m.conn.Close(context.Background())
	}
}

func startPostgresContainer() error {
	cmd := exec.Command("docker-compose", "-f", "startup-utils/docker-compose.yml", "up", "-d", "postgres")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error starting PostgreSQL container: %v, output: %s", err, output)
	}
	return nil
}

func InitialModel() Model {
	connectionString := "postgresql://postgres:password@localhost:5432/todos"
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Println("Unable to connect to database:", err)
		log.Println("Attempting to start the PostgreSQL container...")
		if err := startPostgresContainer(); err != nil {
			log.Fatalf("Failed to start PostgreSQL container: %v", err)
		}

		time.Sleep(15 * time.Second)

		conn, err = pgx.Connect(context.Background(), connectionString)
		if err != nil {
			log.Fatalf("Still unable to connect to database after starting container: %v", err)
		}
	}

	fmt.Println("Connected to the database!")

	sql := `SELECT * FROM todos`
	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var todos []TodoItem

	for rows.Next() {
		var todo TodoItem
		err := rows.Scan(&todo.id, &todo.title, &todo.description, &todo.created, &todo.due)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}

	return Model{
		choices:   []string{"add", "list"},
		ToDos:     todos,
		conn:      conn,
		ViewState: ViewList,
		AddModel:  AddTodoModel{conn: conn},
	}
}

func (m AddTodoModel) Init() tea.Cmd {
	return nil
}

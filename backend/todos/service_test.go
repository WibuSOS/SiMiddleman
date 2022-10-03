package todos

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	todo, status, err := service.CreateTodos(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, todo)
}

func TestGetTodoService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Task: "task 1",
	}

	service.CreateTodos(req)

	todos, status, err := service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, false, todos[0].Done)
	assert.Equal(t, req.Task, todos[0].Task)

	req = DataRequest{
		Task: "task 2",
	}

	service.CreateTodos(req)

	todos, status, err = service.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, 2, len(todos))
	assert.Equal(t, false, todos[1].Done)
	assert.Equal(t, req.Task, todos[1].Task)
}

package todos

import (
	"strconv"

	"github.com/WibuSOS/sinarmas/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetTodos() ([]models.Todos, error)
	CreateTodos(task string) (models.Todos, error)
	CheckTodo(taskId string) error
	DeleteTodo(taskId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTodos() ([]models.Todos, error) {
	var todos []models.Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}

func (r *repository) CreateTodos(task string) (models.Todos, error) {
	todo := models.Todos{
		Task: task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		return models.Todos{}, res.Error
	}

	return todo, nil
}

func (r *repository) CheckTodo(taskId string) error {
	idConv, _ := strconv.Atoi(taskId)
	todo := models.Todos{}
	res := r.db.First(&todo, idConv)

	if res.Error != nil {
		return res.Error
	}

	if todo.Done {
		res = r.db.Model(&todo).Where("ID = ?", idConv).Update("Done", false)
	} else {
		res = r.db.Model(&todo).Where("ID = ?", idConv).Update("Done", true)
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *repository) DeleteTodo(taskId string) error {
	idConv, _ := strconv.Atoi(taskId)
	todo := models.Todos{}

	res := r.db.Where("ID = ?", idConv).Delete(&todo)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

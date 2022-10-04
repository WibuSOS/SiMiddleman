package users

import (
	"log"
	"net/http"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// func (h *Handler) GetUser(c *gin.Context) {
// 	todos, status, err := h.Service.GetTodos()
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(status, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(status, gin.H{
// 		"message": "success",
// 		"data":    todos,
// 	})
// }

func (h *Handler) CreateUser(c *gin.Context) {
	var req models.Users
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := h.Service.CreateUser(&req)
	if err != nil {
		log.Println(err.Error())
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
	})
}

// func (h *Handler) UpdateUser(c *gin.Context) {
// 	taskId := c.Param("task_id")
// 	status, err := h.Service.CheckTodo(taskId)

// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(status, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(status, gin.H{
// 		"message": "success update checklist " + taskId,
// 	})
// }

// func (h *Handler) DeleteUser(c *gin.Context) {
// 	taskId := c.Param("task_id")
// 	status, err := h.Service.DeleteTodo(taskId)

// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(status, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(status, gin.H{
// 		"message": "success delete checklist " + taskId,
// 	})
// }

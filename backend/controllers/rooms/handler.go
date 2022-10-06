package rooms

import (
	"net/http"

	"github.com/WibuSOS/sinarmas/utils/errors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		error := errors.NewBadRequestError(err.Error())
		errors.LogError(error)
		c.JSON(error.Status, gin.H{"message": error.Message})
		return
	}

	err := h.Service.CreateRoom(&req)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	// todos, status, err := h.Service.GetTodos()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(status, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(status, gin.H{
	// 	"message": "success",
	// 	"data":    todos,
	// })

	c.JSON(http.StatusOK, gin.H{
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

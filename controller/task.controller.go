package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController interface {
	Create(c *gin.Context)
}

type taskController struct{}

var (
	_taskService service.TaskService
)

func NewTaskController(taskService service.TaskService) TaskController {
	_taskService = taskService
	return &taskController{}
}
func (*taskController) Create(c *gin.Context) {
	var (
		task   entity.Task
		result entity.MongoResult
	)
	if err := json.NewDecoder(c.Request.Body).Decode(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error decoding body",
			"data":    nil,
			"error":   true,
		})
		return
	}
	if task.Name == "" || task.ListID == "" || task.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "some required fields are empty",
			"data":    nil,
			"error":   true,
		})
		return
	}
	task.ID = primitive.NewObjectID()
	task.UpdatedAt = time.Now()
	task.CreatedAt = time.Now()
	task.Done = false
	result = _taskService.Create(task)
	c.JSON(http.StatusOK, gin.H{
		"message": result.Message,
		"data":    result.InsertedID,
		"error":   false,
	})
}

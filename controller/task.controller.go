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
		input  struct {
			Name        string `bson:"name" json:"name"`
			Description string `bson:"description" json:"description"`
			UserID      string `bson:"user_id" json:"user_id"`
			ListID      string `bson:"list_id" json:"list_id"`
		}
	)
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error decoding body",
			"data":    nil,
			"error":   true,
		})
		return
	}
	if input.Name == "" || input.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "some required fields are empty",
			"data":    nil,
			"error":   true,
		})
		return
	}
	list_id, err := primitive.ObjectIDFromHex(input.ListID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	task.Name = input.Name
	task.Description = input.Description
	task.UserID = input.UserID
	task.ListID = list_id
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

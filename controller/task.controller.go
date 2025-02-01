package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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
func (*taskController) Get(c *gin.Context) {
	var (
		list_id string
		result  entity.MongoResult
	)
	list_id = c.Param("list_id")
	if list_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "list id not provided",
			"data":    nil,
			"error":   true,
		})
		return
	}
	listID, err := primitive.ObjectIDFromHex(list_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	result = _taskService.GetAll(listID)
	if !result.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Message,
			"data":    result.Data,
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result.Message,
		"data":    result.Data,
		"error":   false,
	})
}
func (*taskController) Update(c *gin.Context) {
	var (
		task_id  string
		task     entity.Task
		response entity.MongoResult
	)
	task_id = c.Param("task_id")
	if task_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "task id not provided",
			"data":    nil,
			"error":   true,
		})
		return
	}
	taskID, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	if err = json.NewDecoder(c.Request.Body).Decode(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	response = _taskService.Update(taskID, task)
	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": response.Message,
			"data":    nil,
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response.Message,
		"data":    response.Data,
		"error":   false,
	})
}
func (*taskController) Delete(c *gin.Context) {
	task_id := c.Param("task_id")
	if task_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "task id not provided",
			"data":    nil,
			"error":   true,
		})
		return
	}
	taskID, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid task ID format: %v", err.Error()),
			"data":    nil,
			"error":   true,
		})
		return
	}
	results := _taskService.DeleteTask(taskID)
	if !results.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": results.Message,
			"data":    results.Data,
			"error":   results.Success,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": results.Message,
		"data":    results.Data,
		"error":   !results.Success,
	})
}

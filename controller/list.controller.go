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

type ListController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	DeleteList(c *gin.Context)
}

type listController struct{}

var (
	_liserService service.ListService
)

func NewListController(liserService service.ListService) ListController {
	_liserService = liserService
	return &listController{}
}
func (*listController) Create(c *gin.Context) {
	var (
		list   entity.List
		result entity.MongoResult
	)

	if err := json.NewDecoder(c.Request.Body).Decode(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	list.ID = primitive.NewObjectID()
	list.UpdatedAt = time.Now()
	list.CreatedAt = time.Now()
	result = _liserService.Create(list)
	if !result.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Message,
			"data":    nil,
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result.Message,
		"data":    nil,
		"error":   false,
	})
}
func (*listController) GetAll(c *gin.Context) {
	var (
		_list []entity.List
	)
	user_id := c.Param("user_id")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user id not provided",
			"data":    nil,
			"error":   true,
		})
		return
	}
	_list, err := _liserService.FetchAll(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully fetched all items",
		"data":    _list,
		"error":   false,
	})
}
func (*listController) DeleteList(c *gin.Context) {
	var (
		list_id  string
		response entity.MongoResult
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
	listId, err := primitive.ObjectIDFromHex(list_id)
	if err != nil {
		if !response.Success {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"data":    nil,
				"error":   true,
			})
			return
		}
	}
	response = _liserService.Delete(listId)
	if !response.Success || response.Data == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": response.Message,
			"data":    response.Data,
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

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
	Create(*gin.Context)
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
			"error":   false,
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
			"error":   false,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": result.Message,
		"data":    nil,
		"error":   false,
	})
}

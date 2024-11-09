package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bonjourrog/taskm/entity"
	"github.com/bonjourrog/taskm/service"
	"github.com/bonjourrog/taskm/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController interface {
	UserRegister(c *gin.Context)
}

type authController struct{}

var (
	_authService service.AuthService
)

func NewAuthController(authService service.AuthService) AuthController {
	_authService = authService
	return &authController{}
}
func (*authController) UserRegister(c *gin.Context) {
	var (
		user    entity.User
		mResult entity.MongoResult
	)
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	if user.UserName == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "some required fields are missing",
			"data":    nil,
			"error":   true,
		})
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		return
	}
	user.ID = primitive.NewObjectID()
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	user.Password = hashedPassword
	mResult = _authService.Register(user)
	if !mResult.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": mResult.Message,
			"data":    nil,
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": mResult.Message,
		"data":    mResult.InsertedID,
		"error":   false,
	})
}

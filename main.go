package main

import (
	"os"

	"github.com/bonjourrog/taskm/controller"
	"github.com/bonjourrog/taskm/middleware"
	"github.com/bonjourrog/taskm/repository"
	"github.com/bonjourrog/taskm/repository/authrepo"
	taskrepo "github.com/bonjourrog/taskm/repository/taskRepo"
	"github.com/bonjourrog/taskm/routes"
	"github.com/bonjourrog/taskm/service"
	"github.com/joho/godotenv"
)

var (
	taskRepo       taskrepo.Task             = taskrepo.NewTasksRepo()
	taskService    service.TaskService       = service.NewTaskService(taskRepo)
	taskController controller.TaskController = controller.NewTaskController(taskService)
)

var (
	authRepo       authrepo.AuthRepo         = authrepo.NewAuthRepo()
	authService    service.AuthService       = service.NewAuthService(authRepo)
	authController controller.AuthController = controller.NewAuthController(authService)
)

var (
	listRepo       repository.ListRepo       = repository.NewListRepository()
	ListService    service.ListService       = service.NewListService(listRepo)
	listController controller.ListController = controller.NewListController(ListService)

	httpRouter routes.Router = routes.NewGinRouter()
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	httpRouter.POST("/api/auth/register/", authController.UserRegister)
	httpRouter.POST("/api/auth/sign-in", authController.Login)
	httpRouter.POST("/api/list", listController.Create, middleware.ValidateToken)
	httpRouter.DELETE("/api/list/:list_id", listController.DeleteList, middleware.ValidateToken)
	httpRouter.PUT("/api/list/:list_id", listController.UpdateList, middleware.ValidateToken)
	httpRouter.GET("/api/list/:user_id", listController.GetAll, middleware.ValidateToken)
	httpRouter.POST("/api/task/", taskController.Create, middleware.ValidateToken)
	httpRouter.SERVE(os.Getenv("PORT"))

}

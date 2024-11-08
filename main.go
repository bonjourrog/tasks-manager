package main

import (
	"os"

	"github.com/bonjourrog/taskm/controller"
	"github.com/bonjourrog/taskm/repository"
	"github.com/bonjourrog/taskm/routes"
	"github.com/bonjourrog/taskm/service"
	"github.com/joho/godotenv"
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
	httpRouter.POST("/api/list/", listController.Create)
	httpRouter.GET("/api/list/:user_id", listController.GetAll)
	httpRouter.SERVE(os.Getenv("PORT"))

}

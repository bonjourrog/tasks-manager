package routes

import (
	"net/http"

	"github.com/bonjourrog/taskm/middleware"
	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	ginDispatch = gin.Default()
)

func NewGinRouter() Router {
	ginDispatch.Use(middleware.CorsMiddleware)
	return &ginRouter{}
}
func (*ginRouter) GET(uri string, f func(*gin.Context), middlewares ...gin.HandlerFunc) {
	ginDispatch.GET(uri, append(middlewares, f)...)
}
func (*ginRouter) POST(uri string, f func(*gin.Context), middlewares ...gin.HandlerFunc) {
	ginDispatch.POST(uri, append(middlewares, f)...)
}
func (*ginRouter) PUT(uri string, f func(*gin.Context), middlewares ...gin.HandlerFunc) {
	ginDispatch.PUT(uri, append(middlewares, f)...)
}
func (*ginRouter) DELETE(uri string, f func(*gin.Context), middlewares ...gin.HandlerFunc) {
	ginDispatch.DELETE(uri, append(middlewares, f)...)
}
func (*ginRouter) SERVE(port string) {
	if err := http.ListenAndServe(port, ginDispatch); err != nil {
		panic(err)
	}
}

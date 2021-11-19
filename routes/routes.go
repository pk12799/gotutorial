package routes

import "github.com/gin-gonic/gin"

func setupRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.getTodos)
		v1.POST("todo", Controllers.CreateTodos)
		v1.GET("todo/:id", Controllers.gettodo)
		v1.PUT("todo/:id", Controllers.Update)
		v1.DELETE("todo/id", Controllers.Deletetodo)
	}
	return r
}

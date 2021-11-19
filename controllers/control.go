package controllers

import (
	"net/http"

	"../Models"
	"github.com/gin-gonic/gin"
)

func CreateATodo(c *gin.Context) {
	var todo Models.TODO
	c.BindJSON(&todo)
	// DB call to create a todo
	// Config.DB.Create(todo).Error;
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

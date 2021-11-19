//previous

// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// var router *gin.Engine

// func setup() {
// 	router = gin.Default()
// 	router.LoadHTMLGlob("templetes/*")
// 	router.GET("/", func(c *gin.Context) {

// 		// Call the HTML method of the Context to render a template
// 		c.HTML(
// 			// Set the HTTP status to 200 (OK)
// 			http.StatusOK,
// 			// Use the index.html template
// 			"index.html",
// 			// Pass the data that the page uses (in this case, 'title')
// 			gin.H{
// 				"title": "Home Page",
// 			},
// 		)

// 	})
// 	router.Run()
// }

// func main() {
// 	setup()
// }
// .......new ........



package main

r := gin.Default()
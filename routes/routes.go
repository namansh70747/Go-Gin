package routes

import (
	"github.com/gin-gonic/gin"

	"go-gin-notes/controllers" // Replace 'go-gin-notes' with your actual module name if different
)

/*
github.com/gin-gonic/gin → This brings in the Gin framework so you can define routes.
"go-gin-notes/controllers" → You're importing functions from another package where your
*/
// SetupRoutes sets up all routes for the app
// r *gin.Engine
//r *gin.Engine → You're receiving the router from main.go and adding routes to it here.
func SetupRoutes(r *gin.Engine) {
	r.GET("/notes", controllers.GetNotes)
	r.GET("/notes/:id", controllers.GetNoteByID)
	r.POST("/notes", controllers.CreateNote)
}

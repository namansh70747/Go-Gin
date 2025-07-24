package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-gin-notes/config"
	"go-gin-notes/models"
)

// GetNotes handles GET /notes and returns all notes from the database
func GetNotes(c *gin.Context) {
	var notes []models.Note
	config.DB.Find(&notes)
	c.JSON(http.StatusOK, notes)
}

// GetNoteByID handles GET /notes/:id and returns a single note by its ID
func GetNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	var note models.Note
	result := config.DB.First(&note, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

// CreateNote handles POST /notes and creates a new note in the database
func CreateNote(c *gin.Context) {
	var newNote models.Note
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newNote)
	c.JSON(http.StatusCreated, newNote)
}

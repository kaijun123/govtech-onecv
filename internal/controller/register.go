package controller

import (
	"govtech-onecv/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStudentHandler(c *gin.Context, database *db.Database) {
	var teacherSchema db.TeacherSchema

	if err := c.Bind(&teacherSchema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}

	if len(teacherSchema.Students) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}
	teacher := teacherSchema.Teacher

	// Check if the teacher is registered
	if result := database.DB.First(teacherSchema, teacher); result.Error != nil {
		// teacher does not exist in the database
		database.DB.Create(teacherSchema)

	} else {

	}

	// If not, just write

}

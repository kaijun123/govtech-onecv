package controller

import (
	"govtech-onecv/internal/db"
	"govtech-onecv/internal/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotificationHandler(c *gin.Context, database *db.Database) {
	var req Request

	if result := c.Bind(&req); result != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}
	teacher := req.Teacher
	notif := req.Notification

	if notif == "" || teacher == "" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}

	// Get students registered under a teacher
	var teacherSchema db.TeacherSchema
	if result := database.DB.First(&teacherSchema); result.Error != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("Database error; cannot fetch data"))
		return
	}
	// log.Println("teacherSchema: ", teacherSchema)

	studentMap := make(map[string]bool)

	students := teacherSchema.Students
	for _, s := range students {
		// Retrieve student info to check for suspension
		var studentSchema db.StudentSchema
		if result := database.DB.Find(&studentSchema, "student=?", s); result.Error != nil {
			// cannot find the student
			c.JSON(http.StatusBadRequest, NewErrorResponse("Database error; cannot fetch data"))
			return
		} else if studentSchema.Student != "" && !studentSchema.Suspend {
			// student exists and is not suspended, add to map
			studentMap[s] = true
		}

	}
	// log.Println("studentMap: ", studentMap)

	// Get the names specially mentioned using the @
	words := strings.Split(notif, " ")
	// log.Println("words: ", words)

	for _, w := range words {
		if strings.HasPrefix(w, "@") {
			trimmedWord := strings.TrimPrefix(w, "@")
			if _, ok := studentMap[trimmedWord]; !ok {

				// Check if the student exists and if the student is suspended
				// Do not add if the student does not exist
				var studentSchema db.StudentSchema
				if result := database.DB.Find(&studentSchema, "student=?", trimmedWord); result.Error != nil {
					c.JSON(http.StatusInternalServerError, NewErrorResponse("Database error; cannot fetch data"))
					return
				} else if studentSchema.Student != "" && !studentSchema.Suspend {
					// student exists and is not suspended, add to map
					// log.Println("trimmedWord: ", trimmedWord)
					// log.Println("studentSchema: ", studentSchema)
					studentMap[trimmedWord] = true
				}
			}
		}
	}

	studentArray := util.MapToArray(studentMap)
	// log.Println("studentArray: ", studentArray)

	c.JSON(http.StatusOK, Response{Recipients: studentArray})
}

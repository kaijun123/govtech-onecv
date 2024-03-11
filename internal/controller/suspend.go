package controller

import (
	"govtech-onecv/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuspendStudentHandler(c *gin.Context, database *db.Database) {
	var req Request

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}

	student := req.Student
	if student == "" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}
	// log.Println("Student: ", student)

	// Find the student in the db
	var data db.StudentSchema
	if result := database.DB.Find(&data, "student=?", student); result.Error != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("student does not exist"))
		return
	} else if data.Student == "" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("student does not exist"))
		return
	}

	// log.Println("data: ", data)

	if !data.Suspend {
		// Student not suspended yet
		data.Suspend = true
		if result := database.DB.Save(data); result.Error != nil {
			c.JSON(http.StatusBadRequest, NewErrorResponse("unable to update"))
			return
		} else {
			c.JSON(http.StatusNonAuthoritativeInfo, NewSuccessResponse(""))
			return
		}
	} else {
		c.JSON(http.StatusNonAuthoritativeInfo, NewSuccessResponse("student already suspended"))
		return
	}
}

package controller

import (
	"govtech-onecv/internal/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonStudentHandler(c *gin.Context, database *db.Database) {
	teacherParams, exists := c.GetQueryArray("teacher")
	if !exists {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}

	// log.Println("teacherParams: ", teacherParams)

	studentTeacherMap := make(map[string]int)
	for _, teacher := range teacherParams {
		// Do something with each teacher email

		// log.Println("teacher: ", teacher)

		var teacherSchema db.TeacherSchema
		if result := database.DB.First(&teacherSchema, "teacher=?", teacher); result.Error != nil {
			c.JSON(http.StatusBadRequest, NewErrorResponse("Database error; cannot fetch data"))
			return
		}

		students := teacherSchema.Students
		// log.Println("students: ", students)

		for _, s := range students {
			if _, ok := studentTeacherMap[s]; !ok {
				studentTeacherMap[s] = 1
			} else {
				studentTeacherMap[s] += 1
			}
		}
	}
	// log.Println("studentTeacherMap: ", studentTeacherMap)

	commonStudents := []string{}
	for student := range studentTeacherMap {
		if studentTeacherMap[student] == len(teacherParams) {
			commonStudents = append(commonStudents, student)
		}
	}

	log.Println(commonStudents)
	c.JSON(http.StatusOK, Response{Students: commonStudents})
}

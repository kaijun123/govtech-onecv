package controller

import (
	"govtech-onecv/internal/db"
	"govtech-onecv/internal/util"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterStudentHandler(c *gin.Context, database *db.Database) {
	var req Request

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}

	newStudents := req.Students
	teacher := req.Teacher
	if len(newStudents) == 0 || teacher == "" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("invalid request body"))
		return
	}

	log.Println("newStudents: ", newStudents)
	log.Println("teacher: ", teacher)

	// Use transactions
	tx := database.DB.Begin()

	// Check if the teacher exists
	var teacherSchema db.TeacherSchema
	if result := database.DB.Find(&teacherSchema, "teacher=?", teacher); result.Error != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("Database error; cannot fetch data"))
		return
	} else if teacherSchema.Teacher == "" {
		// Teacher not in the database
		log.Println("teacherSchema: ", teacherSchema)

		// Update TeacherSchema
		teacherSchema.Teacher = teacher
		teacherSchema.Students = newStudents
		if result := tx.Create(&teacherSchema); result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, NewErrorResponse("Database error; cannot write data"))
			return
		}

		// Update StudentSchema
		if status, errorResp := updateStudentSchema(tx, teacher, newStudents); status != http.StatusOK {
			c.JSON(status, errorResp)
			return
		}

	} else {
		// Teacher in the database
		log.Println("teacherSchema: ", teacherSchema)

		// Get the new student list
		oldStudents := teacherSchema.Students
		combinedStudents := util.MergeArrayWithoutDuplicate(oldStudents, newStudents)
		teacherSchema.Students = combinedStudents
		log.Println("oldStudents: ", oldStudents)
		log.Println("newStudents: ", newStudents)
		log.Println("combinedStudents: ", combinedStudents)

		// Update the TeacherSchema
		if result := tx.Save(&teacherSchema); result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, NewErrorResponse("Database error; cannot write data"))
			return
		}

		// Update the StudentSchema
		if status, errorResp := updateStudentSchema(tx, teacher, newStudents); status != http.StatusOK {
			c.JSON(status, errorResp)
			return
		}

	}
	tx.Commit()
	c.JSON(http.StatusOK, NewSuccessResponse(""))
}

// Update StudentSchema
func updateStudentSchema(tx *gorm.DB, teacher string, newStudents []string) (int, ErrorResponse) {
	for _, s := range newStudents {
		var studentSchema db.StudentSchema
		if result := tx.Find(&studentSchema, "student=?", s); result.Error != nil {
			tx.Rollback()
			return http.StatusInternalServerError, NewErrorResponse("Database error; cannot fetch data")

		} else if studentSchema.Student == "" {
			// Student does not exist. Create new student object
			studentSchema.Student = s
			studentSchema.Suspend = false
			studentSchema.Teachers = []string{teacher}
			if result := tx.Create(&studentSchema); result.Error != nil {
				tx.Rollback()
				return http.StatusInternalServerError, NewErrorResponse("Database error; cannot write data")
			}

		} else if studentSchema.Student != "" {
			// Student exists. Append to existing student object
			studentSchema.Teachers = util.AppendArrayWithoutDuplicate(studentSchema.Teachers, teacher)
			if result := tx.Save(&studentSchema); result.Error != nil {
				tx.Rollback()
				return http.StatusInternalServerError, NewErrorResponse("Database error; cannot write data")
			}
		}
	}
	return http.StatusOK, ErrorResponse{}
}

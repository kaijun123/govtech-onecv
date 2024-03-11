package main

import (
	"fmt"
	"govtech-onecv/internal/controller"
	"govtech-onecv/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("test")

	database := &db.Database{}
	database.Init()
	database.AutoMigrate()
	database.Seed()

	r := gin.Default()

	api := r.Group("/api")
	{

		api.GET("/db", func(c *gin.Context) {
			student1 := db.StudentSchema{}
			database.DB.Find(&student1)
			log.Println("student1: ", student1)
		})

		api.GET("/health", func(c *gin.Context) {
			controller.HealthController(c)
		})

		// POST /api/register
		api.POST("/register", func(c *gin.Context) {

		})

		// GET /api/commonstudents
		api.POST("/commonstudents", func(c *gin.Context) {
			controller.CommonStudentHandler(c, database)
		})

		// POST /api/suspend
		api.POST("/suspend", func(c *gin.Context) {
			controller.SuspendStudentHandler(c, database)
		})

		// POST /api/retrievefornotifications
		api.POST("/retrievefornotifications", func(c *gin.Context) {
			controller.NotificationHandler(c, database)
		})
	}

	r.Run()

}

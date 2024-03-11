package main

import (
	"govtech-onecv/internal/controller"
	"govtech-onecv/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	database := &db.Database{}
	database.Init()
	database.AutoMigrate()
	// database.Seed()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			controller.HealthController(c)
		})

		// POST /api/register
		api.POST("/register", func(c *gin.Context) {
			controller.RegisterStudentHandler(c, database)
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

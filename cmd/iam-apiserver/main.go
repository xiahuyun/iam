package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiahuyun/iam/internal/apiserver/controller"
)

func main() {

	// create gin router
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// create health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "IAM API Server is running",
		})
	})

	api := router.Group("/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", controller.ListUsers)
			users.GET("/:id", controller.GetUser)
			users.POST("", controller.CreateUser)
			users.PUT("/:id", controller.UpdateUser)
			users.DELETE("/:id", controller.DeleteUser)
		}
	}

	secureServer := &http.Server{
		Addr: ":8090",
	}
	secureServer.Handler = router

	log.Println("Starting IAM API Server on :8090...")
	if err := secureServer.ListenAndServeTLS("/etc/iam/cert/iam-apiserver.pem", "/etc/iam/cert/iam-apiserver-key.pem"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

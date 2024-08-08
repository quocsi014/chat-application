package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/quocsi014/modules/auth/handler"
	"github.com/quocsi014/modules/auth/repository"
	"github.com/quocsi014/modules/auth/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
    	if err != nil {
        	log.Fatalf("Error loading .env file")
    	}

    // Lấy biến môi trường
	dsn :=os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1Group := r.Group("/api/v1")
	{
		authGroup := v1Group.Group("/auth")
		authRepo := repository.NewAuthRepository(db)
		authService := service.NewAuthService(authRepo)
		authHandler := handler.NewAuthHandler(authService)
		authHandler.SetupRoute(authGroup)

	}
	r.Run()
}

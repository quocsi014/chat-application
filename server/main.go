package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/quocsi014/modules/auth/handler"
	"github.com/quocsi014/modules/auth/repository"
	"github.com/quocsi014/modules/auth/repository/rds"
	"github.com/quocsi014/modules/auth/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1Group := r.Group("/api/v1")
	{
		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
		authGroup := v1Group.Group("/auth")
		authRepo := repository.NewAuthRepository(db)
		accountCachingRepository := rds.NewAccountCaching(rdb, time.Minute*5)
		emailService := service.NewGEmailService()
		authService := service.NewAuthService(authRepo, accountCachingRepository, jwtSecretKey)
		authHandler := handler.NewAuthHandler(authService, *emailService)
		authHandler.SetupRoute(authGroup)

	}
	r.Run()
}

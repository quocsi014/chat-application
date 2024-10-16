package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/quocsi014/config"
	auth_handler "github.com/quocsi014/modules/auth/handler"
	auth_repository "github.com/quocsi014/modules/auth/repository"
	"github.com/quocsi014/modules/auth/repository/rds"
	auth_service "github.com/quocsi014/modules/auth/service"
	conversation_request_handler "github.com/quocsi014/modules/conversation-request/handler"
	conversation_request_repository "github.com/quocsi014/modules/conversation-request/repository"
	conversation_request_service "github.com/quocsi014/modules/conversation-request/service"
	conversation_handler "github.com/quocsi014/modules/conversation/handler"
	conversation_repository "github.com/quocsi014/modules/conversation/repository"
	conversation_service "github.com/quocsi014/modules/conversation/service"
	message_handler "github.com/quocsi014/modules/message/handler"
	message_repository "github.com/quocsi014/modules/message/repository"
	message_service "github.com/quocsi014/modules/message/service"

	"github.com/quocsi014/modules/user_information/handler"
	"github.com/quocsi014/modules/user_information/repository"
	"github.com/quocsi014/modules/user_information/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.DBConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // Phải đặt thành false khi AllowAllOrigins là true
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1Group := r.Group("/api/v1")
	{
		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
		authGroup := v1Group.Group("/auth")
		{
			authRepo := auth_repository.NewAuthRepository(db)
			accountCachingRepository := rds.NewAccountCaching(rdb, time.Minute*5)
			emailService := auth_service.NewGEmailService()
			authService := auth_service.NewAuthService(authRepo, accountCachingRepository, jwtSecretKey)
			authHandler := auth_handler.NewAuthHandler(authService, *emailService)
			authHandler.SetupRoute(authGroup)
		}

		userGroup := v1Group.Group("users")
		userRepo := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepo)
		userHandler := handler.NewUserHandler(userService)
		userHandler.SetupRoute(userGroup)

		conversationRequestGroup := v1Group.Group("/conversations")
		{
			conversationRequestRepo := conversation_request_repository.NewConversationRequestRepository(db)
			conversationRequestService := conversation_request_service.NewConversationRequestService(conversationRequestRepo, userService)
			conversationRequestHandler := conversation_request_handler.NewConversationRequestHandler(conversationRequestService)
			conversationRequestHandler.SetupRoute(conversationRequestGroup)
		}

		conversationGroup := v1Group.Group("/conversations")
		{
			conversationRepo := conversation_repository.NewConversationRepository(db)
			conversationService := conversation_service.NewConversationService(conversationRepo)
			conversationHandler := conversation_handler.NewConversationHandler(conversationService)
			conversationHandler.SetupRoute(conversationGroup)
			messageGroup := conversationGroup.Group("/:conversation_id/messages")
			{
				messageRepo := message_repository.NewMessageRepository(db)
				messageService := message_service.NewMessageService(messageRepo)
				messageHandler := message_handler.NewMessageHandler(messageService)
				messageHandler.SetupRoute(messageGroup)
			}
		}
    v1Group.GET("/ws", message_handler.HandleMessageWebSocket())

	}
	r.Run()
}

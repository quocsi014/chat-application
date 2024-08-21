package handler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/auth/entity"
	"github.com/quocsi014/modules/auth/service"
)

type IAccountService interface{
	Login(ctx context.Context, account entity.Account) (string, error)
	CreateEmailVerification(ctx context.Context, email, otp string) error
	VerifyOTP(ctx context.Context, email, otp string) (string, error)	
	Register(ctx context.Context, account *entity.Account) error
}


type AuthHandler struct{
	service IAccountService
	emailService service.EmailService
}

func NewAuthHandler(service IAccountService, emailService service.EmailService) *AuthHandler{
	return &AuthHandler{
		service: service,
		emailService: emailService,
	}
}

func (c *AuthHandler)Login() func (ctx *gin.Context){
	return func(ctx *gin.Context) {
		var account *entity.Account = &entity.Account{}
		if err := ctx.ShouldBind(account); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		jwtToken, err := c.service.Login(ctx.Request.Context(), *account)
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": jwtToken,
			})
	}
}

func genOtp() string{
	return strconv.Itoa(rand.Intn(9000) + 1000)
}

func (handler *AuthHandler)EmailRegister() func (ctx *gin.Context){
	return func(ctx *gin.Context){
		otp := genOtp()
		email := ctx.Query("email")
		if err := handler.service.CreateEmailVerification(ctx.Request.Context(), email, otp); err != nil{
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		if err := handler.emailService.SendOtp(email, otp); err != nil{
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.Status(http.StatusOK)

	}
}

type OTPVerification struct{
	Email string `json:"email"`
	Otp string `json:"otp"`
}

func (handler *AuthHandler)VerifyOTP() func(ctx *gin.Context){
	return func(ctx *gin.Context){
		otpVerification := OTPVerification{}
		if err := ctx.ShouldBind(&otpVerification); err != nil{
			fmt.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}
		token, err := handler.service.VerifyOTP(ctx, otpVerification.Email, otpVerification.Otp)
		if err != nil{
			ctx.JSON(err.(*app_error.AppError).StatusCode, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func (handler *AuthHandler)VerifyEmailVerificationOTP() func(ctx *gin.Context){
	return func(ctx *gin.Context){
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token is required"))
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Kiểm tra phương pháp ký
    			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        			return nil, app_error.ErrUnauthenticatedError(errors.New(fmt.Sprintf("Unexpected signing method %v", token.Header["alg"])), "invalid token") 
    			}
    			return []byte("your_secret_key"), nil
		})
		if err != nil || !token.Valid {
    			ctx.JSON(http.StatusUnauthorized, err)
    			ctx.Abort()
    			return
		}

		tokenClaims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
            		ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "invalid token"))
            		ctx.Abort()
            		return
        	}
		tokenEmail := tokenClaims["email"].(string)
		if tokenEmail != ctx.Param("email"){
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token is invalid"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func (handler *AuthHandler)Register() func(ctx *gin.Context){
	return func(ctx *gin.Context){
		account := entity.Account{}
		if err := ctx.ShouldBind(&account); err != nil{
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}
		if err := handler.service.Register(ctx, &account); err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
			return
		}
		ctx.Status(http.StatusCreated)

	}
}
func(handler *AuthHandler)SetupRoute(group *gin.RouterGroup){
	group.POST("/login", handler.Login())
	group.POST("/register/mail", handler.EmailRegister())
	group.POST("/register/verify_otp", handler.VerifyOTP())
	group.POST("register/:email", handler.VerifyEmailVerificationOTP(), handler.Register())
}

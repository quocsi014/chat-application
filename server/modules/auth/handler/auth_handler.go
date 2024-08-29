package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/helper"
	"github.com/quocsi014/middleware"
	"github.com/quocsi014/modules/auth/entity"
	"github.com/quocsi014/modules/auth/service"
)

type IAccountService interface{
	Login(ctx context.Context, account entity.LoginAccount) (string, error)
	GetJwtSecretKey() string
	Register(ctx context.Context, account entity.Account) (string, error)
	VerifyAccount(ctx context.Context, email string) (string, error)
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
		var account *entity.LoginAccount = &entity.LoginAccount{}
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



func (handler *AuthHandler)VerifyEmailVerificationOTP() func(ctx *gin.Context){
	return func(ctx *gin.Context){
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token is required"))
			ctx.Abort()
			return
		}
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Kiểm tra phương pháp ký
    			if token.Method != jwt.SigningMethodHS256 {
			    return nil, app_error.ErrUnauthenticatedError(errors.New(fmt.Sprintf("Unexpected signing method %v", token.Header["alg"])), "invalid token")
			}
    			return []byte(handler.service.GetJwtSecretKey()), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
        			if ve.Errors&jwt.ValidationErrorExpired != 0 {
            				ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(err, "Token has expired"))
        			} else {
         	   			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(err, "Invalid token"))
        			}
    			} else {
        			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(err, "Invalid token"))
    			}
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
		ctx.Set("token", token)
		ctx.Next()
	}
}

func (handler *AuthHandler)Register() func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		account := entity.Account{}
		if err := ctx.ShouldBind(&account); err != nil{
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
		}
		token, err:= handler.service.Register(ctx, account)
		if err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
			return
		}
		if err := handler.emailService.SendRegistrationVerification(*account.Email, token); err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
		}
		ctx.JSON(http.StatusOK, token)
	}
}
func (handler *AuthHandler)VerifyRegistration() func(ctx *gin.Context){
	return func(ctx *gin.Context){
		token, exist := ctx.Get("token")
		if !exist{
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token is required"))
			return
		}
		tokenClaims, err := helper.GetMapClaims(token.(*jwt.Token))
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		email := tokenClaims["email"].(string)
		accessToken, err := handler.service.VerifyAccount(ctx, email)
		if err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"token": accessToken,
		})

	}
}

func(handler *AuthHandler)SetupRoute(group *gin.RouterGroup){
	group.POST("/login", handler.Login())
	group.POST("/register", handler.Register())
	group.POST("/verify", middleware.VerifyToken(), handler.VerifyRegistration())
}

package handler

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quocsi014/modules/auth/entity"
	"github.com/quocsi014/modules/auth/service"
)

type IAccountService interface{
	Login(ctx context.Context, account entity.Account) (string, error)
	CreateEmailVerification(ctx context.Context, email, otp string) error
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
func(handler *AuthHandler)SetupRoute(group *gin.RouterGroup){
	group.POST("/login", handler.Login())
	group.POST("/register/mail", handler.EmailRegister())

}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quocsi014/modules/auth/entity"
)

type IAccountService interface{
	Login(ctx *gin.Context, account *entity.Account) error
}

type AuthHandler struct{
	service IAccountService
}

func NewAuthHandler(service IAccountService) *AuthHandler{
	return &AuthHandler{
		service: service,
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

		if err := c.service.Login(ctx, account); err != nil{
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":"ok",
			})
	}
}

func(handler *AuthHandler)SetupRoute(group *gin.RouterGroup){
	group.POST("/login", handler.Login())

}

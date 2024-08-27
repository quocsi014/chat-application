package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/user_information/entity"
)


type IUserService interface{
	CreateUser(ctx context.Context, user *entity.User) error
}
type UserHandler struct{
	service IUserService
}

func NewUserHandler(service IUserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler)CreateUser() func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		user := entity.User{}
		if err := ctx.ShouldBind(&user); err != nil{
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
		}
		userId,_ := ctx.Get("userId")
		user.Id = userId.(string)
		if err := handler.service.CreateUser(ctx, &user); err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
		}
		ctx.Status(http.StatusOK)
	}
}

func (handler *UserHandler)SetupRoute(group *gin.RouterGroup){
	group.POST("", handler.CreateUser())
}

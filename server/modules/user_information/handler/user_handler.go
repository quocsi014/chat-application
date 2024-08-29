package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/helper"
	"github.com/quocsi014/middleware"
	"github.com/quocsi014/modules/user_information/entity"
)


type IUserService interface{
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, userId string, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
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
		token,_ := ctx.Get("token")
		jwtMapClaims,err := helper.GetMapClaims(token.(*jwt.Token))
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, err)
		}
		id := jwtMapClaims["user_id"].(string)	


		user := entity.User{}
		if err := ctx.ShouldBind(&user); err != nil{
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
		}
		user.Id = id
		if err := handler.service.CreateUser(ctx, &user); err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
		}
		ctx.Status(http.StatusOK)
	}
}

func (handler *UserHandler)SetupRoute(group *gin.RouterGroup){
	group.POST("", middleware.VerifyToken(), handler.CreateUser())
	group.PUT("/profile", middleware.VerifyToken(), handler.UpdateProfile())
	group.GET("/profile/:username", handler.GetUserProfile())
}

func (handler *UserHandler) UpdateProfile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token, _ := ctx.Get("token")
		jwtMapClaims, err := helper.GetMapClaims(token.(*jwt.Token))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		userId := jwtMapClaims["user_id"].(string)

		var user entity.User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}

		if err := handler.service.UpdateUser(ctx, userId, &user); err != nil {
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (handler *UserHandler) GetUserProfile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		
		user, err := handler.service.GetUserByUsername(ctx, username)
		if err != nil {
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, errResponse.Err)
			return
		}
		
		ctx.JSON(http.StatusOK, user)
	}
}

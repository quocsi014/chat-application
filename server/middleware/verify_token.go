package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/helper"
)


func VerifyToken() func(ctx *gin.Context){
	return func(ctx *gin.Context) {
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
    			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
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
		ctx.Set("token", token)
		ctx.Next()
	}
}

func VerifyUser() func(*gin.Context){
	return func(ctx *gin.Context){
		token, ok := ctx.Get("token")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token is required"))
			return
		}

		jwtMapClaims,err := helper.GetMapClaims(token.(*jwt.Token))
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, err)
		}
		tokenUserId := jwtMapClaims["user_id"].(string)	
		user_id := ctx.Query("user_id")
		if tokenUserId != user_id{
			ctx.JSON(http.StatusForbidden, app_error.ErrPermissionDenied())
		}
		ctx.Next()
	}
}

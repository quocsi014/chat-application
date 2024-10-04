package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
)

func GetMapClaims(token *jwt.Token) (jwt.MapClaims, error) {
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, app_error.ErrUnauthenticatedError(nil, "invalid token")
	}
	return tokenClaims, nil
}

func GetUserId(ctx *gin.Context) string {
	token, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token không tồn tại"))
		return ""
	}

	claims, err := GetMapClaims(token.(*jwt.Token))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(err, "Unauthorized"))
		return ""
	}

	recipientId, ok := claims["user_id"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Invalid token"))
		return ""
	}

	return recipientId
}

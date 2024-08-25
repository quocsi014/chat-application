package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
)

func GetMapClaims(token *jwt.Token) (jwt.MapClaims, error){
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
            	return nil, app_error.ErrUnauthenticatedError(nil, "invalid token")
        }
	return tokenClaims, nil
}

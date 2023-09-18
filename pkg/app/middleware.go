package app

import (
	"fmt"
	"net/http"

	"github.com/AbdulrahmanDaud10/savannah-info-customer-order-service/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthorizeJWT -> to authorize JWT Token
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		fmt.Println("-------------------------", authHeader)
		// if authHeader == "" {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 		"error": "No Authorization header found"})

		// }
		tokenString := authHeader[len(BearerSchema):]

		if token, err := api.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("customerID", claims["customerID"])
					fmt.Println("during authorization", claims["customerID"])
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}

	}

}

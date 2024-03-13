package midware

import (
	"blog/database"
	"blog/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var Secret = utils.CreatConfig("auth").GetString("secret")

func checkToken(token string) (int, error) {
	if len(token) == 0 {
		return -1, errors.New("Auth token not found")
	}
	// var jwtHeader *utils.JwtHeader
	var jwtPayload *utils.JwtPayload
	var err error

	_, jwtPayload, err = utils.VertifyJWT(token, Secret)
	if err != nil {
		return -1, err
	}
	expirationTime := time.Unix(jwtPayload.Expiration, 0)
	if time.Now().After(expirationTime) {
		return -1, err
	}
	return jwtPayload.ID, nil
}

func getCookieToken(ctx *gin.Context) string {
	var refreshToken = ""
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == "refresh_token" {
			refreshToken = cookie.Value
		}
	}
	return refreshToken
}

// func chechRefresh(refreshToken string) (string,error){

// }

func AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/login" {
			return
		}

		authToken := ctx.Request.Header.Get("auth_token")
		refreshToken := getCookieToken(ctx)

		_, jwtAuthPayload, authErr := utils.VertifyJWT(authToken, Secret)
		_, jwtRefreshPayload, refreshErr := utils.VertifyJWT(refreshToken, Secret)
		if authErr == nil && refreshErr == nil {
			ctx.Set("uid", jwtAuthPayload.ID)
			return
		}
		if authErr == nil && refreshErr != nil {
			fmt.Printf("refreshErr: %v\n", refreshErr)
			newRefreshToken := utils.GetToken(jwtAuthPayload.ID)
			database.SetToken(authToken, newRefreshToken)
			ctx.SetCookie("refresh_token", newRefreshToken, int(time.Hour)*24*7,"/","localhost",false,true)
			return
		}
		if authErr != nil && refreshErr == nil {
			oldAuthToken,err := database.GetToken(refreshToken)
			if oldAuthToken != "" && err == nil{
				newAuthToken := utils.GetToken(jwtRefreshPayload.ID)
				database.SetToken(newAuthToken,refreshToken)
				ctx.Header("auth_token",newAuthToken)
				return
			}
		}
		ctx.JSON(http.StatusForbidden,gin.H{"msg":"Authorise error,please Login"})
	}
}

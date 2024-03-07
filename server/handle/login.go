package handle

import (
	"blog/database"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct{
	UserID int `json:"id"`
	Username string `json:"username"`
}

//	@Summary		"用户登陆"
//	@Description	"使用用户名和密码登录"
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username	body		loginUser	true	"用户名"
//	@Success		200			{object}	loginUser	"成功更新用户信息"
//	@Failure		400			{string}	string		"更新用户信息失败"
//	@Router			/login [post]
func Login(ctx *gin.Context) {
	var UserInfo loginUser
	var err error
	if err := ctx.BindJSON(&UserInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorInfo": "Invalid login infomation"})
		return
	}

	var User *database.USERS
	if User, err = database.GetUser(UserInfo.Username); err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusBadRequest, gin.H{"errorInfo": "User Not Found"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"errorInfo": "Login Failed"})
		}
		return
	}

	if User.PASSWORD == UserInfo.Password {
		ctx.JSON(http.StatusAccepted, response{UserID: User.ID,Username: User.USERNAME})
	} else {
		ctx.JSON(http.StatusBadRequest, "Wrong password")
		utils.LogRus.Warnf("%s登录失败", UserInfo.Username)
	}
}

package handle

import (
	// "blog/database"
	"blog/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		"根据用户id获得blog"
//	@Description	"根据用户id获得blog"
//	@Tags			blog
//	@Accept			json
//	@Produce		json
//	@Param			uid	path		string					true	"用户名"
//	@Success		200	{array}		database.BlogPreview	"获取用户的所有blog成功"
//	@Failure		400	{string}	string					"获取用户的所有blog失败"
//	@Router			/blog/{uid} [get]
func GetAllBlogByName(ctx *gin.Context) {
	var uid int
	var err error
	var response []*database.BlogPreview

	if uid, err = strconv.Atoi(ctx.Param("uid")); err != nil {
		ctx.JSON(http.StatusBadRequest, "获取参数失败")
		return
	}
	fmt.Printf("uid: %v\n", uid)
	if response, err = database.GetBlogsByWriterID(uid); err != nil {
		ctx.JSON(http.StatusBadRequest, "查询blog失败")
		return
	}
	ctx.JSON(http.StatusOK, response)
}

//	@Summary		"获得所有blog"
//	@Description	"获得所有blog"
//	@Tags			blog
//	@Accept			json
//	@Produce		json
//	@Param			uid		path		string					true	"uid"
//	@Param			blogID	path		string					true	"blogID"
//	@Success		200		{array}		database.BlogPreview	"获取blog成功"
//	@Failure		400		{string}	string					"获取blog失败"
//	@Router			/blog/{uid}/{blogID} [get]
func GetBlogContent(ctx *gin.Context) {
	var blogID int
	var err error
	var response *database.BLOGS
	if blogID, err = strconv.Atoi(ctx.Param("blogid")); err != nil {
		ctx.JSON(http.StatusBadRequest, "获取参数失败")
		return
	}
	fmt.Printf("blogID: %v\n", blogID)
	if response,err = database.GetBlogByID(blogID); err != nil {
		ctx.JSON(http.StatusBadRequest, "数据库查询失败")
		return
	}
	ctx.JSON(http.StatusOK,response)
}

//	@Summary		"创建blog"
//	@Description	"创建blog"
//	@Tags			blog
//	@Accept			json
//	@Produce		json
//	@Param			uid			path		string			true	"UID"
//	@Param			blogcontent	body		database.BLOGS	true	"blogContent"
//	@Success		200			{string}	string			"获取blog成功"
//	@Failure		400			{string}	string			"获取blog失败"
//	@Router			/blog/create/{uid} [post]
func GreateBlog(ctx *gin.Context){
	var blog database.BLOGS
	if err := ctx.BindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest,"获取blog内容失败")
		return
	}
	if err :=database.CreateBlog(&blog);err != nil {
		ctx.JSON(http.StatusBadRequest,"创建blog失败")
		return
	}
	ctx.JSON(http.StatusOK, "创建成功")
}

//	@Summary		"删除blog"
//	@Description	"删除blog"
//	@Tags			blog
//	@Accept			json
//	@Produce		json
//	@Param			uid		path		string	true	"blogID"
//	@Param			blogid	path		string	true	"blogid"
//	@Success		200		{string}	string	"获取blog成功"
//	@Failure		400		{string}	string	"获取blog失败"
//	@Router			/blog/delete/{uid}/{blogid} [delete]
func DeleteBlog(ctx *gin.Context){
	var uid int
	var blogid int
	var err error
	if uid,err = strconv.Atoi(ctx.Param("uid"));err != nil {
		ctx.JSON(http.StatusBadRequest,"删除blog失败")
		return
	}
	if blogid,err = strconv.Atoi(ctx.Param("blogid"));err != nil {
		ctx.JSON(http.StatusBadRequest,"删除blog失败")
		return
	}
	fmt.Printf("uid: %v\n", uid)
	fmt.Printf("blogid: %v\n", blogid)
	if err := database.DeleteBlogByID(blogid,uid);err != nil {
		ctx.JSON(http.StatusBadRequest,"删除blog失败")
	}
	ctx.JSON(http.StatusOK, "删除成功")
}
package test

import (
	"blog/database"
	"blog/utils"
	"fmt"
	"testing"
)

func init() {
	utils.InitLog("log")
}

func TestCreateBlog(t *testing.T) {
	Blog := database.BLOGS{
		WRITER_ID:     1,
		TITLE:         "test222",
		ARTICLE:       "article test2222",
		// CREATE_TIME:   time.Now(),
		// MODIFIED_NAME: time.Now(),
	}

	fmt.Printf("Blog: %v\n", Blog)
	database.CreateBlog(&Blog)
}

func TestDeleteBlogByID(t *testing.T) {
	err := database.DeleteBlogByID(4, 1)
	fmt.Printf("err: %v\n", err)
}

func TestUpdateBlog(t *testing.T) {
	err := database.UpdateBlog(2,"update title","update content")
	fmt.Printf("err: %v\n", err)
}

func TestGetBlogByID(t *testing.T) {
	blog,_ := database.GetBlogByID(2)
	fmt.Printf("err: %v\n", blog)
}

func TestGetBlogsByWriterID(t *testing.T) {
	blog,_ := database.GetBlogsByWriterID(5)
	for _, value := range blog {
		fmt.Printf("value: %v\n", value)
	}
}


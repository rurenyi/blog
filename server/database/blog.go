package database

import (
	"blog/utils"
	"time"
)

type BLOGS struct {
	// BLOG_ID       int       `json:"blogID"`
	WRITER_ID     int       `json:"uid"`
	TITLE         string    `json:"title"`
	ARTICLE       string    `json:"content"`
	// CREATE_TIME   time.Time `json:"createTime"`
	// MODIFIED_NAME time.Time `json:"modifiedTime"`
}

func (BLOGS) TableName() string {
	return "blogs"
}

// 增
func CreateBlog(BlogContent *BLOGS) error {
	if _, err := GetUser(BlogContent.WRITER_ID); err != nil {
		return err
	}
	db := GetdbConnection()

	if err := db.Create(BlogContent).Error; err != nil {
		return err
	}
	utils.LogRus.Warnf("用户ID:%d发布了新博客%s", BlogContent.WRITER_ID, BlogContent.TITLE)
	return nil
}

// 删
func DeleteBlogByID(BlogID int, UserID int) error {
	db := GetdbConnection()
	if err := db.Where("BLOG_ID=? AND WRITER_ID=?", BlogID, UserID).Delete(&BLOGS{}).Error; err != nil {
		return err
	}
	utils.LogRus.Warnf("用户ID%d:删除了博客%d", UserID, BlogID)
	return nil
}

// 改
func UpdateBlog(ID int, title string, text string) error {
	db := GetdbConnection()
	err := db.Model(&BLOGS{}).Where("BLOG_ID=?", ID).Updates(map[string]interface{}{"TITLE": text, "ARTICLE": text}).Error
	if err != nil {
		return err
	}
	return nil
}

// 查
func GetBlogByID(ID int) (*BLOGS, error) {
	db := GetdbConnection()
	var blog BLOGS
	if err := db.Where("BLOG_ID=?", ID).First(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

type BlogPreview struct {
	BLOG_ID     int       `json:"blogID"`
	TITLE       string    `json:"title"`
	CREATE_TIME time.Time `json:"createTime"`
}

func (BlogPreview) TableName() string {
	return "blogs"
}

func GetBlogsByWriterID(WriterID int) ([]*BlogPreview, error) {
	db := GetdbConnection()
	var blog []*BlogPreview
	if err := db.Select("BLOG_ID", "TITLE", "CREATE_TIME").Where("WRITER_ID=?", WriterID).Find(&blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

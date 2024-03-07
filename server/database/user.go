package database

import (
	"blog/utils"
	"errors"

	"gorm.io/gorm"
)

type USERS struct {
	ID       int
	USERNAME string
	PASSWORD string
}

func (USERS) TableName() string { //显示指定表名
	return "USERS"
}

// 增
func CreateNewUser(name string, password string) error {
	pass := utils.MD5(password)
	db := GetdbConnection()
	user := USERS{USERNAME: name, PASSWORD: pass}
	if err := db.Create(&user).Error; err != nil {
		utils.LogRus.Warnf("创建用户出错,%s", err)
		return err
	}
	return nil
}

// 删
func DeleteUserById(ID int) error {
	db := GetdbConnection()
	if err := db.Where("ID=?", ID).Delete(&USERS{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.LogRus.Warnf("删除用户ID %d 未找到", ID)
		} else {
			utils.LogRus.Warnf("删除用户ID %d 出错: %s", ID, err)
		}
		return err
	} else {
		utils.LogRus.Warnf("删除用户ID %d 成功", ID)
		return nil
	}
}

// 改
func UpdateUser(ID int, NewName string) error {
	if _, err := GetUser(ID); err != nil {
		return err
	}
	db := GetdbConnection()
	if err := db.Model(&USERS{}).Where("ID=?", ID).Update("USERNAME", NewName).Error; err != nil {
		return err
	}
	utils.LogRus.Warnf("%d用户更名为了%s", ID, NewName)
	return nil
}

// 查
func GetUser(UserInfo interface{}) (*USERS, error) {
	db := GetdbConnection()
	User := USERS{}
	var err error
	switch v := UserInfo.(type) {
	case int, uint:
		err = db.Where("ID=?", v).First(&User).Error
	case string:
		err = db.Where("USERNAME=?", v).First(&User).Error
	default:
		err = errors.New("未识别的查询参数")
	}
	if err != nil {
		return nil, err
	}
	return &User, nil
}

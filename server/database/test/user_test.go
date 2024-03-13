package test

import (
	"blog/database"
	"blog/utils"
	"fmt"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	utils.InitLog("log")
	database.CreateNewUser("ccc","123456")
}

func TestDeleteUserByID(t *testing.T) {
	utils.InitLog("log")
	database.DeleteUserById(4)
}

func TestUpdateUser(t *testing.T) {
	utils.InitLog("log")
	database.UpdateUser(5,"rurenyi")
}

func TestGetUser(t *testing.T) {
	utils.InitLog("log")
	if User,err := database.GetUser(1); err != nil{
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}else {
		fmt.Printf("User: %v\n", User)
	}
}
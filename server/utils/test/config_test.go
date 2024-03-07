package util_test

import (
	config "blog/utils"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := config.CreatConfig("log")
	config.WatchConfig()
	type TestConfig struct {
		A string `mapstructure:"level"`
		B string `mapstructure:"file"`
	}
	var testc TestConfig
	if err := config.Unmarshal(&testc); err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	} else {
		fmt.Println(testc.A)
		fmt.Println(testc.B)
	}
}

func TestConfig2(t *testing.T){
	logConfig := config.CreatConfig("log")
	logConfig.WatchConfig()
	if !logConfig.IsSet("level"){
		fmt.Printf("level not found")
		t.Fail()
	}else {
		Level := logConfig.GetString("level")
		fmt.Printf("Level: %v\n", Level)
	}
}

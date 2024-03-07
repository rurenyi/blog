package utils

import (
	"fmt"
	"path"
	"runtime"
	"github.com/spf13/viper"
)

var (
	ProjectPath = path.Dir(getCurrentFilePath()+"/../") + "/"
)

func getCurrentFilePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	currentDir := path.Dir(fileName)
	return currentDir
}

func CreatConfig(file string) *viper.Viper {
	config := viper.New()
	configRoot := ProjectPath + "config/"
	config.AddConfigPath(configRoot)
	config.SetConfigName(file)
	config.SetConfigType("yaml")
	configPath := configRoot + file + ".yaml"

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件: %s", configPath))
		}else {
			panic(fmt.Errorf("加载配置文件 %s 出错, 错误: %s",configPath,err))
		}
	}
	return config
}

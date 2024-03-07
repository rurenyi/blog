package util_test

import (
	"blog/utils"
	"testing"
)

func TestLogger(t *testing.T) {
	utils.InitLog("log")
	utils.LogRus.Debug("this is debug log")
	utils.LogRus.Info("this is info log")
	utils.LogRus.Warn("this is warn log")
	utils.LogRus.Error("this is error log")
	// util.LogRus.Panic("this is panic log") //写完日志之后再调panic
}


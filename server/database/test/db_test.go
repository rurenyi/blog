package test

import (
	"blog/database"
	"blog/utils"
	"sync"
	"testing"
)

func TestDBConnection(t *testing.T) {
	utils.InitLog("log")
	a := 100
	wg := sync.WaitGroup{}
	wg.Add(a)
	for i := 0; i < a; i++ {
		go func() {
			defer wg.Done()
			database.GetdbConnection()
		}()
	}
	wg.Wait()
}

func TestDBConnection2t(*testing.T) {
	utils.InitLog("log")
	database.GetdbConnection()
}

func TestDBConnection3(*testing.T) {
	utils.InitLog("log")
	database.GetRedisConnection()
}

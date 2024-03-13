package database

import (
	"blog/utils"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

var (
	db_mysql      *gorm.DB
	db_redis     *redis.Client
	db_mysql_once sync.Once
	db_redis_once sync.Once
	dblog         ormlog.Interface
)

func init() {
	dblog = ormlog.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		ormlog.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      ormlog.Info,
			Colorful:      true,
		},
	)
}

// 创建数据库连接
func createdbConnection(dbName string, user string, password string, host string, port string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true})
	if err != nil {
		utils.LogRus.Panicf("连接数据库失败,%s", dsn)
	}
	sqldb, _ := db.DB()
	sqldb.SetMaxOpenConns(100)
	sqldb.SetMaxIdleConns(20)
	utils.LogRus.Infof("数据库%s连接创建成功", dbName)
	return db
}

// 获取数据库连接
func GetdbConnection() *gorm.DB {
	db_mysql_once.Do(
		func() {
			if db_mysql == nil {
				dbviper := utils.CreatConfig("database")
				dbName := dbviper.GetString("blog.dbname")
				dbUser := dbviper.GetString("blog.username")
				dbPassword := dbviper.GetString("blog.password")
				dbHost := dbviper.GetString("blog.host")
				dbPort := dbviper.GetString("blog.port")
				db_mysql = createdbConnection(dbName, dbUser, dbPassword, dbHost, dbPort)
			}
		},
	)
	return db_mysql
}

// redis

func createRedisConnect (addr string,password string,db int) *redis.Client {
	cli := redis.NewClient(
		&redis.Options{
			Addr: addr,
			Password: password,
			DB: db,
	})
	if err := cli.Ping().Err(); err != nil {
		utils.LogRus.Panicf("redis %s 连接失败",addr)
	}else {
		utils.LogRus.Info("redis连接成功")
	}
	return cli
}

func GetRedisConnection () *redis.Client{
	db_redis_once.Do(func() {
		if db_redis == nil{
			dbviper := utils.CreatConfig("database")
			host := dbviper.GetString("redis.host")
			password := ""
			db := dbviper.GetInt("redis.database")
			db_redis = createRedisConnect(host, password, db)
		}
	})
	return db_redis
}
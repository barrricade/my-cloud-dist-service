package services

import (
	"fmt"

	"github.com/axiaoxin-com/goutils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 获取带有ctx 和自定义 logger 的 gorm db 实例
func DB() *gorm.DB {
	env := viper.GetString("env")
	fmt.Println(env, "what is env")
	// if env == "sqlite" {
	dbname := viper.GetString("sqlite3.dev.dbname")
	MaxIdleConns := viper.GetString("sqlite3.dev.max_idle_conns")
	MaxOpenConns := viper.GetString("sqlite3.dev.max_open_conns")
	fmt.Println("what is conns", MaxIdleConns, MaxOpenConns, dbname)
	if dbname == "" {
		dbname = "./cloud.db"
	}
	db, err := goutils.NewGormSQLite3(goutils.DBConfig{DBName: dbname, GormConfig: &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}, MaxIdleConns: 10, MaxOpenConns: 10})
	if err != nil {
		panic(err)
	}
	return db
	// return db.WithContext(ctx)
	// }

	// db, err := goutils.GormMySQL(env)
	// if err != nil {
	// 	panic(env + " get gorm mysql instance error:" + err.Error())
	// }
	// // logging.Debug(ctx, "using gorm mysql:"+env)
	// db = db.Session(&gorm.Session{
	// 	Logger: logging.NewGormLogger(zap.InfoLevel, zap.DebugLevel, viper.GetDuration("logging.access_logger.slow_threshold")*time.Millisecond),
	// })
	// return db
	// return db.WithContext(ctx)
}

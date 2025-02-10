package initialize

import (
	"Tiktok/app/order/global"
	"Tiktok/app/order/model"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/driver/sqlite"
)

func InitializeDB() {
	Info := global.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Info.User, Info.Password, Info.Host, Info.Port, Info.Name)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}
	zap.S().Info("数据库连接成功")

	err = global.DB.AutoMigrate(&model.Order{}, &model.OrderAddress{}, &model.OrderItem{}) // 替换为你的模型结构
    if err != nil {
        panic(err)
    }
    zap.S().Info("表迁移成功")
}



var TestDB *gorm.DB

func InitializeTestDB() {
    var err error
    TestDB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to initialize test database, got error %v", err)
    }

    err = TestDB.AutoMigrate(&model.Order{}, &model.OrderItem{}, &model.OrderAddress{})
    if err != nil {
        log.Fatalf("failed to migrate test database, got error %v", err)
    }
}

func CleanupTestDB() {
    TestDB.Exec("DROP TABLE orders")
    TestDB.Exec("DROP TABLE order_items")
    TestDB.Exec("DROP TABLE order_addresses")
}
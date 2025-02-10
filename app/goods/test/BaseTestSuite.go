package service

import (
	"Toktik/app/goods/biz/model"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseTestSuite struct {
	suite.Suite
	DB    *gorm.DB
	Trans *gorm.DB // 用于事务回滚测试数据
}

func (suite *BaseTestSuite) SetupSuite() {
	// MySQL 数据库连接配置
	dsn := "root:root@tcp(localhost:13316)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	suite.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		suite.T().Fatalf("failed to connect to database: %v", err)
	}

	// 自动迁移模型
	err = suite.DB.AutoMigrate(&model.Goods{})
	if err != nil {
		suite.T().Fatalf("failed to migrate models: %v", err)
	}
}

func (suite *BaseTestSuite) SetupTest() {
	// 每个测试方法执行前启动事务
	suite.Trans = suite.DB.Begin()
}

func (suite *BaseTestSuite) TearDownTest() {
	// 每个测试方法执行后回滚事务
	suite.Trans.Rollback()
}
func (suite *BaseTestSuite) TearDownSuite() {
	// 如果你需要在整个测试套件结束后清理，可以在这里执行清理操作
	// 例如，清理所有数据
	suite.DB.Exec("DELETE FROM goods WHERE deleted_at IS NULL")
}

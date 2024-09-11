package connect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type MySQLClientConfig struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

func NewMySQLEngine(config MySQLClientConfig) (*xorm.Engine, error) {
	if config.Port == "" {
		config.Port = "3306"
	}
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DBName)

	engine, err := xorm.NewEngine("mysql", mysqlDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	// 测试连接
	err = engine.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return engine, nil
}

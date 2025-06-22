package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "time"
    "github.com/spf13/viper"
)

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
    dsn := viper.GetString("database.dsn")
    newLogger := logger.New(
        log.New(log.Writer(), "", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second,
            LogLevel:      logger.Info,
            Colorful:      true,
        },
    )
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        return nil, err
    }
    return db, nil
}
package mysql

import (
	"context"
	"fmt"

	"github.com/grpc_plugins/config"
	"github.com/grpc_user_info/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	masterDB *gorm.DB
	slaveDB  *gorm.DB
)

const (
	Database  = "PsyExp"
	TableName = "user_infos"
)

func MasterClient() (*gorm.DB, error) {
	if masterDB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.Db.Master.User, config.Config.Db.Master.Passwd,
			config.Config.Db.Master.IP, config.Config.Db.Master.Port, Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		masterDB = db
	}
	return masterDB, nil
}

func SlaveClient() (*gorm.DB, error) {
	if slaveDB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.Db.Slave.User, config.Config.Db.Slave.Passwd,
			config.Config.Db.Slave.IP, config.Config.Db.Slave.Port, Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		slaveDB = db
	}
	return slaveDB, nil
}

type UserInfoDao interface {
	AddUser(context.Context, *entity.UserInfo) (string, error)
	QueryUser(context.Context, int64) (*entity.UserInfo, error)
	UpdateUser(context.Context, *entity.UserInfo) error
	GetUser(context.Context, string) (*entity.UserInfo, error)
	QueryUsers(context.Context, []int64) ([]*entity.UserInfo, error)
}

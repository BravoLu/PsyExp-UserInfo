package mysql

import (
	"context"
	"fmt"

	"github.com/grpc_plugins/log"
	"github.com/grpc_user_info/internal/entity"
)

type UserInfoDaoImpl struct {
}

func (s *UserInfoDaoImpl) AddUser(
	ctx context.Context,
	u *entity.UserInfo,
) (string, error) {
	tx, err := MasterClient()
	if err != nil {
		return "", err
	}
	if err := tx.WithContext(ctx).Table(TableName).Debug().Save(u).Error; err != nil {
		log.Errorf("add user error: %+v", err)
		return "", err
	}
	return "", nil
}

func (s *UserInfoDaoImpl) QueryUser(
	ctx context.Context,
	uid int64,
) (*entity.UserInfo, error) {
	tx, err := SlaveClient()
	if err != nil {
		return nil, err
	}
	var user *entity.UserInfo
	if err := tx.WithContext(ctx).
		Table(TableName).
		Debug().
		Where("id = ?", uid).
		Find(&user).Error; err != nil {
		log.Errorf("query user error: %+v", err)
		return nil, err
	}
	return user, nil
}

func (s *UserInfoDaoImpl) UpdateUser(
	ctx context.Context,
	uid int64,
	u *entity.UserInfo,
) error {
	tx, err := MasterClient()
	if err != nil {
		return err
	}
	if err := tx.WithContext(ctx).
		Table(TableName).
		Debug().
		Where("id = ?", uid).
		Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserInfoDaoImpl) GetUser(
	ctx context.Context,
	email string,
) (*entity.UserInfo, error) {
	tx, err := SlaveClient()
	if err != nil {
		return nil, err
	}
	var res *entity.UserInfo
	result := tx.WithContext(ctx).
		Table(TableName).
		Debug().
		Where("email = ?", email).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 1 {
		return res, nil
	}
	return nil, nil
}

func (s *UserInfoDaoImpl) QueryUsers(
	ctx context.Context,
	uids []int64,
) ([]*entity.UserInfo, error) {
	tx, err := SlaveClient()
	if err != nil {
		return nil, err
	}
	var res []*entity.UserInfo
	result := tx.WithContext(ctx).
		Table(TableName).
		Debug().
		Where("id in ?", uids).
		Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no valid row.")
	}
	return res, nil
}

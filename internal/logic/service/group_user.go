package service

import (
	"context"
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
)

type groupUserService struct{}

var GroupUserService = new(groupUserService)

// ListByUserId 获取用户所加入的群组
func (*groupUserService) ListByUserId(ctx context.Context, appId, userId int64) ([]model.Group, error) {
	groups, err := dao.GroupUserDao.ListByUserId(appId, userId)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// GetUsers 获取群组的所有用户信息
func (*groupUserService) GetUsers(appId, groupId int64) ([]model.GroupUser, error) {
	users, err := dao.GroupUserDao.ListUser(appId, groupId)
	if err != nil {
		return nil, err
	}
	return users, err
}

// AddUser 给群组添加用户
func (*groupUserService) AddUser(appId, groupId, userId int64, label, extra string) error {
	err := dao.GroupUserDao.Add(appId, groupId, userId, label, extra)
	if err != nil {
		return err
	}

	err = dao.GroupDao.UpdateUserNum(appId, groupId, 1)
	return err
}

// DeleteUser 从群组移除用户
func (*groupUserService) DeleteUser(appId, groupId, userId int64) error {
	err := dao.GroupUserDao.Delete(appId, groupId, userId)
	if err != nil {
		return err
	}

	err = dao.GroupDao.UpdateUserNum(appId, groupId, -1)
	return err
}

// Update 更新群组用户信息
func (*groupUserService) Update(ctx context.Context, appId, groupId int64, userId int64, label, extra string) error {
	err := dao.GroupUserDao.Update(appId, groupId, userId, label, extra)
	return err
}

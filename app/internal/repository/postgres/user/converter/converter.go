package converter

import (
	"github.com/kms-qwe/auth/internal/model"
	modelRepo "github.com/kms-qwe/auth/internal/repository/postgres/user/model"
)

// ToUserFromRepo convert serivce model to repo model
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToRepoFromUser convert serivce model to repo model
func ToRepoFromUser(user *model.User) *modelRepo.User {
	return &modelRepo.User{
		ID:        user.ID,
		Info:      ToRepoFromUserInfo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToUserInfoFromRepo convert serivce model to repo model
func ToUserInfoFromRepo(userInfo *modelRepo.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     userInfo.Role,
	}
}

// ToRepoFromUserInfo convert serivce model to repo model
func ToRepoFromUserInfo(userInfo *model.UserInfo) *modelRepo.UserInfo {
	return &modelRepo.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     userInfo.Role,
	}
}

// ToUserInfoUpdateFromRepo convert serivce model to repo model
func ToUserInfoUpdateFromRepo(userInfoUpdate *modelRepo.UserInfoUpdate) *model.UserInfoUpdate {
	return &model.UserInfoUpdate{
		ID:    userInfoUpdate.ID,
		Name:  userInfoUpdate.Name,
		Email: userInfoUpdate.Email,
		Role:  userInfoUpdate.Role,
	}
}

// ToRepoFromUserInfoUpdate convert serivce model to repo model
func ToRepoFromUserInfoUpdate(userInfoUpdate *model.UserInfoUpdate) *modelRepo.UserInfoUpdate {
	return &modelRepo.UserInfoUpdate{
		ID:    userInfoUpdate.ID,
		Name:  userInfoUpdate.Name,
		Email: userInfoUpdate.Email,
		Role:  userInfoUpdate.Role,
	}
}

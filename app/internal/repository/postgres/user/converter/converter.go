package converter

import (
	"github.com/kms-qwe/auth/internal/model"
	modelRepo "github.com/kms-qwe/auth/internal/repository/postgres/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToRepoFromUser(user *model.User) *modelRepo.User {
	return &modelRepo.User{
		ID:        user.ID,
		Info:      ToRepoFromUserInfo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(userInfo *modelRepo.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     userInfo.Role,
	}
}

func ToRepoFromUserInfo(userInfo *model.UserInfo) *modelRepo.UserInfo {
	return &modelRepo.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     userInfo.Role,
	}
}

func ToUserInfoUpdateFromRepo(userInfoUpdate *modelRepo.UserInfoUpdate) *model.UserInfoUpdate {
	return &model.UserInfoUpdate{
		ID:    userInfoUpdate.ID,
		Name:  userInfoUpdate.Name,
		Email: userInfoUpdate.Email,
		Role:  userInfoUpdate.Role,
	}
}

func ToRepoFromUserInfoUpdate(userInfoUpdate *model.UserInfoUpdate) *modelRepo.UserInfoUpdate {
	return &modelRepo.UserInfoUpdate{
		ID:    userInfoUpdate.ID,
		Name:  userInfoUpdate.Name,
		Email: userInfoUpdate.Email,
		Role:  userInfoUpdate.Role,
	}
}

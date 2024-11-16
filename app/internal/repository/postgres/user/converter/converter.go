package converter

import (
	"database/sql"
	"time"

	"github.com/kms-qwe/auth/internal/model"
	modelRepo "github.com/kms-qwe/auth/internal/repository/postgres/user/model"
)

// ToUserFromRepo convert serivce model to repo model
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: NullTimeToPtrTime(user.UpdatedAt),
	}
}

// ToRepoFromUser convert serivce model to repo model
func ToRepoFromUser(user *model.User) *modelRepo.User {
	return &modelRepo.User{
		ID:        user.ID,
		Info:      ToRepoFromUserInfo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: PtrTimeToNullTime(user.UpdatedAt),
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
	var name, email *string
	if userInfoUpdate.Name.Valid {
		nameStr := userInfoUpdate.Name.String
		name = &nameStr
	}
	if userInfoUpdate.Email.Valid {
		emailStr := userInfoUpdate.Email.String
		email = &emailStr
	}
	return &model.UserInfoUpdate{
		ID:    userInfoUpdate.ID,
		Name:  name,
		Email: email,
		Role:  userInfoUpdate.Role,
	}
}

// ToRepoFromUserInfoUpdate convert serivce model to repo model
func ToRepoFromUserInfoUpdate(userInfoUpdate *model.UserInfoUpdate) *modelRepo.UserInfoUpdate {
	return &modelRepo.UserInfoUpdate{
		ID:    userInfoUpdate.ID,
		Name:  PtrStringToNullString(userInfoUpdate.Name),
		Email: PtrStringToNullString(userInfoUpdate.Email),
		Role:  userInfoUpdate.Role,
	}
}

// PtrStringToNullString convert *string to sql.NullString
func PtrStringToNullString(s *string) sql.NullString {
	var ns sql.NullString

	if s != nil {
		ns.Valid = true
		ns.String = *s
	}

	return ns
}

// NullStringToPtrString convert sql.NullString to *string
func NullStringToPtrString(ns sql.NullString) *string {
	var s *string

	if ns.Valid {
		sringValue := ns.String
		s = &sringValue
	}

	return s
}

// PtrTimeToNullTime convert *time.Time to sql.NullTime
func PtrTimeToNullTime(t *time.Time) sql.NullTime {
	var nt sql.NullTime

	if t != nil {
		nt.Time = *t
		nt.Valid = true
	}

	return nt
}

// NullTimeToPtrTime convert sql.NullTime to *time.Time
func NullTimeToPtrTime(nt sql.NullTime) *time.Time {
	var t *time.Time

	if nt.Valid {
		timeValue := nt.Time
		t = &timeValue
	}

	return t
}

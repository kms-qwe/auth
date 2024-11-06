package converter

import (
	"database/sql"

	"github.com/kms-qwe/auth/internal/model"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromDesc(user *desc.User) *model.User {
	return &model.User{
		ID:        user.Id,
		Info:      ToUserInfoFromDesc(user.Info),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: TimestampToNullTime(user.UpdatedAt),
	}
}

func ToDescFromUser(user *model.User) *desc.User {

	return &desc.User{
		Id:        user.ID,
		Info:      ToDescFromUserInfo(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: ConvertNullTimeToTimestamp(user.UpdatedAt),
	}
}

func ToUserInfoFromDesc(userInfo *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     int32(userInfo.Role),
	}
}

func ToDescFromUserInfo(userInfo *model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     desc.Role(userInfo.Role),
	}
}

func ToUserInfoUpdateFromDesc(userInfoUpdate *desc.UserInfoUpdate) *model.UserInfoUpdate {
	return &model.UserInfoUpdate{
		ID:    userInfoUpdate.Id,
		Name:  StringValueToNullString(userInfoUpdate.Name),
		Email: StringValueToNullString(userInfoUpdate.Email),
		Role:  int32(userInfoUpdate.Role),
	}
}

func ToDescFromUserInfoUpdate(userInfoUpdate *model.UserInfoUpdate) *desc.UserInfoUpdate {
	return &desc.UserInfoUpdate{
		Id:    userInfoUpdate.ID,
		Name:  NullStringToStringValue(userInfoUpdate.Name),
		Email: NullStringToStringValue(userInfoUpdate.Email),
		Role:  desc.Role(userInfoUpdate.Role),
	}
}

func TimestampToNullTime(ts *timestamppb.Timestamp) sql.NullTime {
	if ts == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: ts.AsTime(), Valid: true}
}

func ConvertNullTimeToTimestamp(nt sql.NullTime) *timestamppb.Timestamp {
	var ts *timestamppb.Timestamp
	if nt.Valid {
		ts = timestamppb.New(nt.Time)
	}
	return ts
}

func StringValueToNullString(sv *wrapperspb.StringValue) sql.NullString {
	if sv == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: sv.String(), Valid: true}
}

func NullStringToStringValue(ns sql.NullString) *wrapperspb.StringValue {
	var sv *wrapperspb.StringValue
	if ns.Valid {
		sv = wrapperspb.String(ns.String)
	}
	return sv
}

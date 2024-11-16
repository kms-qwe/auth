package converter

import (
	"time"

	"github.com/kms-qwe/auth/internal/model"
	desc "github.com/kms-qwe/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ToUserFromAPI convert API model to service model
func ToUserFromAPI(user *desc.User) *model.User {
	return &model.User{
		ID:        user.Id,
		Info:      ToUserInfoFromAPI(user.Info),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: TimestampToPtrTime(user.UpdatedAt),
	}
}

// ToAPIFromUser convert service model to API model
func ToAPIFromUser(user *model.User) *desc.User {

	return &desc.User{
		Id:        user.ID,
		Info:      ToAPIFromUserInfo(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: PtrTimeToTimestamp(user.UpdatedAt),
	}
}

// ToUserInfoFromAPI convert API model to service model
func ToUserInfoFromAPI(userInfo *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     int32(userInfo.Role),
	}
}

// ToAPIFromUserInfo convert service model to API model
func ToAPIFromUserInfo(userInfo *model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Role:     desc.Role(userInfo.Role),
	}
}

// ToUserInfoUpdateFromAPI convert API model to service model
func ToUserInfoUpdateFromAPI(userInfoUpdate *desc.UserInfoUpdate) *model.UserInfoUpdate {
	return &model.UserInfoUpdate{
		ID:    userInfoUpdate.Id,
		Name:  StringValueToPtrString(userInfoUpdate.Name),
		Email: StringValueToPtrString(userInfoUpdate.Email),
		Role:  model.Role(userInfoUpdate.Role),
	}
}

// ToAPIFromUserInfoUpdate convert service model to API model
func ToAPIFromUserInfoUpdate(userInfoUpdate *model.UserInfoUpdate) *desc.UserInfoUpdate {
	return &desc.UserInfoUpdate{
		Id:    userInfoUpdate.ID,
		Name:  PtrStringToStringValue(userInfoUpdate.Name),
		Email: PtrStringToStringValue(userInfoUpdate.Email),
		Role:  desc.Role(userInfoUpdate.Role),
	}
}

// TimestampToPtrTime convert *timestamppb.Timestamp to *time.Time
func TimestampToPtrTime(ts *timestamppb.Timestamp) *time.Time {
	var t *time.Time
	if ts != nil {
		timeValue := ts.AsTime()
		t = &timeValue
	}
	return t
}

// PtrTimeToTimestamp convert *time.Time to *timestamppb.Timestamp
func PtrTimeToTimestamp(t *time.Time) *timestamppb.Timestamp {
	var ts *timestamppb.Timestamp
	if t != nil {
		ts = timestamppb.New(*t)
	}
	return ts
}

// StringValueToPtrString convert *wrapperspb.StringValue to *string
func StringValueToPtrString(sv *wrapperspb.StringValue) *string {
	var s *string
	if sv != nil {
		stringValue := sv.Value
		s = &stringValue
	}
	return s
}

// PtrStringToStringValue convert *string to *wrapperspb.StringValue
func PtrStringToStringValue(s *string) *wrapperspb.StringValue {
	var sv *wrapperspb.StringValue
	if s != nil {
		sv = wrapperspb.String(*s)
	}
	return sv
}

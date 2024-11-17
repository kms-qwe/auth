package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kms-qwe/auth/internal/model"
	"github.com/kms-qwe/auth/internal/repository/postgres/user/converter"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

func queryCreateUser(_ context.Context, info *model.UserInfo) (string, []interface{}, error) {
	repoInfo := converter.ToRepoFromUserInfo(info)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(repoInfo.Email, repoInfo.Email, repoInfo.Password, repoInfo.Role).
		Suffix("RETURNING id")

	return builder.ToSql()
}

func queryGetUser(_ context.Context, id int64) (string, []interface{}, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	return builder.ToSql()
}

func queryUpdateUser(_ context.Context, userInfoUpdate *model.UserInfoUpdate) (string, []interface{}, error) {
	repoUserInfoUpdate := converter.ToRepoFromUserInfoUpdate(userInfoUpdate)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, repoUserInfoUpdate.Name).
		Set(emailColumn, repoUserInfoUpdate.Email).
		Set(roleColumn, repoUserInfoUpdate.Role).
		Set(updatedAtColumn, sq.Expr("NOW()")).
		Where(sq.Eq{"id": repoUserInfoUpdate.ID})

	return builder.ToSql()
}

func queryDeleteUser(_ context.Context, id int64) (string, []interface{}, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})
	return builder.ToSql()
}

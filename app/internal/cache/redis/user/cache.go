package user

import (
	"context"
	"strconv"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/kms-qwe/auth/internal/cache"
	"github.com/kms-qwe/auth/internal/cache/redis/user/converter"
	"github.com/kms-qwe/auth/internal/config"

	modelCache "github.com/kms-qwe/auth/internal/cache/redis/user/model"
	"github.com/kms-qwe/auth/internal/model"
	client "github.com/kms-qwe/platform_common/pkg/client/cache"
)

type ch struct {
	cl  client.RedisCache
	cfg config.RedisConfig
}

// NewUserCache initializes a new redis user cache instance using the provided redis client and ttl.
func NewUserCache(cl client.RedisCache, cfg config.RedisConfig) cache.UserCache {
	return &ch{
		cl:  cl,
		cfg: cfg,
	}
}

func idToKey(id int64) string {
	return strconv.FormatInt(id, 10)
}

// Expire sets key to expire, if ttl is zero ttl sets to default from config
func (c *ch) Expire(ctx context.Context, id int64, ttl time.Duration) error {

	if ttl == 0 {
		ttl = c.cfg.TTL()
	}

	key := idToKey(id)

	err := c.cl.Expire(ctx, key, ttl)
	if err != nil {
		return err
	}

	return nil
}

// Set sets user to cache
func (c *ch) Set(ctx context.Context, user *model.User) error {
	cacheUser := converter.ToCacheFromUser(user)

	key := idToKey(cacheUser.ID)
	err := c.cl.HashSet(ctx, key, cacheUser)
	if err != nil {
		return err
	}

	return nil
}

// Get gets user from cache
func (c *ch) Get(ctx context.Context, id int64) (*model.User, error) {
	key := idToKey(id)

	values, err := c.cl.HGetAll(ctx, key)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, model.ErrorUserNotFound
	}

	var user modelCache.User
	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromCache(&user), nil
}

// Delete deletes user from cache
func (c *ch) Delete(ctx context.Context, id int64) error {
	key := idToKey(id)

	err := c.cl.Delete(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

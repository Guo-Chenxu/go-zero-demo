package model

import (
	"context"
	"fmt"
	// "github.com/SpectatorNan/gorm-zero/gormc"
	"gorm.io/gorm"
)

type (
	myUserModel interface {
		FindById(ctx context.Context, id int64) (*User, error)
	}
)

func (m *defaultUserModel) FindById(ctx context.Context, id int64) (*User, error) {
	// 缓存的键
	gozeroStudyUserIdKey := fmt.Sprintf("%s%v", cacheGozeroStudyUserIdPrefix, id)
	
	var resp User
	err := m.QueryCtx(ctx, &resp, gozeroStudyUserIdKey, func(conn *gorm.DB, v interface{}) error {
		err := conn.Model(&User{}).Where("id = ?", id).First(&resp).Error
		return err
	})
	return &resp, err
}

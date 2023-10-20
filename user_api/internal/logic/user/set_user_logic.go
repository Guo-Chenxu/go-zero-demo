package user

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"user_api/model"

	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserLogic {
	return &SetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserLogic) SetUser(req *types.SetUserReq) (resp string, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("req:  ", req)
	user := &model.User{
		Name:     sql.NullString{String: req.Name, Valid: len(req.Name) != 0},
		Password: req.Password,
		Mobile:   req.Mobile,
		Type:     req.Type,
		CreateAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdateAt: time.Now(),
	}
	fmt.Println("user:  ", user)
	// todo 测试事务
	err = l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		l.svcCtx.UserModel.Insert(l.ctx, db, user)
		if user.Name.String == "张四" {
			return gorm.ErrDuplicatedKey
		}
		return nil
	})

	resp = strconv.Itoa(int(user.Id))
	return
}

package user

import (
	"context"
	"database/sql"
	"fmt"

	"gozero_study/internal/svc"
	"gozero_study/internal/types"
	"gozero_study/model"

	"github.com/zeromicro/go-zero/core/logx"
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
	user := &model.User{Name: sql.NullString{String: req.Name, Valid: len(req.Name) != 0}, Password: req.Password, Mobile: req.Mobile, Type: req.Type}
	fmt.Println("user:  ", user)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	resp = "插入成功"
	return
}

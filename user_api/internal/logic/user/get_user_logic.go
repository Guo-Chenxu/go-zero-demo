package user

import (
	"context"

	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserRep, err error) {
	// todo: add your logic here and delete this line
	// user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	user, err := l.svcCtx.UserModel.FindById(l.ctx, req.Id)
	if user == nil {
		return
	}
	return &types.GetUserRep{
		Id:       user.Id,
		Name:     user.Name.String,
		Password: user.Password,
		Mobile:   user.Mobile,
		Type:     user.Type,
		CreateAt: user.CreateAt.Time.String(),
		UpdateAt: user.UpdateAt.String(),
	}, err
}

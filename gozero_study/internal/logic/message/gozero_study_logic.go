package message

import (
	"context"

	"gozero_study/internal/svc"
	"gozero_study/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Gozero_studyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGozero_studyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Gozero_studyLogic {
	return &Gozero_studyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Gozero_studyLogic) Gozero_study(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: req.Name,
	}, err
}

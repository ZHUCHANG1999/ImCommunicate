package logic

import (
	"context"
	"imooc/traning/user/models"

	"imooc/traning/user/rpc/internal/svc"
	"imooc/traning/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Id:    in.Id,
		Name:  in.Name,
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &user.CreateResp{}, nil
}

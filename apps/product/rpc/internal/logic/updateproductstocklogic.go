package logic

import (
	"context"

	"go-zero-mall/apps/product/rpc/internal/svc"
	"go-zero-mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStockLogic {
	return &UpdateProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStockLogic) UpdateProductStock(in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.UpdateProductStockResponse{}, nil
}

package service

import (
	"authentication-center/global"
	"authentication-center/internal/dao"
	"authentication-center/pkg/otgorm"

	"context"
)

/**
* @Author: super
* @Date: 2020-09-22 09:41
* @Description: 统一配置context
**/

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}

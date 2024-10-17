package impl

import (
	"context"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger"
)

func (i *FormServiceImpl) CreateForm(ctx context.Context, ins *form.Form) (*form.Form, error) {
	// 直接打印日志
	i.l.Debug("create form")
	// 带Format的日志打印, fmt.Sprintf()
	i.l.Debugf("create form: %s", ins.Name)
	// 携带额外meta数据, 常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create form with meta kv")
	return ins, nil
}

func (i *FormServiceImpl) QueryForm(ctx context.Context, req *form.QueryFormRequest) (
	*form.FormSet, error) {
	return nil, nil
}

func (i *FormServiceImpl) DescribeForm(ctx context.Context, req *form.QueryFormRequest) (
	*form.Form, error) {
	return nil, nil
}

func (i *FormServiceImpl) UpdateForm(ctx context.Context, req *form.UpdateFormRequest) (
	*form.Form, error) {
	return nil, nil
}

func (i *FormServiceImpl) DeleteForm(ctx context.Context, req *form.DeleteFormRequest) (
	*form.Form, error) {
	return nil, nil
}

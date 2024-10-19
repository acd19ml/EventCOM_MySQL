package impl

import (
	"context"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
)

func (i *FormServiceImpl) CreateForm(ctx context.Context, ins *form.Form) (*form.Form, error) {
	// // 直接打印日志
	// i.l.Named("Create").Debug("create form")
	// i.l.Info("create form")
	// // 带Format的日志打印, fmt.Sprintf()
	// i.l.Debugf("create form: %s", ins.Name)
	// // 携带额外meta数据, 常用于Trace系统
	// i.l.With(logger.NewAny("request-id", "req01")).Debug("create form with meta kv")

	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 默认值注入
	ins.InjectDefault()

	// dao模块，负责把对象入库
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

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

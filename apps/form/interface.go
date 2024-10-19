package form

import "context"

// form app service 的接口定义
type Service interface {
	// 创建表格
	CreateForm(context.Context, *Form) (*Form, error)
	// 查询表格列表
	QueryForm(context.Context, *QueryFormRequest) (*FormSet, error)
	// 查询表格详情
	DescribeForm(context.Context, *QueryFormRequest) (*Form, error)
	// 表格更新
	UpdateForm(context.Context, *UpdateFormRequest) (*Form, error)
	// 表格删除, 比如前端需要 打印当前删除表格的名称或者其他信息
	DeleteForm(context.Context, *DeleteFormRequest) (*Form, error)
}

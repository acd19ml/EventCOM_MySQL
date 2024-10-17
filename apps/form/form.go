package form

import (
	"context"
	"time"
)

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

type FormSet struct {
	Items []*Form
	Total int
}

type FieldSet struct {
	Fields []*Field
}

func NewForm() *Form {
	return &Form{
		Head:     &Head{},
		FieldSet: &FieldSet{},
	}
}

// Form模型的定义
type Form struct {
	// 表格公共属性部分
	*Head
	// 表格独有属性部分
	*FieldSet
}

type Head struct {
	ID        uint64    `json:"id"`         // 表单ID
	Name      string    `json:"name"`       // 表单名称
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

type Field struct {
	ID                uint64     `json:"id"`                           // 字段ID
	Label             string     `json:"label"`                        // 字段标签
	Type              FieldType  `json:"-"`                            // 字段类型 (text, number, selection, date)
	Required          bool       `json:"required"`                     // 是否必填
	Description       string     `json:"description,omitempty"`        // 字段描述
	MinValue          *int       `json:"min_value,omitempty"`          // 数字类型的最小值
	MaxValue          *int       `json:"max_value,omitempty"`          // 数字类型的最大值
	MinDate           *time.Time `json:"min_date,omitempty"`           // 日期类型的最小日期
	MaxDate           *time.Time `json:"max_date,omitempty"`           // 日期类型的最大日期
	MultipleSelection bool       `json:"multiple_selection,omitempty"` // 是否允许多选
	Options           []string   `json:"options,omitempty"`            // 选择类型的可选项
	CreatedAt         time.Time  `json:"created_at"`                   // 创建时间
	UpdatedAt         time.Time  `json:"updated_at"`                   // 更新时间
}

type QueryFormRequest struct {
}

type UpdateFormRequest struct {
	Name string `json:"name"` // 表单名称
	*Field
}

type DeleteFormRequest struct {
	Id string
}

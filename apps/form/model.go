package form

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type FormSet struct {
	Items []*Form
	Total int
}

func NewForm() *Form {
	return &Form{
		Head:     &Head{},
		FieldSet: []*Field{},
	}
}

// Form模型的定义
type Form struct {
	// 表格公共属性部分
	*Head
	// 表格独有属性部分
	FieldSet []*Field
}

func (f *Form) Validate() error {
	return validate.Struct(f)
}

func (f *Form) InjectDefault() {
	if f.CreatedAt == 0 {
		f.CreatedAt = time.Now().UnixMilli()
	}
}

type Head struct {
	Id        string `json:"id" validate:"required"`   // 表单ID
	Name      string `json:"name" validate:"required"` // 表单名称
	CreatedAt int64  `json:"created_at"`               // 创建时间
	UpdatedAt int64  `json:"updated_at"`               // 更新时间
}

type Field struct {
	Id                string   `json:"id" validate:"required"`       // 字段ID
	Head_Id           string   `json:"head_id" validate:"required"`  // 字段ID
	Label             string   `json:"label" validate:"required"`    // 字段标签
	Type              string   `json:"type" validate:"required"`     // 字段类型 (text, number, selection, date)
	Required          bool     `json:"required"`                     // 是否必填
	Description       string   `json:"description,omitempty"`        // 字段描述
	MinValue          int64    `json:"min_value,omitempty"`          // 数字类型的最小值
	MaxValue          int64    `json:"max_value,omitempty"`          // 数字类型的最大值
	MinDate           int64    `json:"min_date,omitempty"`           // 日期类型的最小日期
	MaxDate           int64    `json:"max_date,omitempty"`           // 日期类型的最大日期
	MultipleSelection bool     `json:"multiple_selection,omitempty"` // 是否允许多选
	Options           []string `json:"options,omitempty"`            // 选择类型的可选项
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

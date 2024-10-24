package form

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func NewFormSet() *FormSet {
	return &FormSet{
		Items: []*Form{},
	}
}

type FormSet struct {
	Total int     `json:"total"`
	Items []*Form `json:"forms"`
}

func (s *FormSet) Add(item *Form) {
	s.Items = append(s.Items, item)
}

func (f *Form) AddField(field *Field) {
	f.FieldSet = append(f.FieldSet, field)
}
func NewForm() *Form {
	return &Form{
		Head:     &Head{},
		FieldSet: []*Field{},
	}
}

func NewField() *Field {
	return &Field{}
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

func NewQueryFormFromHTTP(r *http.Request) *QueryFormRequest {
	req := NewQueryFormRequest()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryFormRequest() *QueryFormRequest {
	return &QueryFormRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type QueryFormRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func (q *QueryFormRequest) OffSet() int64 {
	return int64((q.PageNumber - 1) * q.PageSize)
}

func (q *QueryFormRequest) GetPageSize() uint {
	return uint(q.PageSize)
}

func NewDescribeFormRequestWithId(id string) *DescribeFormRequest {
	return &DescribeFormRequest{
		Id: id,
	}
}

type DescribeFormRequest struct {
	Id string
}

type UpdateFormRequest struct {
	Name string `json:"name"` // 表单名称
	*Field
}

type DeleteFormRequest struct {
	Id string
}

package form

type FieldType interface {
	Validate(input interface{}) error // 验证字段输入
	GetType() string                  // 获取字段类型名称
}

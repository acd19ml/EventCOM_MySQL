package impl

import (
	"fmt"
	"time"
)

// 文本字段类型
type TextField struct{}

func (f TextField) Validate(input interface{}) error {
	_, ok := input.(string)
	if !ok {
		return fmt.Errorf("invalid input type, expected string")
	}
	return nil
}

func (f TextField) GetType() string {
	return "text"
}

// 数字字段类型
type NumberField struct {
	Min int
	Max int
}

func (f NumberField) Validate(input interface{}) error {
	val, ok := input.(int)
	if !ok {
		return fmt.Errorf("invalid input type, expected number")
	}
	if val < f.Min || val > f.Max {
		return fmt.Errorf("value out of range")
	}
	return nil
}

func (f NumberField) GetType() string {
	return "number"
}

// 选择类型的字段
type SelectionField struct {
	Options           []string // 可供选择的选项
	MultipleSelection bool     // 是否允许多选
}

// 验证选择类型字段的输入
func (f SelectionField) Validate(input interface{}) error {
	// 单选的情况，输入应该是一个字符串
	if !f.MultipleSelection {
		value, ok := input.(string)
		if !ok {
			return fmt.Errorf("invalid input type, expected string")
		}
		// 检查输入的值是否在选项列表中
		if !contains(f.Options, value) {
			return fmt.Errorf("invalid selection: %s", value)
		}
		return nil
	}

	// 多选的情况，输入应该是一个字符串切片
	values, ok := input.([]string)
	if !ok {
		return fmt.Errorf("invalid input type, expected []string")
	}
	// 检查每个输入值是否在选项列表中
	for _, v := range values {
		if !contains(f.Options, v) {
			return fmt.Errorf("invalid selection: %s", v)
		}
	}
	return nil
}

// 返回字段类型
func (f SelectionField) GetType() string {
	return "selection"
}

// 检查选项列表中是否包含某个值
func contains(options []string, value string) bool {
	for _, option := range options {
		if option == value {
			return true
		}
	}
	return false
}

// 表示日期类型的字段
type DateField struct {
	MinDate *time.Time // 允许的最小日期
	MaxDate *time.Time // 允许的最大日期
}

// 验证日期类型字段的输入
func (f DateField) Validate(input interface{}) error {
	value, ok := input.(time.Time)
	if !ok {
		return fmt.Errorf("invalid input type, expected time.Time")
	}
	// 如果设置了最小日期，检查输入是否小于最小日期
	if f.MinDate != nil && value.Before(*f.MinDate) {
		return fmt.Errorf("date is too early, must be after %s", f.MinDate.Format("2006-01-02"))
	}
	// 如果设置了最大日期，检查输入是否大于最大日期
	if f.MaxDate != nil && value.After(*f.MaxDate) {
		return fmt.Errorf("date is too late, must be before %s", f.MaxDate.Format("2006-01-02"))
	}
	return nil
}

// 返回字段类型
func (f DateField) GetType() string {
	return "date"
}

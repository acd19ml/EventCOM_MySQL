package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/apps/form/impl"
	"github.com/acd19ml/EventCOM_MySQL/conf"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
)

var (
	// 定义对象是满足该接口的实例
	service form.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)

	// 创建一个Form对象
	ins := form.NewForm()
	ins.Head.Id = "id_01"
	ins.Head.Name = "test form"

	// 创建多个 Field 并加入到 FieldSet 中
	ins.FieldSet = []*form.Field{
		{
			Id:                "01",
			Head_Id:           "id_01", // 与 Head 关联的 ID
			Label:             "Field 1",
			Type:              impl.NumberField{}.GetType(),
			Required:          true,
			Description:       "This is field 1",
			MinValue:          0,
			MaxValue:          100,
			MinDate:           0,
			MaxDate:           0,
			MultipleSelection: false,
			Options:           []string{},
		},
		{
			Id:                "02",
			Head_Id:           "id_01", // 与 Head 关联的 ID
			Label:             "Field 2",
			Type:              impl.SelectionField{}.GetType(),
			Required:          true,
			Description:       "This is field 2",
			MinValue:          0,
			MaxValue:          0,
			MinDate:           0,
			MaxDate:           0,
			MultipleSelection: false,
			Options:           []string{"Option 1", "Option 2"},
		},
	}

	// 调用 CreateForm 方法，将 Form 和 FieldSet 插入数据库
	ins, err := service.CreateForm(context.Background(), ins)

	// 验证是否没有错误
	if should.NoError(err) {
		// 输出插入后的结果
		fmt.Printf("Inserted Form: %+v\n", ins)
	}

	// 检查插入的 Form 是否符合预期
	should.Equal("test form", ins.Name)
	should.Len(ins.FieldSet, 2)

	// 验证每个 Field
	should.Equal("Field 1", ins.FieldSet[0].Label)
	should.Equal("Field 2", ins.FieldSet[1].Label)

}

func init() {
	// 测试用例的配置文件
	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
	if err != nil {
		panic(err)
	}
	//需要初始化全局logger
	//为什么不设计为默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())

	// form service 的具体实现
	service = impl.NewFormServiceImpl()
}

package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/mcube/logger"
	"github.com/acd19ml/EventCOM_MySQL/mcube/sqlbuilder"
)

func (i *FormServiceImpl) CreateForm(ctx context.Context, ins *form.Form) (*form.Form, error) {
	// // 直接打印日志
	i.l.Named("Create").Debug("create form")
	i.l.Info("create form")
	// 带Format的日志打印, fmt.Sprintf()
	i.l.Debugf("create form: %s", ins.Name)
	// // 携带额外meta数据, 常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create form with meta kv")

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

	b := sqlbuilder.NewBuilder(QueryFormSQL)
	if req.Keywords != "" {
		b.Where("h.`name`LIKE ?",
			"%"+req.Keywords+"%",
		)
	}

	b.Limit(req.OffSet(), req.GetPageSize())

	querySQL, args := b.Build()
	i.l.Debugf("query sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := form.NewFormSet()
	for rows.Next() {
		// 每扫描一行，就读出一行数据
		ins := form.NewForm()

		if err := rows.Scan(
			&ins.Id, &ins.Name,
		); err != nil {
			return nil, err
		}

		set.Add(ins)
	}

	// total 统计
	countSQL, args := b.BuildCount()
	i.l.Debugf("count sql: %s, args: %v", countSQL, args)
	countStmt, err := i.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer countStmt.Close()
	// 执行count语句
	if err := countStmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}

	return set, nil
}

func (i *FormServiceImpl) DescribeForm(ctx context.Context, req *form.DescribeFormRequest) (
	*form.Form, error) {
	b := sqlbuilder.NewBuilder(DescribeFormSQL)
	if req.Id != "" {
		b.Where("h.`id` = ?",
			req.Id,
		)
	}

	querySQL, args := b.Build()
	i.l.Debugf("query sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ins := form.NewForm()
	for rows.Next() {
		// 每扫描一行，就读出一行数据
		field := form.NewField()
		var optionsJSON []byte // 用于接收 JSON 格式的 options 字段

		if err := rows.Scan(
			&ins.Id, &ins.Name, &field.Label, &field.Type, &field.Required,
			&field.Description, &field.MinValue, &field.MaxValue, &field.MinDate,
			&field.MaxDate, &field.MultipleSelection, &optionsJSON,
		); err != nil {
			return nil, err
		}
		// 将 JSON 字符串转换为 []string
		if err := json.Unmarshal(optionsJSON, &field.Options); err != nil {
			return nil, fmt.Errorf("failed to unmarshal options: %v", err)
		}

		ins.AddField(field)
	}

	return ins, nil
}

func (i *FormServiceImpl) UpdateForm(ctx context.Context, req *form.UpdateFormRequest) (
	*form.Form, error) {
	return nil, nil
}

func (i *FormServiceImpl) DeleteForm(ctx context.Context, req *form.DeleteFormRequest) (
	*form.Form, error) {
	return nil, nil
}

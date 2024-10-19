package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
)

func (i *FormServiceImpl) save(ctx context.Context, ins *form.Form) error {
	var (
		err error
	)

	// 把数据入库到Head与Field表
	// 一次需要向两个表录入数据，开启一个事务，要么都成功，要么都失败
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	// 通过defer处理事务提交
	// 1. 没有报错Commit
	// 2. 有报错Rollback
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error, %s", err)
			}
		}
	}()

	// 插入Head数据
	hstmt, err := tx.PrepareContext(ctx, InsertHeadSQL)
	if err != nil {
		return err
	}
	defer hstmt.Close()

	r, err := hstmt.ExecContext(ctx,
		ins.Head.Id, ins.Head.Name, ins.Head.CreatedAt, ins.Head.UpdatedAt,
	)
	if err != nil {
		return err
	} else {
		fmt.Printf("insert head success, %v", r)
	}

	// 插入Field数据
	fstmt, err := tx.PrepareContext(ctx, InsertFieldSQL)
	if err != nil {
		return err
	}
	defer fstmt.Close()

	for _, field := range ins.FieldSet {
		// 插入的 options 字段为 []string
		optionsJSON, err := json.Marshal(field.Options)
		if err != nil {
			return fmt.Errorf("failed to marshal options: %v", err)
		}
		f, err := fstmt.ExecContext(ctx,
			field.Id, field.Head_Id, field.Label, field.Type, field.Required, field.Description,
			field.MinValue, field.MaxValue, field.MinDate, field.MaxDate, field.MultipleSelection, string(optionsJSON),
		)
		if err != nil {
			return err
		} else {
			fmt.Printf("insert field success, %v", f)
		}
	}
	return nil
}

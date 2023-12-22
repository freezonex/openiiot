// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query_openiiot

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"freezonex/openiiot/biz/dal/model_openiiot"
)

func newFlowCore(db *gorm.DB, opts ...gen.DOOption) flowCore {
	_flowCore := flowCore{}

	_flowCore.flowCoreDo.UseDB(db, opts...)
	_flowCore.flowCoreDo.UseModel(&model_openiiot.FlowCore{})

	tableName := _flowCore.flowCoreDo.TableName()
	_flowCore.ALL = field.NewAsterisk(tableName)
	_flowCore.ID = field.NewInt64(tableName, "id")
	_flowCore.FlowID = field.NewInt64(tableName, "flow_id")
	_flowCore.CoreID = field.NewInt64(tableName, "core_id")
	_flowCore.Script = field.NewString(tableName, "script")

	_flowCore.fillFieldMap()

	return _flowCore
}

type flowCore struct {
	flowCoreDo flowCoreDo

	ALL    field.Asterisk
	ID     field.Int64
	FlowID field.Int64
	CoreID field.Int64
	Script field.String

	fieldMap map[string]field.Expr
}

func (f flowCore) Table(newTableName string) *flowCore {
	f.flowCoreDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f flowCore) As(alias string) *flowCore {
	f.flowCoreDo.DO = *(f.flowCoreDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *flowCore) updateTableName(table string) *flowCore {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.FlowID = field.NewInt64(table, "flow_id")
	f.CoreID = field.NewInt64(table, "core_id")
	f.Script = field.NewString(table, "script")

	f.fillFieldMap()

	return f
}

func (f *flowCore) WithContext(ctx context.Context) *flowCoreDo { return f.flowCoreDo.WithContext(ctx) }

func (f flowCore) TableName() string { return f.flowCoreDo.TableName() }

func (f flowCore) Alias() string { return f.flowCoreDo.Alias() }

func (f flowCore) Columns(cols ...field.Expr) gen.Columns { return f.flowCoreDo.Columns(cols...) }

func (f *flowCore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *flowCore) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 4)
	f.fieldMap["id"] = f.ID
	f.fieldMap["flow_id"] = f.FlowID
	f.fieldMap["core_id"] = f.CoreID
	f.fieldMap["script"] = f.Script
}

func (f flowCore) clone(db *gorm.DB) flowCore {
	f.flowCoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f flowCore) replaceDB(db *gorm.DB) flowCore {
	f.flowCoreDo.ReplaceDB(db)
	return f
}

type flowCoreDo struct{ gen.DO }

func (f flowCoreDo) Debug() *flowCoreDo {
	return f.withDO(f.DO.Debug())
}

func (f flowCoreDo) WithContext(ctx context.Context) *flowCoreDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f flowCoreDo) ReadDB() *flowCoreDo {
	return f.Clauses(dbresolver.Read)
}

func (f flowCoreDo) WriteDB() *flowCoreDo {
	return f.Clauses(dbresolver.Write)
}

func (f flowCoreDo) Session(config *gorm.Session) *flowCoreDo {
	return f.withDO(f.DO.Session(config))
}

func (f flowCoreDo) Clauses(conds ...clause.Expression) *flowCoreDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f flowCoreDo) Returning(value interface{}, columns ...string) *flowCoreDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f flowCoreDo) Not(conds ...gen.Condition) *flowCoreDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f flowCoreDo) Or(conds ...gen.Condition) *flowCoreDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f flowCoreDo) Select(conds ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f flowCoreDo) Where(conds ...gen.Condition) *flowCoreDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f flowCoreDo) Order(conds ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f flowCoreDo) Distinct(cols ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f flowCoreDo) Omit(cols ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f flowCoreDo) Join(table schema.Tabler, on ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f flowCoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f flowCoreDo) RightJoin(table schema.Tabler, on ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f flowCoreDo) Group(cols ...field.Expr) *flowCoreDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f flowCoreDo) Having(conds ...gen.Condition) *flowCoreDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f flowCoreDo) Limit(limit int) *flowCoreDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f flowCoreDo) Offset(offset int) *flowCoreDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f flowCoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *flowCoreDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f flowCoreDo) Unscoped() *flowCoreDo {
	return f.withDO(f.DO.Unscoped())
}

func (f flowCoreDo) Create(values ...*model_openiiot.FlowCore) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f flowCoreDo) CreateInBatches(values []*model_openiiot.FlowCore, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f flowCoreDo) Save(values ...*model_openiiot.FlowCore) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f flowCoreDo) First() (*model_openiiot.FlowCore, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowCore), nil
	}
}

func (f flowCoreDo) Take() (*model_openiiot.FlowCore, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowCore), nil
	}
}

func (f flowCoreDo) Last() (*model_openiiot.FlowCore, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowCore), nil
	}
}

func (f flowCoreDo) Find() ([]*model_openiiot.FlowCore, error) {
	result, err := f.DO.Find()
	return result.([]*model_openiiot.FlowCore), err
}

func (f flowCoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.FlowCore, err error) {
	buf := make([]*model_openiiot.FlowCore, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f flowCoreDo) FindInBatches(result *[]*model_openiiot.FlowCore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f flowCoreDo) Attrs(attrs ...field.AssignExpr) *flowCoreDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f flowCoreDo) Assign(attrs ...field.AssignExpr) *flowCoreDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f flowCoreDo) Joins(fields ...field.RelationField) *flowCoreDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f flowCoreDo) Preload(fields ...field.RelationField) *flowCoreDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f flowCoreDo) FirstOrInit() (*model_openiiot.FlowCore, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowCore), nil
	}
}

func (f flowCoreDo) FirstOrCreate() (*model_openiiot.FlowCore, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowCore), nil
	}
}

func (f flowCoreDo) FindByPage(offset int, limit int) (result []*model_openiiot.FlowCore, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f flowCoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f flowCoreDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f flowCoreDo) Delete(models ...*model_openiiot.FlowCore) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *flowCoreDo) withDO(do gen.Dao) *flowCoreDo {
	f.DO = *do.(*gen.DO)
	return f
}

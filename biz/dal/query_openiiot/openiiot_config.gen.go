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

func newOpeniiotConfig(db *gorm.DB, opts ...gen.DOOption) openiiotConfig {
	_openiiotConfig := openiiotConfig{}

	_openiiotConfig.openiiotConfigDo.UseDB(db, opts...)
	_openiiotConfig.openiiotConfigDo.UseModel(&model_openiiot.OpeniiotConfig{})

	tableName := _openiiotConfig.openiiotConfigDo.TableName()
	_openiiotConfig.ALL = field.NewAsterisk(tableName)
	_openiiotConfig.Pkey = field.NewInt32(tableName, "pkey")
	_openiiotConfig.Name = field.NewString(tableName, "name")
	_openiiotConfig.Value = field.NewString(tableName, "value")
	_openiiotConfig.AddDate = field.NewTime(tableName, "add_date")

	_openiiotConfig.fillFieldMap()

	return _openiiotConfig
}

type openiiotConfig struct {
	openiiotConfigDo openiiotConfigDo

	ALL     field.Asterisk
	Pkey    field.Int32
	Name    field.String
	Value   field.String
	AddDate field.Time

	fieldMap map[string]field.Expr
}

func (o openiiotConfig) Table(newTableName string) *openiiotConfig {
	o.openiiotConfigDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o openiiotConfig) As(alias string) *openiiotConfig {
	o.openiiotConfigDo.DO = *(o.openiiotConfigDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *openiiotConfig) updateTableName(table string) *openiiotConfig {
	o.ALL = field.NewAsterisk(table)
	o.Pkey = field.NewInt32(table, "pkey")
	o.Name = field.NewString(table, "name")
	o.Value = field.NewString(table, "value")
	o.AddDate = field.NewTime(table, "add_date")

	o.fillFieldMap()

	return o
}

func (o *openiiotConfig) WithContext(ctx context.Context) *openiiotConfigDo {
	return o.openiiotConfigDo.WithContext(ctx)
}

func (o openiiotConfig) TableName() string { return o.openiiotConfigDo.TableName() }

func (o openiiotConfig) Alias() string { return o.openiiotConfigDo.Alias() }

func (o openiiotConfig) Columns(cols ...field.Expr) gen.Columns {
	return o.openiiotConfigDo.Columns(cols...)
}

func (o *openiiotConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *openiiotConfig) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 4)
	o.fieldMap["pkey"] = o.Pkey
	o.fieldMap["name"] = o.Name
	o.fieldMap["value"] = o.Value
	o.fieldMap["add_date"] = o.AddDate
}

func (o openiiotConfig) clone(db *gorm.DB) openiiotConfig {
	o.openiiotConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o openiiotConfig) replaceDB(db *gorm.DB) openiiotConfig {
	o.openiiotConfigDo.ReplaceDB(db)
	return o
}

type openiiotConfigDo struct{ gen.DO }

func (o openiiotConfigDo) Debug() *openiiotConfigDo {
	return o.withDO(o.DO.Debug())
}

func (o openiiotConfigDo) WithContext(ctx context.Context) *openiiotConfigDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o openiiotConfigDo) ReadDB() *openiiotConfigDo {
	return o.Clauses(dbresolver.Read)
}

func (o openiiotConfigDo) WriteDB() *openiiotConfigDo {
	return o.Clauses(dbresolver.Write)
}

func (o openiiotConfigDo) Session(config *gorm.Session) *openiiotConfigDo {
	return o.withDO(o.DO.Session(config))
}

func (o openiiotConfigDo) Clauses(conds ...clause.Expression) *openiiotConfigDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o openiiotConfigDo) Returning(value interface{}, columns ...string) *openiiotConfigDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o openiiotConfigDo) Not(conds ...gen.Condition) *openiiotConfigDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o openiiotConfigDo) Or(conds ...gen.Condition) *openiiotConfigDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o openiiotConfigDo) Select(conds ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o openiiotConfigDo) Where(conds ...gen.Condition) *openiiotConfigDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o openiiotConfigDo) Order(conds ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o openiiotConfigDo) Distinct(cols ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o openiiotConfigDo) Omit(cols ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o openiiotConfigDo) Join(table schema.Tabler, on ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o openiiotConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o openiiotConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o openiiotConfigDo) Group(cols ...field.Expr) *openiiotConfigDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o openiiotConfigDo) Having(conds ...gen.Condition) *openiiotConfigDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o openiiotConfigDo) Limit(limit int) *openiiotConfigDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o openiiotConfigDo) Offset(offset int) *openiiotConfigDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o openiiotConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *openiiotConfigDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o openiiotConfigDo) Unscoped() *openiiotConfigDo {
	return o.withDO(o.DO.Unscoped())
}

func (o openiiotConfigDo) Create(values ...*model_openiiot.OpeniiotConfig) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o openiiotConfigDo) CreateInBatches(values []*model_openiiot.OpeniiotConfig, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o openiiotConfigDo) Save(values ...*model_openiiot.OpeniiotConfig) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o openiiotConfigDo) First() (*model_openiiot.OpeniiotConfig, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.OpeniiotConfig), nil
	}
}

func (o openiiotConfigDo) Take() (*model_openiiot.OpeniiotConfig, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.OpeniiotConfig), nil
	}
}

func (o openiiotConfigDo) Last() (*model_openiiot.OpeniiotConfig, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.OpeniiotConfig), nil
	}
}

func (o openiiotConfigDo) Find() ([]*model_openiiot.OpeniiotConfig, error) {
	result, err := o.DO.Find()
	return result.([]*model_openiiot.OpeniiotConfig), err
}

func (o openiiotConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.OpeniiotConfig, err error) {
	buf := make([]*model_openiiot.OpeniiotConfig, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o openiiotConfigDo) FindInBatches(result *[]*model_openiiot.OpeniiotConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o openiiotConfigDo) Attrs(attrs ...field.AssignExpr) *openiiotConfigDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o openiiotConfigDo) Assign(attrs ...field.AssignExpr) *openiiotConfigDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o openiiotConfigDo) Joins(fields ...field.RelationField) *openiiotConfigDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o openiiotConfigDo) Preload(fields ...field.RelationField) *openiiotConfigDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o openiiotConfigDo) FirstOrInit() (*model_openiiot.OpeniiotConfig, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.OpeniiotConfig), nil
	}
}

func (o openiiotConfigDo) FirstOrCreate() (*model_openiiot.OpeniiotConfig, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.OpeniiotConfig), nil
	}
}

func (o openiiotConfigDo) FindByPage(offset int, limit int) (result []*model_openiiot.OpeniiotConfig, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o openiiotConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o openiiotConfigDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o openiiotConfigDo) Delete(models ...*model_openiiot.OpeniiotConfig) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *openiiotConfigDo) withDO(do gen.Dao) *openiiotConfigDo {
	o.DO = *do.(*gen.DO)
	return o
}

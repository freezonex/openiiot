// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query_openiiot

import (
	"context"
	"freezonex/openiiot/biz/dal/model_openiiot"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newWmsOutbound(db *gorm.DB, opts ...gen.DOOption) wmsOutbound {
	_wmsOutbound := wmsOutbound{}

	_wmsOutbound.wmsOutboundDo.UseDB(db, opts...)
	_wmsOutbound.wmsOutboundDo.UseModel(&model_openiiot.WmsOutbound{})

	tableName := _wmsOutbound.wmsOutboundDo.TableName()
	_wmsOutbound.ALL = field.NewAsterisk(tableName)
	_wmsOutbound.ID = field.NewInt64(tableName, "id")
	_wmsOutbound.RefID = field.NewString(tableName, "ref_id")
	_wmsOutbound.Type = field.NewString(tableName, "type")
	_wmsOutbound.Source = field.NewString(tableName, "source")
	_wmsOutbound.Note = field.NewString(tableName, "note")
	_wmsOutbound.Operator = field.NewString(tableName, "operator")
	_wmsOutbound.Status = field.NewString(tableName, "status")
	_wmsOutbound.UpdateTime = field.NewTime(tableName, "update_time")
	_wmsOutbound.CreateTime = field.NewTime(tableName, "create_time")

	_wmsOutbound.fillFieldMap()

	return _wmsOutbound
}

type wmsOutbound struct {
	wmsOutboundDo wmsOutboundDo

	ALL        field.Asterisk
	ID         field.Int64
	RefID      field.String
	Type       field.String
	Source     field.String
	Note       field.String
	Operator   field.String
	Status     field.String
	UpdateTime field.Time
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (w wmsOutbound) Table(newTableName string) *wmsOutbound {
	w.wmsOutboundDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsOutbound) As(alias string) *wmsOutbound {
	w.wmsOutboundDo.DO = *(w.wmsOutboundDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsOutbound) updateTableName(table string) *wmsOutbound {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.RefID = field.NewString(table, "ref_id")
	w.Type = field.NewString(table, "type")
	w.Source = field.NewString(table, "source")
	w.Note = field.NewString(table, "note")
	w.Operator = field.NewString(table, "operator")
	w.Status = field.NewString(table, "status")
	w.UpdateTime = field.NewTime(table, "update_time")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *wmsOutbound) WithContext(ctx context.Context) *wmsOutboundDo {
	return w.wmsOutboundDo.WithContext(ctx)
}

func (w wmsOutbound) TableName() string { return w.wmsOutboundDo.TableName() }

func (w wmsOutbound) Alias() string { return w.wmsOutboundDo.Alias() }

func (w wmsOutbound) Columns(cols ...field.Expr) gen.Columns { return w.wmsOutboundDo.Columns(cols...) }

func (w *wmsOutbound) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsOutbound) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["ref_id"] = w.RefID
	w.fieldMap["type"] = w.Type
	w.fieldMap["source"] = w.Source
	w.fieldMap["note"] = w.Note
	w.fieldMap["operator"] = w.Operator
	w.fieldMap["status"] = w.Status
	w.fieldMap["update_time"] = w.UpdateTime
	w.fieldMap["create_time"] = w.CreateTime
}

func (w wmsOutbound) clone(db *gorm.DB) wmsOutbound {
	w.wmsOutboundDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsOutbound) replaceDB(db *gorm.DB) wmsOutbound {
	w.wmsOutboundDo.ReplaceDB(db)
	return w
}

type wmsOutboundDo struct{ gen.DO }

func (w wmsOutboundDo) Debug() *wmsOutboundDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsOutboundDo) WithContext(ctx context.Context) *wmsOutboundDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsOutboundDo) ReadDB() *wmsOutboundDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsOutboundDo) WriteDB() *wmsOutboundDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsOutboundDo) Session(config *gorm.Session) *wmsOutboundDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsOutboundDo) Clauses(conds ...clause.Expression) *wmsOutboundDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsOutboundDo) Returning(value interface{}, columns ...string) *wmsOutboundDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsOutboundDo) Not(conds ...gen.Condition) *wmsOutboundDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsOutboundDo) Or(conds ...gen.Condition) *wmsOutboundDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsOutboundDo) Select(conds ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsOutboundDo) Where(conds ...gen.Condition) *wmsOutboundDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsOutboundDo) Order(conds ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsOutboundDo) Distinct(cols ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsOutboundDo) Omit(cols ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsOutboundDo) Join(table schema.Tabler, on ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsOutboundDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsOutboundDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsOutboundDo) Group(cols ...field.Expr) *wmsOutboundDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsOutboundDo) Having(conds ...gen.Condition) *wmsOutboundDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsOutboundDo) Limit(limit int) *wmsOutboundDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsOutboundDo) Offset(offset int) *wmsOutboundDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsOutboundDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsOutboundDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsOutboundDo) Unscoped() *wmsOutboundDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsOutboundDo) Create(values ...*model_openiiot.WmsOutbound) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsOutboundDo) CreateInBatches(values []*model_openiiot.WmsOutbound, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsOutboundDo) Save(values ...*model_openiiot.WmsOutbound) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsOutboundDo) First() (*model_openiiot.WmsOutbound, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutbound), nil
	}
}

func (w wmsOutboundDo) Take() (*model_openiiot.WmsOutbound, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutbound), nil
	}
}

func (w wmsOutboundDo) Last() (*model_openiiot.WmsOutbound, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutbound), nil
	}
}

func (w wmsOutboundDo) Find() ([]*model_openiiot.WmsOutbound, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsOutbound), err
}

func (w wmsOutboundDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsOutbound, err error) {
	buf := make([]*model_openiiot.WmsOutbound, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsOutboundDo) FindInBatches(result *[]*model_openiiot.WmsOutbound, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsOutboundDo) Attrs(attrs ...field.AssignExpr) *wmsOutboundDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsOutboundDo) Assign(attrs ...field.AssignExpr) *wmsOutboundDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsOutboundDo) Joins(fields ...field.RelationField) *wmsOutboundDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsOutboundDo) Preload(fields ...field.RelationField) *wmsOutboundDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsOutboundDo) FirstOrInit() (*model_openiiot.WmsOutbound, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutbound), nil
	}
}

func (w wmsOutboundDo) FirstOrCreate() (*model_openiiot.WmsOutbound, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutbound), nil
	}
}

func (w wmsOutboundDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsOutbound, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wmsOutboundDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsOutboundDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsOutboundDo) Delete(models ...*model_openiiot.WmsOutbound) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsOutboundDo) withDO(do gen.Dao) *wmsOutboundDo {
	w.DO = *do.(*gen.DO)
	return w
}

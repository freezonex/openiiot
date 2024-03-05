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

func newWmsStocktaking(db *gorm.DB, opts ...gen.DOOption) wmsStocktaking {
	_wmsStocktaking := wmsStocktaking{}

	_wmsStocktaking.wmsStocktakingDo.UseDB(db, opts...)
	_wmsStocktaking.wmsStocktakingDo.UseModel(&model_openiiot.WmsStocktaking{})

	tableName := _wmsStocktaking.wmsStocktakingDo.TableName()
	_wmsStocktaking.ALL = field.NewAsterisk(tableName)
	_wmsStocktaking.ID = field.NewInt64(tableName, "id")
	_wmsStocktaking.RefID = field.NewString(tableName, "ref_id")
	_wmsStocktaking.Type = field.NewString(tableName, "type")
	_wmsStocktaking.StorageLocation = field.NewInt64(tableName, "storage_location")
	_wmsStocktaking.Operator = field.NewString(tableName, "operator")
	_wmsStocktaking.Result = field.NewString(tableName, "result")
	_wmsStocktaking.UpdateTime = field.NewTime(tableName, "update_time")
	_wmsStocktaking.CreateTime = field.NewTime(tableName, "create_time")

	_wmsStocktaking.fillFieldMap()

	return _wmsStocktaking
}

type wmsStocktaking struct {
	wmsStocktakingDo wmsStocktakingDo

	ALL             field.Asterisk
	ID              field.Int64
	RefID           field.String
	Type            field.String
	StorageLocation field.Int64
	Operator        field.String
	Result          field.String
	UpdateTime      field.Time
	CreateTime      field.Time

	fieldMap map[string]field.Expr
}

func (w wmsStocktaking) Table(newTableName string) *wmsStocktaking {
	w.wmsStocktakingDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsStocktaking) As(alias string) *wmsStocktaking {
	w.wmsStocktakingDo.DO = *(w.wmsStocktakingDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsStocktaking) updateTableName(table string) *wmsStocktaking {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.RefID = field.NewString(table, "ref_id")
	w.Type = field.NewString(table, "type")
	w.StorageLocation = field.NewInt64(table, "storage_location")
	w.Operator = field.NewString(table, "operator")
	w.Result = field.NewString(table, "result")
	w.UpdateTime = field.NewTime(table, "update_time")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *wmsStocktaking) WithContext(ctx context.Context) *wmsStocktakingDo {
	return w.wmsStocktakingDo.WithContext(ctx)
}

func (w wmsStocktaking) TableName() string { return w.wmsStocktakingDo.TableName() }

func (w wmsStocktaking) Alias() string { return w.wmsStocktakingDo.Alias() }

func (w wmsStocktaking) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsStocktakingDo.Columns(cols...)
}

func (w *wmsStocktaking) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsStocktaking) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["ref_id"] = w.RefID
	w.fieldMap["type"] = w.Type
	w.fieldMap["storage_location"] = w.StorageLocation
	w.fieldMap["operator"] = w.Operator
	w.fieldMap["result"] = w.Result
	w.fieldMap["update_time"] = w.UpdateTime
	w.fieldMap["create_time"] = w.CreateTime
}

func (w wmsStocktaking) clone(db *gorm.DB) wmsStocktaking {
	w.wmsStocktakingDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsStocktaking) replaceDB(db *gorm.DB) wmsStocktaking {
	w.wmsStocktakingDo.ReplaceDB(db)
	return w
}

type wmsStocktakingDo struct{ gen.DO }

func (w wmsStocktakingDo) Debug() *wmsStocktakingDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsStocktakingDo) WithContext(ctx context.Context) *wmsStocktakingDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsStocktakingDo) ReadDB() *wmsStocktakingDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsStocktakingDo) WriteDB() *wmsStocktakingDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsStocktakingDo) Session(config *gorm.Session) *wmsStocktakingDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsStocktakingDo) Clauses(conds ...clause.Expression) *wmsStocktakingDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsStocktakingDo) Returning(value interface{}, columns ...string) *wmsStocktakingDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsStocktakingDo) Not(conds ...gen.Condition) *wmsStocktakingDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsStocktakingDo) Or(conds ...gen.Condition) *wmsStocktakingDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsStocktakingDo) Select(conds ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsStocktakingDo) Where(conds ...gen.Condition) *wmsStocktakingDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsStocktakingDo) Order(conds ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsStocktakingDo) Distinct(cols ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsStocktakingDo) Omit(cols ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsStocktakingDo) Join(table schema.Tabler, on ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsStocktakingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsStocktakingDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsStocktakingDo) Group(cols ...field.Expr) *wmsStocktakingDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsStocktakingDo) Having(conds ...gen.Condition) *wmsStocktakingDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsStocktakingDo) Limit(limit int) *wmsStocktakingDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsStocktakingDo) Offset(offset int) *wmsStocktakingDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsStocktakingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsStocktakingDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsStocktakingDo) Unscoped() *wmsStocktakingDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsStocktakingDo) Create(values ...*model_openiiot.WmsStocktaking) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsStocktakingDo) CreateInBatches(values []*model_openiiot.WmsStocktaking, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsStocktakingDo) Save(values ...*model_openiiot.WmsStocktaking) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsStocktakingDo) First() (*model_openiiot.WmsStocktaking, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktaking), nil
	}
}

func (w wmsStocktakingDo) Take() (*model_openiiot.WmsStocktaking, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktaking), nil
	}
}

func (w wmsStocktakingDo) Last() (*model_openiiot.WmsStocktaking, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktaking), nil
	}
}

func (w wmsStocktakingDo) Find() ([]*model_openiiot.WmsStocktaking, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsStocktaking), err
}

func (w wmsStocktakingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsStocktaking, err error) {
	buf := make([]*model_openiiot.WmsStocktaking, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsStocktakingDo) FindInBatches(result *[]*model_openiiot.WmsStocktaking, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsStocktakingDo) Attrs(attrs ...field.AssignExpr) *wmsStocktakingDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsStocktakingDo) Assign(attrs ...field.AssignExpr) *wmsStocktakingDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsStocktakingDo) Joins(fields ...field.RelationField) *wmsStocktakingDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsStocktakingDo) Preload(fields ...field.RelationField) *wmsStocktakingDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsStocktakingDo) FirstOrInit() (*model_openiiot.WmsStocktaking, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktaking), nil
	}
}

func (w wmsStocktakingDo) FirstOrCreate() (*model_openiiot.WmsStocktaking, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktaking), nil
	}
}

func (w wmsStocktakingDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsStocktaking, count int64, err error) {
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

func (w wmsStocktakingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsStocktakingDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsStocktakingDo) Delete(models ...*model_openiiot.WmsStocktaking) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsStocktakingDo) withDO(do gen.Dao) *wmsStocktakingDo {
	w.DO = *do.(*gen.DO)
	return w
}

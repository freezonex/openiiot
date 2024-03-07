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

func newWmsOutboundRecord(db *gorm.DB, opts ...gen.DOOption) wmsOutboundRecord {
	_wmsOutboundRecord := wmsOutboundRecord{}

	_wmsOutboundRecord.wmsOutboundRecordDo.UseDB(db, opts...)
	_wmsOutboundRecord.wmsOutboundRecordDo.UseModel(&model_openiiot.WmsOutboundRecord{})

	tableName := _wmsOutboundRecord.wmsOutboundRecordDo.TableName()
	_wmsOutboundRecord.ALL = field.NewAsterisk(tableName)
	_wmsOutboundRecord.ID = field.NewInt64(tableName, "id")
	_wmsOutboundRecord.OutboundID = field.NewInt64(tableName, "outbound_id")
	_wmsOutboundRecord.StockLocationID = field.NewInt64(tableName, "stock_location_id")
	_wmsOutboundRecord.MaterialID = field.NewInt64(tableName, "material_id")
	_wmsOutboundRecord.Quantity = field.NewInt32(tableName, "quantity")

	_wmsOutboundRecord.fillFieldMap()

	return _wmsOutboundRecord
}

type wmsOutboundRecord struct {
	wmsOutboundRecordDo wmsOutboundRecordDo

	ALL             field.Asterisk
	ID              field.Int64
	OutboundID      field.Int64
	StockLocationID field.Int64
	MaterialID      field.Int64
	Quantity        field.Int32

	fieldMap map[string]field.Expr
}

func (w wmsOutboundRecord) Table(newTableName string) *wmsOutboundRecord {
	w.wmsOutboundRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsOutboundRecord) As(alias string) *wmsOutboundRecord {
	w.wmsOutboundRecordDo.DO = *(w.wmsOutboundRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsOutboundRecord) updateTableName(table string) *wmsOutboundRecord {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.OutboundID = field.NewInt64(table, "outbound_id")
	w.StockLocationID = field.NewInt64(table, "stock_location_id")
	w.MaterialID = field.NewInt64(table, "material_id")
	w.Quantity = field.NewInt32(table, "quantity")

	w.fillFieldMap()

	return w
}

func (w *wmsOutboundRecord) WithContext(ctx context.Context) *wmsOutboundRecordDo {
	return w.wmsOutboundRecordDo.WithContext(ctx)
}

func (w wmsOutboundRecord) TableName() string { return w.wmsOutboundRecordDo.TableName() }

func (w wmsOutboundRecord) Alias() string { return w.wmsOutboundRecordDo.Alias() }

func (w wmsOutboundRecord) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsOutboundRecordDo.Columns(cols...)
}

func (w *wmsOutboundRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsOutboundRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["id"] = w.ID
	w.fieldMap["outbound_id"] = w.OutboundID
	w.fieldMap["stock_location_id"] = w.StockLocationID
	w.fieldMap["material_id"] = w.MaterialID
	w.fieldMap["quantity"] = w.Quantity
}

func (w wmsOutboundRecord) clone(db *gorm.DB) wmsOutboundRecord {
	w.wmsOutboundRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsOutboundRecord) replaceDB(db *gorm.DB) wmsOutboundRecord {
	w.wmsOutboundRecordDo.ReplaceDB(db)
	return w
}

type wmsOutboundRecordDo struct{ gen.DO }

func (w wmsOutboundRecordDo) Debug() *wmsOutboundRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsOutboundRecordDo) WithContext(ctx context.Context) *wmsOutboundRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsOutboundRecordDo) ReadDB() *wmsOutboundRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsOutboundRecordDo) WriteDB() *wmsOutboundRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsOutboundRecordDo) Session(config *gorm.Session) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsOutboundRecordDo) Clauses(conds ...clause.Expression) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsOutboundRecordDo) Returning(value interface{}, columns ...string) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsOutboundRecordDo) Not(conds ...gen.Condition) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsOutboundRecordDo) Or(conds ...gen.Condition) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsOutboundRecordDo) Select(conds ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsOutboundRecordDo) Where(conds ...gen.Condition) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsOutboundRecordDo) Order(conds ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsOutboundRecordDo) Distinct(cols ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsOutboundRecordDo) Omit(cols ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsOutboundRecordDo) Join(table schema.Tabler, on ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsOutboundRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsOutboundRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsOutboundRecordDo) Group(cols ...field.Expr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsOutboundRecordDo) Having(conds ...gen.Condition) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsOutboundRecordDo) Limit(limit int) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsOutboundRecordDo) Offset(offset int) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsOutboundRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsOutboundRecordDo) Unscoped() *wmsOutboundRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsOutboundRecordDo) Create(values ...*model_openiiot.WmsOutboundRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsOutboundRecordDo) CreateInBatches(values []*model_openiiot.WmsOutboundRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsOutboundRecordDo) Save(values ...*model_openiiot.WmsOutboundRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsOutboundRecordDo) First() (*model_openiiot.WmsOutboundRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutboundRecord), nil
	}
}

func (w wmsOutboundRecordDo) Take() (*model_openiiot.WmsOutboundRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutboundRecord), nil
	}
}

func (w wmsOutboundRecordDo) Last() (*model_openiiot.WmsOutboundRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutboundRecord), nil
	}
}

func (w wmsOutboundRecordDo) Find() ([]*model_openiiot.WmsOutboundRecord, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsOutboundRecord), err
}

func (w wmsOutboundRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsOutboundRecord, err error) {
	buf := make([]*model_openiiot.WmsOutboundRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsOutboundRecordDo) FindInBatches(result *[]*model_openiiot.WmsOutboundRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsOutboundRecordDo) Attrs(attrs ...field.AssignExpr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsOutboundRecordDo) Assign(attrs ...field.AssignExpr) *wmsOutboundRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsOutboundRecordDo) Joins(fields ...field.RelationField) *wmsOutboundRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsOutboundRecordDo) Preload(fields ...field.RelationField) *wmsOutboundRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsOutboundRecordDo) FirstOrInit() (*model_openiiot.WmsOutboundRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutboundRecord), nil
	}
}

func (w wmsOutboundRecordDo) FirstOrCreate() (*model_openiiot.WmsOutboundRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsOutboundRecord), nil
	}
}

func (w wmsOutboundRecordDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsOutboundRecord, count int64, err error) {
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

func (w wmsOutboundRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsOutboundRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsOutboundRecordDo) Delete(models ...*model_openiiot.WmsOutboundRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsOutboundRecordDo) withDO(do gen.Dao) *wmsOutboundRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}

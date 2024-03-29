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

func newWmsInboundRecord(db *gorm.DB, opts ...gen.DOOption) wmsInboundRecord {
	_wmsInboundRecord := wmsInboundRecord{}

	_wmsInboundRecord.wmsInboundRecordDo.UseDB(db, opts...)
	_wmsInboundRecord.wmsInboundRecordDo.UseModel(&model_openiiot.WmsInboundRecord{})

	tableName := _wmsInboundRecord.wmsInboundRecordDo.TableName()
	_wmsInboundRecord.ALL = field.NewAsterisk(tableName)
	_wmsInboundRecord.ID = field.NewInt64(tableName, "id")
	_wmsInboundRecord.InboundID = field.NewInt64(tableName, "inbound_id")
	_wmsInboundRecord.StockLocationID = field.NewInt64(tableName, "stock_location_id")
	_wmsInboundRecord.MaterialID = field.NewInt64(tableName, "material_id")
	_wmsInboundRecord.Quantity = field.NewInt32(tableName, "quantity")

	_wmsInboundRecord.fillFieldMap()

	return _wmsInboundRecord
}

type wmsInboundRecord struct {
	wmsInboundRecordDo wmsInboundRecordDo

	ALL             field.Asterisk
	ID              field.Int64
	InboundID       field.Int64
	StockLocationID field.Int64
	MaterialID      field.Int64
	Quantity        field.Int32

	fieldMap map[string]field.Expr
}

func (w wmsInboundRecord) Table(newTableName string) *wmsInboundRecord {
	w.wmsInboundRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsInboundRecord) As(alias string) *wmsInboundRecord {
	w.wmsInboundRecordDo.DO = *(w.wmsInboundRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsInboundRecord) updateTableName(table string) *wmsInboundRecord {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.InboundID = field.NewInt64(table, "inbound_id")
	w.StockLocationID = field.NewInt64(table, "stock_location_id")
	w.MaterialID = field.NewInt64(table, "material_id")
	w.Quantity = field.NewInt32(table, "quantity")

	w.fillFieldMap()

	return w
}

func (w *wmsInboundRecord) WithContext(ctx context.Context) *wmsInboundRecordDo {
	return w.wmsInboundRecordDo.WithContext(ctx)
}

func (w wmsInboundRecord) TableName() string { return w.wmsInboundRecordDo.TableName() }

func (w wmsInboundRecord) Alias() string { return w.wmsInboundRecordDo.Alias() }

func (w wmsInboundRecord) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsInboundRecordDo.Columns(cols...)
}

func (w *wmsInboundRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsInboundRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["id"] = w.ID
	w.fieldMap["inbound_id"] = w.InboundID
	w.fieldMap["stock_location_id"] = w.StockLocationID
	w.fieldMap["material_id"] = w.MaterialID
	w.fieldMap["quantity"] = w.Quantity
}

func (w wmsInboundRecord) clone(db *gorm.DB) wmsInboundRecord {
	w.wmsInboundRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsInboundRecord) replaceDB(db *gorm.DB) wmsInboundRecord {
	w.wmsInboundRecordDo.ReplaceDB(db)
	return w
}

type wmsInboundRecordDo struct{ gen.DO }

func (w wmsInboundRecordDo) Debug() *wmsInboundRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsInboundRecordDo) WithContext(ctx context.Context) *wmsInboundRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsInboundRecordDo) ReadDB() *wmsInboundRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsInboundRecordDo) WriteDB() *wmsInboundRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsInboundRecordDo) Session(config *gorm.Session) *wmsInboundRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsInboundRecordDo) Clauses(conds ...clause.Expression) *wmsInboundRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsInboundRecordDo) Returning(value interface{}, columns ...string) *wmsInboundRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsInboundRecordDo) Not(conds ...gen.Condition) *wmsInboundRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsInboundRecordDo) Or(conds ...gen.Condition) *wmsInboundRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsInboundRecordDo) Select(conds ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsInboundRecordDo) Where(conds ...gen.Condition) *wmsInboundRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsInboundRecordDo) Order(conds ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsInboundRecordDo) Distinct(cols ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsInboundRecordDo) Omit(cols ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsInboundRecordDo) Join(table schema.Tabler, on ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsInboundRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsInboundRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsInboundRecordDo) Group(cols ...field.Expr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsInboundRecordDo) Having(conds ...gen.Condition) *wmsInboundRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsInboundRecordDo) Limit(limit int) *wmsInboundRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsInboundRecordDo) Offset(offset int) *wmsInboundRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsInboundRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsInboundRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsInboundRecordDo) Unscoped() *wmsInboundRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsInboundRecordDo) Create(values ...*model_openiiot.WmsInboundRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsInboundRecordDo) CreateInBatches(values []*model_openiiot.WmsInboundRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsInboundRecordDo) Save(values ...*model_openiiot.WmsInboundRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsInboundRecordDo) First() (*model_openiiot.WmsInboundRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsInboundRecord), nil
	}
}

func (w wmsInboundRecordDo) Take() (*model_openiiot.WmsInboundRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsInboundRecord), nil
	}
}

func (w wmsInboundRecordDo) Last() (*model_openiiot.WmsInboundRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsInboundRecord), nil
	}
}

func (w wmsInboundRecordDo) Find() ([]*model_openiiot.WmsInboundRecord, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsInboundRecord), err
}

func (w wmsInboundRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsInboundRecord, err error) {
	buf := make([]*model_openiiot.WmsInboundRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsInboundRecordDo) FindInBatches(result *[]*model_openiiot.WmsInboundRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsInboundRecordDo) Attrs(attrs ...field.AssignExpr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsInboundRecordDo) Assign(attrs ...field.AssignExpr) *wmsInboundRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsInboundRecordDo) Joins(fields ...field.RelationField) *wmsInboundRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsInboundRecordDo) Preload(fields ...field.RelationField) *wmsInboundRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsInboundRecordDo) FirstOrInit() (*model_openiiot.WmsInboundRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsInboundRecord), nil
	}
}

func (w wmsInboundRecordDo) FirstOrCreate() (*model_openiiot.WmsInboundRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsInboundRecord), nil
	}
}

func (w wmsInboundRecordDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsInboundRecord, count int64, err error) {
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

func (w wmsInboundRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsInboundRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsInboundRecordDo) Delete(models ...*model_openiiot.WmsInboundRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsInboundRecordDo) withDO(do gen.Dao) *wmsInboundRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}

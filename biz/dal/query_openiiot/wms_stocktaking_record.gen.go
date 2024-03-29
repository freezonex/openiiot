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

func newWmsStocktakingRecord(db *gorm.DB, opts ...gen.DOOption) wmsStocktakingRecord {
	_wmsStocktakingRecord := wmsStocktakingRecord{}

	_wmsStocktakingRecord.wmsStocktakingRecordDo.UseDB(db, opts...)
	_wmsStocktakingRecord.wmsStocktakingRecordDo.UseModel(&model_openiiot.WmsStocktakingRecord{})

	tableName := _wmsStocktakingRecord.wmsStocktakingRecordDo.TableName()
	_wmsStocktakingRecord.ALL = field.NewAsterisk(tableName)
	_wmsStocktakingRecord.ID = field.NewInt64(tableName, "id")
	_wmsStocktakingRecord.StocktakingID = field.NewInt64(tableName, "stocktaking_id")
	_wmsStocktakingRecord.StockLocationID = field.NewInt64(tableName, "stock_location_id")
	_wmsStocktakingRecord.MaterialID = field.NewInt64(tableName, "material_id")
	_wmsStocktakingRecord.Quantity = field.NewInt32(tableName, "quantity")
	_wmsStocktakingRecord.StockQuantity = field.NewInt32(tableName, "stock_quantity")
	_wmsStocktakingRecord.Discrepancy = field.NewInt32(tableName, "discrepancy")

	_wmsStocktakingRecord.fillFieldMap()

	return _wmsStocktakingRecord
}

type wmsStocktakingRecord struct {
	wmsStocktakingRecordDo wmsStocktakingRecordDo

	ALL             field.Asterisk
	ID              field.Int64
	StocktakingID   field.Int64
	StockLocationID field.Int64
	MaterialID      field.Int64
	Quantity        field.Int32
	StockQuantity   field.Int32
	Discrepancy     field.Int32

	fieldMap map[string]field.Expr
}

func (w wmsStocktakingRecord) Table(newTableName string) *wmsStocktakingRecord {
	w.wmsStocktakingRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsStocktakingRecord) As(alias string) *wmsStocktakingRecord {
	w.wmsStocktakingRecordDo.DO = *(w.wmsStocktakingRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsStocktakingRecord) updateTableName(table string) *wmsStocktakingRecord {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.StocktakingID = field.NewInt64(table, "stocktaking_id")
	w.StockLocationID = field.NewInt64(table, "stock_location_id")
	w.MaterialID = field.NewInt64(table, "material_id")
	w.Quantity = field.NewInt32(table, "quantity")
	w.StockQuantity = field.NewInt32(table, "stock_quantity")
	w.Discrepancy = field.NewInt32(table, "discrepancy")

	w.fillFieldMap()

	return w
}

func (w *wmsStocktakingRecord) WithContext(ctx context.Context) *wmsStocktakingRecordDo {
	return w.wmsStocktakingRecordDo.WithContext(ctx)
}

func (w wmsStocktakingRecord) TableName() string { return w.wmsStocktakingRecordDo.TableName() }

func (w wmsStocktakingRecord) Alias() string { return w.wmsStocktakingRecordDo.Alias() }

func (w wmsStocktakingRecord) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsStocktakingRecordDo.Columns(cols...)
}

func (w *wmsStocktakingRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsStocktakingRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["stocktaking_id"] = w.StocktakingID
	w.fieldMap["stock_location_id"] = w.StockLocationID
	w.fieldMap["material_id"] = w.MaterialID
	w.fieldMap["quantity"] = w.Quantity
	w.fieldMap["stock_quantity"] = w.StockQuantity
	w.fieldMap["discrepancy"] = w.Discrepancy
}

func (w wmsStocktakingRecord) clone(db *gorm.DB) wmsStocktakingRecord {
	w.wmsStocktakingRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsStocktakingRecord) replaceDB(db *gorm.DB) wmsStocktakingRecord {
	w.wmsStocktakingRecordDo.ReplaceDB(db)
	return w
}

type wmsStocktakingRecordDo struct{ gen.DO }

func (w wmsStocktakingRecordDo) Debug() *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsStocktakingRecordDo) WithContext(ctx context.Context) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsStocktakingRecordDo) ReadDB() *wmsStocktakingRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsStocktakingRecordDo) WriteDB() *wmsStocktakingRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsStocktakingRecordDo) Session(config *gorm.Session) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsStocktakingRecordDo) Clauses(conds ...clause.Expression) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsStocktakingRecordDo) Returning(value interface{}, columns ...string) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsStocktakingRecordDo) Not(conds ...gen.Condition) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsStocktakingRecordDo) Or(conds ...gen.Condition) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsStocktakingRecordDo) Select(conds ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsStocktakingRecordDo) Where(conds ...gen.Condition) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsStocktakingRecordDo) Order(conds ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsStocktakingRecordDo) Distinct(cols ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsStocktakingRecordDo) Omit(cols ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsStocktakingRecordDo) Join(table schema.Tabler, on ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsStocktakingRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsStocktakingRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsStocktakingRecordDo) Group(cols ...field.Expr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsStocktakingRecordDo) Having(conds ...gen.Condition) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsStocktakingRecordDo) Limit(limit int) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsStocktakingRecordDo) Offset(offset int) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsStocktakingRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsStocktakingRecordDo) Unscoped() *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsStocktakingRecordDo) Create(values ...*model_openiiot.WmsStocktakingRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsStocktakingRecordDo) CreateInBatches(values []*model_openiiot.WmsStocktakingRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsStocktakingRecordDo) Save(values ...*model_openiiot.WmsStocktakingRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsStocktakingRecordDo) First() (*model_openiiot.WmsStocktakingRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktakingRecord), nil
	}
}

func (w wmsStocktakingRecordDo) Take() (*model_openiiot.WmsStocktakingRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktakingRecord), nil
	}
}

func (w wmsStocktakingRecordDo) Last() (*model_openiiot.WmsStocktakingRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktakingRecord), nil
	}
}

func (w wmsStocktakingRecordDo) Find() ([]*model_openiiot.WmsStocktakingRecord, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsStocktakingRecord), err
}

func (w wmsStocktakingRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsStocktakingRecord, err error) {
	buf := make([]*model_openiiot.WmsStocktakingRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsStocktakingRecordDo) FindInBatches(result *[]*model_openiiot.WmsStocktakingRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsStocktakingRecordDo) Attrs(attrs ...field.AssignExpr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsStocktakingRecordDo) Assign(attrs ...field.AssignExpr) *wmsStocktakingRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsStocktakingRecordDo) Joins(fields ...field.RelationField) *wmsStocktakingRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsStocktakingRecordDo) Preload(fields ...field.RelationField) *wmsStocktakingRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsStocktakingRecordDo) FirstOrInit() (*model_openiiot.WmsStocktakingRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktakingRecord), nil
	}
}

func (w wmsStocktakingRecordDo) FirstOrCreate() (*model_openiiot.WmsStocktakingRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStocktakingRecord), nil
	}
}

func (w wmsStocktakingRecordDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsStocktakingRecord, count int64, err error) {
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

func (w wmsStocktakingRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsStocktakingRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsStocktakingRecordDo) Delete(models ...*model_openiiot.WmsStocktakingRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsStocktakingRecordDo) withDO(do gen.Dao) *wmsStocktakingRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}

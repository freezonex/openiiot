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

func newWmsStorageLocation(db *gorm.DB, opts ...gen.DOOption) wmsStorageLocation {
	_wmsStorageLocation := wmsStorageLocation{}

	_wmsStorageLocation.wmsStorageLocationDo.UseDB(db, opts...)
	_wmsStorageLocation.wmsStorageLocationDo.UseModel(&model_openiiot.WmsStorageLocation{})

	tableName := _wmsStorageLocation.wmsStorageLocationDo.TableName()
	_wmsStorageLocation.ALL = field.NewAsterisk(tableName)
	_wmsStorageLocation.ID = field.NewInt64(tableName, "id")
	_wmsStorageLocation.WarehouseID = field.NewInt64(tableName, "warehouse_id")
	_wmsStorageLocation.Name = field.NewString(tableName, "name")
	_wmsStorageLocation.Occupied = field.NewBool(tableName, "occupied")
	_wmsStorageLocation.MaterialName = field.NewString(tableName, "material_name")
	_wmsStorageLocation.UpdateTime = field.NewTime(tableName, "update_time")
	_wmsStorageLocation.CreateTime = field.NewTime(tableName, "create_time")
	_wmsStorageLocation.MaterialQuantity = field.NewString(tableName, "material_quantity")

	_wmsStorageLocation.fillFieldMap()

	return _wmsStorageLocation
}

type wmsStorageLocation struct {
	wmsStorageLocationDo wmsStorageLocationDo

	ALL              field.Asterisk
	ID               field.Int64
	WarehouseID      field.Int64
	Name             field.String
	Occupied         field.Bool
	MaterialName     field.String
	UpdateTime       field.Time
	CreateTime       field.Time
	MaterialQuantity field.String

	fieldMap map[string]field.Expr
}

func (w wmsStorageLocation) Table(newTableName string) *wmsStorageLocation {
	w.wmsStorageLocationDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsStorageLocation) As(alias string) *wmsStorageLocation {
	w.wmsStorageLocationDo.DO = *(w.wmsStorageLocationDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsStorageLocation) updateTableName(table string) *wmsStorageLocation {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.WarehouseID = field.NewInt64(table, "warehouse_id")
	w.Name = field.NewString(table, "name")
	w.Occupied = field.NewBool(table, "occupied")
	w.MaterialName = field.NewString(table, "material_name")
	w.UpdateTime = field.NewTime(table, "update_time")
	w.CreateTime = field.NewTime(table, "create_time")
	w.MaterialQuantity = field.NewString(table, "material_quantity")

	w.fillFieldMap()

	return w
}

func (w *wmsStorageLocation) WithContext(ctx context.Context) *wmsStorageLocationDo {
	return w.wmsStorageLocationDo.WithContext(ctx)
}

func (w wmsStorageLocation) TableName() string { return w.wmsStorageLocationDo.TableName() }

func (w wmsStorageLocation) Alias() string { return w.wmsStorageLocationDo.Alias() }

func (w wmsStorageLocation) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsStorageLocationDo.Columns(cols...)
}

func (w *wmsStorageLocation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsStorageLocation) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["warehouse_id"] = w.WarehouseID
	w.fieldMap["name"] = w.Name
	w.fieldMap["occupied"] = w.Occupied
	w.fieldMap["material_name"] = w.MaterialName
	w.fieldMap["update_time"] = w.UpdateTime
	w.fieldMap["create_time"] = w.CreateTime
	w.fieldMap["material_quantity"] = w.MaterialQuantity
}

func (w wmsStorageLocation) clone(db *gorm.DB) wmsStorageLocation {
	w.wmsStorageLocationDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsStorageLocation) replaceDB(db *gorm.DB) wmsStorageLocation {
	w.wmsStorageLocationDo.ReplaceDB(db)
	return w
}

type wmsStorageLocationDo struct{ gen.DO }

func (w wmsStorageLocationDo) Debug() *wmsStorageLocationDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsStorageLocationDo) WithContext(ctx context.Context) *wmsStorageLocationDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsStorageLocationDo) ReadDB() *wmsStorageLocationDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsStorageLocationDo) WriteDB() *wmsStorageLocationDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsStorageLocationDo) Session(config *gorm.Session) *wmsStorageLocationDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsStorageLocationDo) Clauses(conds ...clause.Expression) *wmsStorageLocationDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsStorageLocationDo) Returning(value interface{}, columns ...string) *wmsStorageLocationDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsStorageLocationDo) Not(conds ...gen.Condition) *wmsStorageLocationDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsStorageLocationDo) Or(conds ...gen.Condition) *wmsStorageLocationDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsStorageLocationDo) Select(conds ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsStorageLocationDo) Where(conds ...gen.Condition) *wmsStorageLocationDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsStorageLocationDo) Order(conds ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsStorageLocationDo) Distinct(cols ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsStorageLocationDo) Omit(cols ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsStorageLocationDo) Join(table schema.Tabler, on ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsStorageLocationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsStorageLocationDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsStorageLocationDo) Group(cols ...field.Expr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsStorageLocationDo) Having(conds ...gen.Condition) *wmsStorageLocationDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsStorageLocationDo) Limit(limit int) *wmsStorageLocationDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsStorageLocationDo) Offset(offset int) *wmsStorageLocationDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsStorageLocationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsStorageLocationDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsStorageLocationDo) Unscoped() *wmsStorageLocationDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsStorageLocationDo) Create(values ...*model_openiiot.WmsStorageLocation) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsStorageLocationDo) CreateInBatches(values []*model_openiiot.WmsStorageLocation, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsStorageLocationDo) Save(values ...*model_openiiot.WmsStorageLocation) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsStorageLocationDo) First() (*model_openiiot.WmsStorageLocation, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocation), nil
	}
}

func (w wmsStorageLocationDo) Take() (*model_openiiot.WmsStorageLocation, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocation), nil
	}
}

func (w wmsStorageLocationDo) Last() (*model_openiiot.WmsStorageLocation, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocation), nil
	}
}

func (w wmsStorageLocationDo) Find() ([]*model_openiiot.WmsStorageLocation, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsStorageLocation), err
}

func (w wmsStorageLocationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsStorageLocation, err error) {
	buf := make([]*model_openiiot.WmsStorageLocation, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsStorageLocationDo) FindInBatches(result *[]*model_openiiot.WmsStorageLocation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsStorageLocationDo) Attrs(attrs ...field.AssignExpr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsStorageLocationDo) Assign(attrs ...field.AssignExpr) *wmsStorageLocationDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsStorageLocationDo) Joins(fields ...field.RelationField) *wmsStorageLocationDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsStorageLocationDo) Preload(fields ...field.RelationField) *wmsStorageLocationDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsStorageLocationDo) FirstOrInit() (*model_openiiot.WmsStorageLocation, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocation), nil
	}
}

func (w wmsStorageLocationDo) FirstOrCreate() (*model_openiiot.WmsStorageLocation, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocation), nil
	}
}

func (w wmsStorageLocationDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsStorageLocation, count int64, err error) {
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

func (w wmsStorageLocationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsStorageLocationDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsStorageLocationDo) Delete(models ...*model_openiiot.WmsStorageLocation) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsStorageLocationDo) withDO(do gen.Dao) *wmsStorageLocationDo {
	w.DO = *do.(*gen.DO)
	return w
}

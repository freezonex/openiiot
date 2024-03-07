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

func newWmsStorageLocationMaterial(db *gorm.DB, opts ...gen.DOOption) wmsStorageLocationMaterial {
	_wmsStorageLocationMaterial := wmsStorageLocationMaterial{}

	_wmsStorageLocationMaterial.wmsStorageLocationMaterialDo.UseDB(db, opts...)
	_wmsStorageLocationMaterial.wmsStorageLocationMaterialDo.UseModel(&model_openiiot.WmsStorageLocationMaterial{})

	tableName := _wmsStorageLocationMaterial.wmsStorageLocationMaterialDo.TableName()
	_wmsStorageLocationMaterial.ALL = field.NewAsterisk(tableName)
	_wmsStorageLocationMaterial.ID = field.NewInt64(tableName, "id")
	_wmsStorageLocationMaterial.StoreLocationID = field.NewInt64(tableName, "store_location_id")
	_wmsStorageLocationMaterial.MaterialID = field.NewInt64(tableName, "material_id")
	_wmsStorageLocationMaterial.Quantity = field.NewInt32(tableName, "quantity")

	_wmsStorageLocationMaterial.fillFieldMap()

	return _wmsStorageLocationMaterial
}

type wmsStorageLocationMaterial struct {
	wmsStorageLocationMaterialDo wmsStorageLocationMaterialDo

	ALL             field.Asterisk
	ID              field.Int64
	StoreLocationID field.Int64
	MaterialID      field.Int64
	Quantity        field.Int32

	fieldMap map[string]field.Expr
}

func (w wmsStorageLocationMaterial) Table(newTableName string) *wmsStorageLocationMaterial {
	w.wmsStorageLocationMaterialDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsStorageLocationMaterial) As(alias string) *wmsStorageLocationMaterial {
	w.wmsStorageLocationMaterialDo.DO = *(w.wmsStorageLocationMaterialDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsStorageLocationMaterial) updateTableName(table string) *wmsStorageLocationMaterial {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.StoreLocationID = field.NewInt64(table, "store_location_id")
	w.MaterialID = field.NewInt64(table, "material_id")
	w.Quantity = field.NewInt32(table, "quantity")

	w.fillFieldMap()

	return w
}

func (w *wmsStorageLocationMaterial) WithContext(ctx context.Context) *wmsStorageLocationMaterialDo {
	return w.wmsStorageLocationMaterialDo.WithContext(ctx)
}

func (w wmsStorageLocationMaterial) TableName() string {
	return w.wmsStorageLocationMaterialDo.TableName()
}

func (w wmsStorageLocationMaterial) Alias() string { return w.wmsStorageLocationMaterialDo.Alias() }

func (w wmsStorageLocationMaterial) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsStorageLocationMaterialDo.Columns(cols...)
}

func (w *wmsStorageLocationMaterial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsStorageLocationMaterial) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 4)
	w.fieldMap["id"] = w.ID
	w.fieldMap["store_location_id"] = w.StoreLocationID
	w.fieldMap["material_id"] = w.MaterialID
	w.fieldMap["quantity"] = w.Quantity
}

func (w wmsStorageLocationMaterial) clone(db *gorm.DB) wmsStorageLocationMaterial {
	w.wmsStorageLocationMaterialDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsStorageLocationMaterial) replaceDB(db *gorm.DB) wmsStorageLocationMaterial {
	w.wmsStorageLocationMaterialDo.ReplaceDB(db)
	return w
}

type wmsStorageLocationMaterialDo struct{ gen.DO }

func (w wmsStorageLocationMaterialDo) Debug() *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsStorageLocationMaterialDo) WithContext(ctx context.Context) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsStorageLocationMaterialDo) ReadDB() *wmsStorageLocationMaterialDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsStorageLocationMaterialDo) WriteDB() *wmsStorageLocationMaterialDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsStorageLocationMaterialDo) Session(config *gorm.Session) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsStorageLocationMaterialDo) Clauses(conds ...clause.Expression) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsStorageLocationMaterialDo) Returning(value interface{}, columns ...string) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsStorageLocationMaterialDo) Not(conds ...gen.Condition) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsStorageLocationMaterialDo) Or(conds ...gen.Condition) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsStorageLocationMaterialDo) Select(conds ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsStorageLocationMaterialDo) Where(conds ...gen.Condition) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsStorageLocationMaterialDo) Order(conds ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsStorageLocationMaterialDo) Distinct(cols ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsStorageLocationMaterialDo) Omit(cols ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsStorageLocationMaterialDo) Join(table schema.Tabler, on ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsStorageLocationMaterialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsStorageLocationMaterialDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsStorageLocationMaterialDo) Group(cols ...field.Expr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsStorageLocationMaterialDo) Having(conds ...gen.Condition) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsStorageLocationMaterialDo) Limit(limit int) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsStorageLocationMaterialDo) Offset(offset int) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsStorageLocationMaterialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsStorageLocationMaterialDo) Unscoped() *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsStorageLocationMaterialDo) Create(values ...*model_openiiot.WmsStorageLocationMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsStorageLocationMaterialDo) CreateInBatches(values []*model_openiiot.WmsStorageLocationMaterial, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsStorageLocationMaterialDo) Save(values ...*model_openiiot.WmsStorageLocationMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsStorageLocationMaterialDo) First() (*model_openiiot.WmsStorageLocationMaterial, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocationMaterial), nil
	}
}

func (w wmsStorageLocationMaterialDo) Take() (*model_openiiot.WmsStorageLocationMaterial, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocationMaterial), nil
	}
}

func (w wmsStorageLocationMaterialDo) Last() (*model_openiiot.WmsStorageLocationMaterial, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocationMaterial), nil
	}
}

func (w wmsStorageLocationMaterialDo) Find() ([]*model_openiiot.WmsStorageLocationMaterial, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsStorageLocationMaterial), err
}

func (w wmsStorageLocationMaterialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsStorageLocationMaterial, err error) {
	buf := make([]*model_openiiot.WmsStorageLocationMaterial, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsStorageLocationMaterialDo) FindInBatches(result *[]*model_openiiot.WmsStorageLocationMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsStorageLocationMaterialDo) Attrs(attrs ...field.AssignExpr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsStorageLocationMaterialDo) Assign(attrs ...field.AssignExpr) *wmsStorageLocationMaterialDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsStorageLocationMaterialDo) Joins(fields ...field.RelationField) *wmsStorageLocationMaterialDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsStorageLocationMaterialDo) Preload(fields ...field.RelationField) *wmsStorageLocationMaterialDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsStorageLocationMaterialDo) FirstOrInit() (*model_openiiot.WmsStorageLocationMaterial, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocationMaterial), nil
	}
}

func (w wmsStorageLocationMaterialDo) FirstOrCreate() (*model_openiiot.WmsStorageLocationMaterial, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsStorageLocationMaterial), nil
	}
}

func (w wmsStorageLocationMaterialDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsStorageLocationMaterial, count int64, err error) {
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

func (w wmsStorageLocationMaterialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsStorageLocationMaterialDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsStorageLocationMaterialDo) Delete(models ...*model_openiiot.WmsStorageLocationMaterial) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsStorageLocationMaterialDo) withDO(do gen.Dao) *wmsStorageLocationMaterialDo {
	w.DO = *do.(*gen.DO)
	return w
}

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

func newWmsRfidMaterial(db *gorm.DB, opts ...gen.DOOption) wmsRfidMaterial {
	_wmsRfidMaterial := wmsRfidMaterial{}

	_wmsRfidMaterial.wmsRfidMaterialDo.UseDB(db, opts...)
	_wmsRfidMaterial.wmsRfidMaterialDo.UseModel(&model_openiiot.WmsRfidMaterial{})

	tableName := _wmsRfidMaterial.wmsRfidMaterialDo.TableName()
	_wmsRfidMaterial.ALL = field.NewAsterisk(tableName)
	_wmsRfidMaterial.ID = field.NewInt64(tableName, "id")
	_wmsRfidMaterial.Rfid = field.NewString(tableName, "rfid")
	_wmsRfidMaterial.MaterialID = field.NewInt64(tableName, "material_id")
	_wmsRfidMaterial.Quantity = field.NewInt32(tableName, "quantity")
	_wmsRfidMaterial.UpdateTime = field.NewTime(tableName, "update_time")
	_wmsRfidMaterial.CreateTime = field.NewTime(tableName, "create_time")

	_wmsRfidMaterial.fillFieldMap()

	return _wmsRfidMaterial
}

type wmsRfidMaterial struct {
	wmsRfidMaterialDo wmsRfidMaterialDo

	ALL        field.Asterisk
	ID         field.Int64
	Rfid       field.String
	MaterialID field.Int64
	Quantity   field.Int32
	UpdateTime field.Time
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (w wmsRfidMaterial) Table(newTableName string) *wmsRfidMaterial {
	w.wmsRfidMaterialDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsRfidMaterial) As(alias string) *wmsRfidMaterial {
	w.wmsRfidMaterialDo.DO = *(w.wmsRfidMaterialDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsRfidMaterial) updateTableName(table string) *wmsRfidMaterial {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Rfid = field.NewString(table, "rfid")
	w.MaterialID = field.NewInt64(table, "material_id")
	w.Quantity = field.NewInt32(table, "quantity")
	w.UpdateTime = field.NewTime(table, "update_time")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *wmsRfidMaterial) WithContext(ctx context.Context) *wmsRfidMaterialDo {
	return w.wmsRfidMaterialDo.WithContext(ctx)
}

func (w wmsRfidMaterial) TableName() string { return w.wmsRfidMaterialDo.TableName() }

func (w wmsRfidMaterial) Alias() string { return w.wmsRfidMaterialDo.Alias() }

func (w wmsRfidMaterial) Columns(cols ...field.Expr) gen.Columns {
	return w.wmsRfidMaterialDo.Columns(cols...)
}

func (w *wmsRfidMaterial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsRfidMaterial) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["rfid"] = w.Rfid
	w.fieldMap["material_id"] = w.MaterialID
	w.fieldMap["quantity"] = w.Quantity
	w.fieldMap["update_time"] = w.UpdateTime
	w.fieldMap["create_time"] = w.CreateTime
}

func (w wmsRfidMaterial) clone(db *gorm.DB) wmsRfidMaterial {
	w.wmsRfidMaterialDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsRfidMaterial) replaceDB(db *gorm.DB) wmsRfidMaterial {
	w.wmsRfidMaterialDo.ReplaceDB(db)
	return w
}

type wmsRfidMaterialDo struct{ gen.DO }

func (w wmsRfidMaterialDo) Debug() *wmsRfidMaterialDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsRfidMaterialDo) WithContext(ctx context.Context) *wmsRfidMaterialDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsRfidMaterialDo) ReadDB() *wmsRfidMaterialDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsRfidMaterialDo) WriteDB() *wmsRfidMaterialDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsRfidMaterialDo) Session(config *gorm.Session) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsRfidMaterialDo) Clauses(conds ...clause.Expression) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsRfidMaterialDo) Returning(value interface{}, columns ...string) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsRfidMaterialDo) Not(conds ...gen.Condition) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsRfidMaterialDo) Or(conds ...gen.Condition) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsRfidMaterialDo) Select(conds ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsRfidMaterialDo) Where(conds ...gen.Condition) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsRfidMaterialDo) Order(conds ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsRfidMaterialDo) Distinct(cols ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsRfidMaterialDo) Omit(cols ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsRfidMaterialDo) Join(table schema.Tabler, on ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsRfidMaterialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsRfidMaterialDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsRfidMaterialDo) Group(cols ...field.Expr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsRfidMaterialDo) Having(conds ...gen.Condition) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsRfidMaterialDo) Limit(limit int) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsRfidMaterialDo) Offset(offset int) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsRfidMaterialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsRfidMaterialDo) Unscoped() *wmsRfidMaterialDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsRfidMaterialDo) Create(values ...*model_openiiot.WmsRfidMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsRfidMaterialDo) CreateInBatches(values []*model_openiiot.WmsRfidMaterial, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsRfidMaterialDo) Save(values ...*model_openiiot.WmsRfidMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsRfidMaterialDo) First() (*model_openiiot.WmsRfidMaterial, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsRfidMaterial), nil
	}
}

func (w wmsRfidMaterialDo) Take() (*model_openiiot.WmsRfidMaterial, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsRfidMaterial), nil
	}
}

func (w wmsRfidMaterialDo) Last() (*model_openiiot.WmsRfidMaterial, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsRfidMaterial), nil
	}
}

func (w wmsRfidMaterialDo) Find() ([]*model_openiiot.WmsRfidMaterial, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsRfidMaterial), err
}

func (w wmsRfidMaterialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsRfidMaterial, err error) {
	buf := make([]*model_openiiot.WmsRfidMaterial, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsRfidMaterialDo) FindInBatches(result *[]*model_openiiot.WmsRfidMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsRfidMaterialDo) Attrs(attrs ...field.AssignExpr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsRfidMaterialDo) Assign(attrs ...field.AssignExpr) *wmsRfidMaterialDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsRfidMaterialDo) Joins(fields ...field.RelationField) *wmsRfidMaterialDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsRfidMaterialDo) Preload(fields ...field.RelationField) *wmsRfidMaterialDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsRfidMaterialDo) FirstOrInit() (*model_openiiot.WmsRfidMaterial, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsRfidMaterial), nil
	}
}

func (w wmsRfidMaterialDo) FirstOrCreate() (*model_openiiot.WmsRfidMaterial, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsRfidMaterial), nil
	}
}

func (w wmsRfidMaterialDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsRfidMaterial, count int64, err error) {
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

func (w wmsRfidMaterialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsRfidMaterialDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsRfidMaterialDo) Delete(models ...*model_openiiot.WmsRfidMaterial) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsRfidMaterialDo) withDO(do gen.Dao) *wmsRfidMaterialDo {
	w.DO = *do.(*gen.DO)
	return w
}
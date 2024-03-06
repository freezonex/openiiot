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

func newWmsMaterial(db *gorm.DB, opts ...gen.DOOption) wmsMaterial {
	_wmsMaterial := wmsMaterial{}

	_wmsMaterial.wmsMaterialDo.UseDB(db, opts...)
	_wmsMaterial.wmsMaterialDo.UseModel(&model_openiiot.WmsMaterial{})

	tableName := _wmsMaterial.wmsMaterialDo.TableName()
	_wmsMaterial.ALL = field.NewAsterisk(tableName)
	_wmsMaterial.ID = field.NewInt64(tableName, "id")
	_wmsMaterial.Rfid = field.NewString(tableName, "rfid")
	_wmsMaterial.ProductCode = field.NewString(tableName, "product_code")
	_wmsMaterial.Name = field.NewString(tableName, "name")
	_wmsMaterial.StorageLocationID = field.NewInt64(tableName, "storage_location_id")
	_wmsMaterial.ProductType = field.NewString(tableName, "product_type")
	_wmsMaterial.Unit = field.NewString(tableName, "unit")
	_wmsMaterial.Note = field.NewString(tableName, "note")
	_wmsMaterial.UpdateTime = field.NewTime(tableName, "update_time")
	_wmsMaterial.CreateTime = field.NewTime(tableName, "create_time")

	_wmsMaterial.fillFieldMap()

	return _wmsMaterial
}

type wmsMaterial struct {
	wmsMaterialDo wmsMaterialDo

	ALL               field.Asterisk
	ID                field.Int64
	Rfid              field.String
	ProductCode       field.String
	Name              field.String
	StorageLocationID field.Int64
	ProductType       field.String
	Unit              field.String
	Note              field.String
	UpdateTime        field.Time
	CreateTime        field.Time

	fieldMap map[string]field.Expr
}

func (w wmsMaterial) Table(newTableName string) *wmsMaterial {
	w.wmsMaterialDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsMaterial) As(alias string) *wmsMaterial {
	w.wmsMaterialDo.DO = *(w.wmsMaterialDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsMaterial) updateTableName(table string) *wmsMaterial {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Rfid = field.NewString(table, "rfid")
	w.ProductCode = field.NewString(table, "product_code")
	w.Name = field.NewString(table, "name")
	w.StorageLocationID = field.NewInt64(table, "storage_location_id")
	w.ProductType = field.NewString(table, "product_type")
	w.Unit = field.NewString(table, "unit")
	w.Note = field.NewString(table, "note")
	w.UpdateTime = field.NewTime(table, "update_time")
	w.CreateTime = field.NewTime(table, "create_time")

	w.fillFieldMap()

	return w
}

func (w *wmsMaterial) WithContext(ctx context.Context) *wmsMaterialDo {
	return w.wmsMaterialDo.WithContext(ctx)
}

func (w wmsMaterial) TableName() string { return w.wmsMaterialDo.TableName() }

func (w wmsMaterial) Alias() string { return w.wmsMaterialDo.Alias() }

func (w wmsMaterial) Columns(cols ...field.Expr) gen.Columns { return w.wmsMaterialDo.Columns(cols...) }

func (w *wmsMaterial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsMaterial) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["id"] = w.ID
	w.fieldMap["rfid"] = w.Rfid
	w.fieldMap["product_code"] = w.ProductCode
	w.fieldMap["name"] = w.Name
	w.fieldMap["storage_location_id"] = w.StorageLocationID
	w.fieldMap["product_type"] = w.ProductType
	w.fieldMap["unit"] = w.Unit
	w.fieldMap["note"] = w.Note
	w.fieldMap["update_time"] = w.UpdateTime
	w.fieldMap["create_time"] = w.CreateTime
}

func (w wmsMaterial) clone(db *gorm.DB) wmsMaterial {
	w.wmsMaterialDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsMaterial) replaceDB(db *gorm.DB) wmsMaterial {
	w.wmsMaterialDo.ReplaceDB(db)
	return w
}

type wmsMaterialDo struct{ gen.DO }

func (w wmsMaterialDo) Debug() *wmsMaterialDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsMaterialDo) WithContext(ctx context.Context) *wmsMaterialDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsMaterialDo) ReadDB() *wmsMaterialDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsMaterialDo) WriteDB() *wmsMaterialDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsMaterialDo) Session(config *gorm.Session) *wmsMaterialDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsMaterialDo) Clauses(conds ...clause.Expression) *wmsMaterialDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsMaterialDo) Returning(value interface{}, columns ...string) *wmsMaterialDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsMaterialDo) Not(conds ...gen.Condition) *wmsMaterialDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsMaterialDo) Or(conds ...gen.Condition) *wmsMaterialDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsMaterialDo) Select(conds ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsMaterialDo) Where(conds ...gen.Condition) *wmsMaterialDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsMaterialDo) Order(conds ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsMaterialDo) Distinct(cols ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsMaterialDo) Omit(cols ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsMaterialDo) Join(table schema.Tabler, on ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsMaterialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsMaterialDo) RightJoin(table schema.Tabler, on ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsMaterialDo) Group(cols ...field.Expr) *wmsMaterialDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsMaterialDo) Having(conds ...gen.Condition) *wmsMaterialDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsMaterialDo) Limit(limit int) *wmsMaterialDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsMaterialDo) Offset(offset int) *wmsMaterialDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsMaterialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wmsMaterialDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsMaterialDo) Unscoped() *wmsMaterialDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsMaterialDo) Create(values ...*model_openiiot.WmsMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsMaterialDo) CreateInBatches(values []*model_openiiot.WmsMaterial, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsMaterialDo) Save(values ...*model_openiiot.WmsMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsMaterialDo) First() (*model_openiiot.WmsMaterial, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsMaterial), nil
	}
}

func (w wmsMaterialDo) Take() (*model_openiiot.WmsMaterial, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsMaterial), nil
	}
}

func (w wmsMaterialDo) Last() (*model_openiiot.WmsMaterial, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsMaterial), nil
	}
}

func (w wmsMaterialDo) Find() ([]*model_openiiot.WmsMaterial, error) {
	result, err := w.DO.Find()
	return result.([]*model_openiiot.WmsMaterial), err
}

func (w wmsMaterialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.WmsMaterial, err error) {
	buf := make([]*model_openiiot.WmsMaterial, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsMaterialDo) FindInBatches(result *[]*model_openiiot.WmsMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsMaterialDo) Attrs(attrs ...field.AssignExpr) *wmsMaterialDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsMaterialDo) Assign(attrs ...field.AssignExpr) *wmsMaterialDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsMaterialDo) Joins(fields ...field.RelationField) *wmsMaterialDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsMaterialDo) Preload(fields ...field.RelationField) *wmsMaterialDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsMaterialDo) FirstOrInit() (*model_openiiot.WmsMaterial, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsMaterial), nil
	}
}

func (w wmsMaterialDo) FirstOrCreate() (*model_openiiot.WmsMaterial, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.WmsMaterial), nil
	}
}

func (w wmsMaterialDo) FindByPage(offset int, limit int) (result []*model_openiiot.WmsMaterial, count int64, err error) {
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

func (w wmsMaterialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsMaterialDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsMaterialDo) Delete(models ...*model_openiiot.WmsMaterial) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsMaterialDo) withDO(do gen.Dao) *wmsMaterialDo {
	w.DO = *do.(*gen.DO)
	return w
}

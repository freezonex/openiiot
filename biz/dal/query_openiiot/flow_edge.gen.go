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

func newFlowEdge(db *gorm.DB, opts ...gen.DOOption) flowEdge {
	_flowEdge := flowEdge{}

	_flowEdge.flowEdgeDo.UseDB(db, opts...)
	_flowEdge.flowEdgeDo.UseModel(&model_openiiot.FlowEdge{})

	tableName := _flowEdge.flowEdgeDo.TableName()
	_flowEdge.ALL = field.NewAsterisk(tableName)
	_flowEdge.ID = field.NewInt64(tableName, "id")
	_flowEdge.FlowID = field.NewInt64(tableName, "flow_id")
	_flowEdge.EdgeID = field.NewInt64(tableName, "edge_id")
	_flowEdge.Script = field.NewString(tableName, "script")
	_flowEdge.Script2 = field.NewString(tableName, "script2")
	_flowEdge.Script3 = field.NewString(tableName, "script3")
	_flowEdge.Script4 = field.NewString(tableName, "script4")

	_flowEdge.fillFieldMap()

	return _flowEdge
}

type flowEdge struct {
	flowEdgeDo flowEdgeDo

	ALL     field.Asterisk
	ID      field.Int64
	FlowID  field.Int64
	EdgeID  field.Int64
	Script  field.String
	Script2 field.String
	Script3 field.String
	Script4 field.String

	fieldMap map[string]field.Expr
}

func (f flowEdge) Table(newTableName string) *flowEdge {
	f.flowEdgeDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f flowEdge) As(alias string) *flowEdge {
	f.flowEdgeDo.DO = *(f.flowEdgeDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *flowEdge) updateTableName(table string) *flowEdge {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.FlowID = field.NewInt64(table, "flow_id")
	f.EdgeID = field.NewInt64(table, "edge_id")
	f.Script = field.NewString(table, "script")
	f.Script2 = field.NewString(table, "script2")
	f.Script3 = field.NewString(table, "script3")
	f.Script4 = field.NewString(table, "script4")

	f.fillFieldMap()

	return f
}

func (f *flowEdge) WithContext(ctx context.Context) *flowEdgeDo { return f.flowEdgeDo.WithContext(ctx) }

func (f flowEdge) TableName() string { return f.flowEdgeDo.TableName() }

func (f flowEdge) Alias() string { return f.flowEdgeDo.Alias() }

func (f flowEdge) Columns(cols ...field.Expr) gen.Columns { return f.flowEdgeDo.Columns(cols...) }

func (f *flowEdge) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *flowEdge) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["flow_id"] = f.FlowID
	f.fieldMap["edge_id"] = f.EdgeID
	f.fieldMap["script"] = f.Script
	f.fieldMap["script2"] = f.Script2
	f.fieldMap["script3"] = f.Script3
	f.fieldMap["script4"] = f.Script4
}

func (f flowEdge) clone(db *gorm.DB) flowEdge {
	f.flowEdgeDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f flowEdge) replaceDB(db *gorm.DB) flowEdge {
	f.flowEdgeDo.ReplaceDB(db)
	return f
}

type flowEdgeDo struct{ gen.DO }

func (f flowEdgeDo) Debug() *flowEdgeDo {
	return f.withDO(f.DO.Debug())
}

func (f flowEdgeDo) WithContext(ctx context.Context) *flowEdgeDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f flowEdgeDo) ReadDB() *flowEdgeDo {
	return f.Clauses(dbresolver.Read)
}

func (f flowEdgeDo) WriteDB() *flowEdgeDo {
	return f.Clauses(dbresolver.Write)
}

func (f flowEdgeDo) Session(config *gorm.Session) *flowEdgeDo {
	return f.withDO(f.DO.Session(config))
}

func (f flowEdgeDo) Clauses(conds ...clause.Expression) *flowEdgeDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f flowEdgeDo) Returning(value interface{}, columns ...string) *flowEdgeDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f flowEdgeDo) Not(conds ...gen.Condition) *flowEdgeDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f flowEdgeDo) Or(conds ...gen.Condition) *flowEdgeDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f flowEdgeDo) Select(conds ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f flowEdgeDo) Where(conds ...gen.Condition) *flowEdgeDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f flowEdgeDo) Order(conds ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f flowEdgeDo) Distinct(cols ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f flowEdgeDo) Omit(cols ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f flowEdgeDo) Join(table schema.Tabler, on ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f flowEdgeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f flowEdgeDo) RightJoin(table schema.Tabler, on ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f flowEdgeDo) Group(cols ...field.Expr) *flowEdgeDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f flowEdgeDo) Having(conds ...gen.Condition) *flowEdgeDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f flowEdgeDo) Limit(limit int) *flowEdgeDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f flowEdgeDo) Offset(offset int) *flowEdgeDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f flowEdgeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *flowEdgeDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f flowEdgeDo) Unscoped() *flowEdgeDo {
	return f.withDO(f.DO.Unscoped())
}

func (f flowEdgeDo) Create(values ...*model_openiiot.FlowEdge) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f flowEdgeDo) CreateInBatches(values []*model_openiiot.FlowEdge, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f flowEdgeDo) Save(values ...*model_openiiot.FlowEdge) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f flowEdgeDo) First() (*model_openiiot.FlowEdge, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowEdge), nil
	}
}

func (f flowEdgeDo) Take() (*model_openiiot.FlowEdge, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowEdge), nil
	}
}

func (f flowEdgeDo) Last() (*model_openiiot.FlowEdge, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowEdge), nil
	}
}

func (f flowEdgeDo) Find() ([]*model_openiiot.FlowEdge, error) {
	result, err := f.DO.Find()
	return result.([]*model_openiiot.FlowEdge), err
}

func (f flowEdgeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model_openiiot.FlowEdge, err error) {
	buf := make([]*model_openiiot.FlowEdge, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f flowEdgeDo) FindInBatches(result *[]*model_openiiot.FlowEdge, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f flowEdgeDo) Attrs(attrs ...field.AssignExpr) *flowEdgeDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f flowEdgeDo) Assign(attrs ...field.AssignExpr) *flowEdgeDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f flowEdgeDo) Joins(fields ...field.RelationField) *flowEdgeDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f flowEdgeDo) Preload(fields ...field.RelationField) *flowEdgeDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f flowEdgeDo) FirstOrInit() (*model_openiiot.FlowEdge, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowEdge), nil
	}
}

func (f flowEdgeDo) FirstOrCreate() (*model_openiiot.FlowEdge, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model_openiiot.FlowEdge), nil
	}
}

func (f flowEdgeDo) FindByPage(offset int, limit int) (result []*model_openiiot.FlowEdge, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f flowEdgeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f flowEdgeDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f flowEdgeDo) Delete(models ...*model_openiiot.FlowEdge) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *flowEdgeDo) withDO(do gen.Dao) *flowEdgeDo {
	f.DO = *do.(*gen.DO)
	return f
}

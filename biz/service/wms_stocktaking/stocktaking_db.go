package wms_stocktaking

import (
	"context"
	"errors"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddStocktakingDB will add Stocktaking record to the DB.
func (a *WmsStocktakingService) AddStocktakingDB(ctx context.Context, refid string, type1 string, storagelocationid int64, operator string, result string) (int64, error) {

	if refid == "" {
		return -1, errors.New("refid can not be empty")
	}
	if type1 == "" {
		return -1, errors.New("type1 can not be empty")
	}
	if storagelocationid == 0 {
		return -1, errors.New("storagelocationid can not be empty")
	}
	if operator == "" {
		return -1, errors.New("operator can not be empty")
	}
	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.RefID.Eq(refid)).First()
	if existRecord != nil {
		return -1, errors.New("refid exist")
	}
	id := common.GetUUID()

	newRecord := &model_openiiot.WmsStocktaking{
		RefID:           refid,
		Type:            type1,
		StorageLocation: storagelocationid,
		Operator:        operator,
		Result:          result,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetStocktakingDB will get stocktaking record from the DB in condition
func (a *WmsStocktakingService) GetStocktakingDB(ctx context.Context, id int64, refid string) ([]*model_openiiot.WmsStocktaking, error) {

	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if refid != "" {
		tx = tx.Where(table.RefID.Eq(refid))
	}

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateStocktakingDB will update stocktaking record from the DB.
func (a *WmsStocktakingService) UpdateStocktakingDB(ctx context.Context, id int64, refid string, type1 string, storagelocationid int64, operator string, result string) error {
	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("stocktaking does not exist")
	}

	updates := make(map[string]interface{})

	if refid != "" {
		updates[table.RefID.ColumnName().String()] = refid
	}
	if storagelocationid != 0 {
		updates[table.StorageLocation.ColumnName().String()] = storagelocationid
	}
	if type1 != "" {
		updates[table.Type.ColumnName().String()] = type1
	}
	if operator != "" {
		updates[table.Operator.ColumnName().String()] = operator
	}
	if result != "" {
		updates[table.Result.ColumnName().String()] = result
	}
	_, err := tx.Updates(updates)
	return err
}

// DeleteStocktakingDB will delete stocktaking record from the DB.
func (a *WmsStocktakingService) DeleteStocktakingDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("stocktaking does not exist")
	}
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

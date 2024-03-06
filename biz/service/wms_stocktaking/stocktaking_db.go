package wms_stocktaking

import (
	"context"
	"errors"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddStocktakingDB will add Stocktaking record to the DB.
func (a *WmsStocktakingService) AddStocktakingDB(ctx context.Context, refid string, Type string, StorageLocationIds string, operator string, result string) (int64, error) {

	if refid == "" {
		return -1, errors.New("refid can not be empty")
	}
	if Type == "" {
		return -1, errors.New("type1 can not be empty")
	}
	if StorageLocationIds == "" {
		return -1, errors.New("type1 can not be empty")
	}
	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.RefID.Eq(refid)).First()
	if existRecord != nil {
		return -1, errors.New("refid exist")
	}
	id := common.GetUUID()

	newRecord := &model_openiiot.WmsStocktaking{
		ID:                 id,
		RefID:              refid,
		Type:               Type,
		StorageLocationIds: StorageLocationIds,
		Operator:           "",
		Result:             result,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetStocktakingDB will get stocktaking record from the DB in condition
func (a *WmsStocktakingService) GetStocktakingDB(ctx context.Context, id int64, refid string, Type string) ([]*model_openiiot.WmsStocktaking, error) {

	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if refid != "" {
		tx = tx.Where(table.RefID.Eq(refid))
	}
	if Type != "" {
		tx = tx.Where(table.Type.Eq(Type))
	}

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateStocktakingDB will update stocktaking record from the DB.
func (a *WmsStocktakingService) UpdateStocktakingDB(ctx context.Context, id int64, refid string, Type string, Result string) error {
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
	if Type != "" {
		updates[table.Type.ColumnName().String()] = Type
	}
	if Result != "" {
		updates[table.Result.ColumnName().String()] = Result
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

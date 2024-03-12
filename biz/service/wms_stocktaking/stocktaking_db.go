package wms_stocktaking

import (
	"context"
	"errors"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddStocktakingDB will add Stocktaking record to the DB.
func (a *WmsStocktakingService) AddStocktakingDB(ctx context.Context, refid string, Type string, Source string, Note string, Status string) (int64, error) {

	if Type == "" {
		return -1, errors.New("Type can not be empty")
	}
	if Source == "" {
		return -1, errors.New("Source can not be empty")
	}
	table := a.db.DBOpeniiotQuery.WmsStocktaking
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	newRecord := &model_openiiot.WmsStocktaking{
		ID:       id,
		RefID:    refid,
		Type:     Type,
		Source:   Source,
		Note:     &Note,
		Operator: "",
		Status:   &Status,
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
func (a *WmsStocktakingService) UpdateStocktakingDB(ctx context.Context, id int64, refid string, Type string) error {
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

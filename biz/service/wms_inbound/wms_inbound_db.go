package wms_inbound

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
	"gorm.io/gen/field"
)

// AddWmsInboundDB will add wms record to the DB.
func (a *WmsInboundService) AddWmsInboundDB(ctx context.Context, RefID string, Type string, StorageLocation string, MaterialName string, Operator string) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsInbound
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	storageLocationService := storagelocation.DefaultStorageLocationService()
	storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, 0, StorageLocation, nil, "")

	var newRecord = &model_openiiot.WmsInbound{
		ID:                id,
		RefID:             RefID,
		Type:              Type,
		StorageLocationID: storageLocationData[0].ID,
		MaterialName:      MaterialName,
		Operator:          Operator,
	}

	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetWmsInboundDB will get wms record from the DB in condition
func (a *WmsInboundService) GetWmsInboundDB(ctx context.Context, id int64, RefId string) ([]*model_openiiot.WmsInbound, error) {
	table := a.db.DBOpeniiotQuery.WmsInbound
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if RefId != "" {
		tx = tx.Where(table.RefID.Eq(RefId))
	}

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateWmsInboundDB will update wms record from the DB.
func (a *WmsInboundService) UpdateWmsInboundDB(ctx context.Context, id int64, RefID string, Type string, StorageLocation string, MaterialName string, Operator string) error {
	table := a.db.DBOpeniiotQuery.WmsInbound
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	storageLocationService := storagelocation.DefaultStorageLocationService()
	storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, 0, StorageLocation, nil, "")

	if existRecord == nil {
		return errors.New("wms does not exist")
	}
	updates := make(map[string]interface{})
	if RefID != "" {
		updates[table.RefID.ColumnName().String()] = RefID
	}
	if Type != "" {
		updates[table.Type.ColumnName().String()] = Type
	}
	if storageLocationData[0].ID != 0 {
		updates[table.StorageLocationID.ColumnName().String()] = storageLocationData[0]
	}
	if MaterialName != "" {
		updates[table.MaterialName.ColumnName().String()] = MaterialName
	}
	if Operator != "" {
		updates[table.Operator.ColumnName().String()] = Operator
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsInboundDB will delete wms record from the DB.
func (a *WmsInboundService) DeleteWmsInboundDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsInbound
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

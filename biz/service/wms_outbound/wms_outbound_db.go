package wms_outbound

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_material"
	"gorm.io/gen/field"
	"strings"
)

// AddWmsOutboundDB will add wms record to the DB.
func (a *WmsOutboundService) AddWmsOutboundDB(ctx context.Context, Type string, StorageLocation []string, MaterialName string) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	storageLocationService := wms_material.DefaultWmsMaterialService()
	storageLocationData, _ := storageLocationService.GetWmsMaterialDB(ctx, 0, "", "", "", "", common.StringToInt64(StorageLocation[0]))

	joinedStorageLocations := strings.Join(StorageLocation, ",") // 使用逗号或其他分隔符

	var newRecord = &model_openiiot.WmsOutbound{
		ID:                 id,
		RefID:              storageLocationData[0].Rfid,
		Type:               Type,
		StorageLocationIds: joinedStorageLocations,
		MaterialName:       MaterialName,
		Operator:           "",
	}

	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetWmsOutboundDB will get wms record from the DB in condition
func (a *WmsOutboundService) GetWmsOutboundDB(ctx context.Context, id int64, RefId string) ([]*model_openiiot.WmsOutbound, error) {
	table := a.db.DBOpeniiotQuery.WmsOutbound
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

// UpdateWmsOutboundDB will update wms record from the DB.
func (a *WmsOutboundService) UpdateWmsOutboundDB(ctx context.Context, id int64, RefID string, Type string, StorageLocationId string, MaterialName string) error {
	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

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
	if StorageLocationId != "" {
		updates[table.StorageLocationIds.ColumnName().String()] = StorageLocationId
	}
	if MaterialName != "" {
		updates[table.MaterialName.ColumnName().String()] = MaterialName
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsOutboundDB will delete wms record from the DB.
func (a *WmsOutboundService) DeleteWmsOutboundDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

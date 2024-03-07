package wms_inbound

//
//import (
//	"context"
//	"errors"
//	"freezonex/openiiot/biz/dal/model_openiiot"
//	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
//	"freezonex/openiiot/biz/service/utils/common"
//	storagelocation "freezonex/openiiot/biz/service/wms_storage_location"
//	"gorm.io/gen/field"
//	"strings"
//)
//
//// AddWmsInboundDB will add wms record to the DB.
//func (a *WmsInboundService) AddWmsInboundDB(ctx context.Context, type1 string, source string, shelfrecords []*freezonex_openiiot_api.ShelfInventory, note string, status string) (int64, error) {
//
//	table := a.db.DBOpeniiotQuery.WmsInbound
//	tx := table.WithContext(ctx)
//	id := common.GetUUID()
//	if type1 == "PDA" {
//
//		if len(rfids) == 0 {
//			return -1, errors.New("rfids cannot be empty")
//		}
//		operator := "default"
//		rfidsStr := strings.Join(rfids, ",")
//		storageLocationService := storagelocation.DefaultStorageLocationService()
//		storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, storagelocationId, "")
//		status1 := "done"
//		var newRecord = &model_openiiot.WmsInbound{
//			ID:     id,
//			RefID:  rfidsStr,
//			Type:   type1,
//			Note:   &note,
//			Status: &status1,
//
//			Source:   source,
//			Operator: operator,
//		}
//
//		err := tx.Create(newRecord)
//		if err != nil {
//			return -1, err
//		}
//
//	}
//
//	if type1 == "manual" {
//		if len(materialids) == 0 {
//			return -1, errors.New("materialids cannot be empty")
//		}
//		operator := "default"
//		rfidsStr := strings.Join(rfids, ",")
//		storageLocationService := storagelocation.DefaultStorageLocationService()
//		storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, storagelocationId, "")
//
//		var newRecord = &model_openiiot.WmsInbound{
//			ID:                id,
//			RefID:             rfidsStr,
//			Type:              type1,
//			StorageLocationID: storageLocationData[0].ID,
//			MaterialName:      materialname,
//			Source:            source,
//			Operator:          operator,
//		}
//
//		err := tx.Create(newRecord)
//		if err != nil {
//			return -1, err
//		}
//	}
//
//	return id, nil
//}
//
//// GetWmsInboundDB will get wms record from the DB in condition
//func (a *WmsInboundService) GetWmsInboundDB(ctx context.Context, id int64, RefId string) ([]*model_openiiot.WmsInbound, error) {
//	table := a.db.DBOpeniiotQuery.WmsInbound
//	tx := table.WithContext(ctx).Select(field.ALL)
//	if id != 0 {
//		tx = tx.Where(table.ID.Eq(id))
//	}
//	if RefId != "" {
//		tx = tx.Where(table.RefID.Eq(RefId))
//	}
//
//	data, err := tx.Find()
//	if err != nil {
//		return nil, err
//	}
//
//	return data, nil
//}
//
//// UpdateWmsInboundDB will update wms record from the DB.
//func (a *WmsInboundService) UpdateWmsInboundDB(ctx context.Context, id int64, RefID string, Type string, StorageLocation string, MaterialName string) error {
//	table := a.db.DBOpeniiotQuery.WmsInbound
//	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
//	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
//
//	storageLocationService := storagelocation.DefaultStorageLocationService()
//	storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, common.StringToInt64(StorageLocation), "")
//
//	if existRecord == nil {
//		return errors.New("wms does not exist")
//	}
//	updates := make(map[string]interface{})
//	if RefID != "" {
//		updates[table.RefID.ColumnName().String()] = RefID
//	}
//	if Type != "" {
//		updates[table.Type.ColumnName().String()] = Type
//	}
//	if storageLocationData[0].ID != 0 {
//		updates[table.StorageLocationID.ColumnName().String()] = storageLocationData[0]
//	}
//	if MaterialName != "" {
//		updates[table.MaterialName.ColumnName().String()] = MaterialName
//	}
//
//	_, err := tx.Updates(updates)
//	return err
//}
//
//// DeleteWmsInboundDB will delete wms record from the DB.
//func (a *WmsInboundService) DeleteWmsInboundDB(ctx context.Context, id int64) error {
//
//	table := a.db.DBOpeniiotQuery.WmsInbound
//	tx := table.WithContext(ctx)
//
//	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
//	if existRecord == nil {
//		return errors.New("wms does not exist")
//	}
//
//	_, err := tx.Where(table.ID.Eq(id)).Delete()
//	return err
//}

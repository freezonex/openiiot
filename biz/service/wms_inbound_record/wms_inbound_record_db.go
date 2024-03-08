package wms_inbound_record

import (
	"context"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	"gorm.io/gen/field"
)

// AddWmsInboundRecordDB will add wms record to the DB.
func (a *WmsInboundRecordService) AddWmsInboundRecordDB(ctx context.Context, InboundId int64, storagelocation int64, MaterialId int64, quantity int32) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsInboundRecord
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	var newRecord = &model_openiiot.WmsInboundRecord{
		ID:              id,
		InboundID:       InboundId,
		MaterialID:      MaterialId,
		StockLocationID: storagelocation,
		Quantity:        quantity,
	}

	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetWmsInboundRecordDB will get wms record from the DB in condition
func (a *WmsInboundRecordService) GetWmsInboundRecordDB(ctx context.Context, id int64, inboundid int64) ([]*model_openiiot.WmsInboundRecord, error) {
	table := a.db.DBOpeniiotQuery.WmsInboundRecord
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if inboundid != 0 {
		tx = tx.Where(table.InboundID.Eq(inboundid))
	}
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

//
//// UpdateWmsInboundRecordDB will update wms record from the DB.
//func (a *WmsInboundRecordService) UpdateWmsInboundRecordDB(ctx context.Context, id int64, InboundId string, Rfid string, MaterialId int64) error {
//	table := a.db.DBOpeniiotQuery.WmsInboundRecord
//	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
//	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
//	if existRecord == nil {
//		return errors.New("wms does not exist")
//	}
//	updates := make(map[string]interface{})
//	if InboundId != "" {
//		updates[table.InboundID.ColumnName().String()] = InboundId
//	}
//	if Rfid != "" {
//		updates[table.Rfid.ColumnName().String()] = Rfid
//	}
//	if MaterialId != 0 {
//		updates[table.MaterialID.ColumnName().String()] = MaterialId
//	}
//	_, err := tx.Updates(updates)
//	return err
//}
//
//// DeleteWmsInboundRecordDB will delete wms record from the DB.
//func (a *WmsInboundRecordService) DeleteWmsInboundRecordDB(ctx context.Context, id int64) error {
//
//	table := a.db.DBOpeniiotQuery.WmsInboundRecord
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

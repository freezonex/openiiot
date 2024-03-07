package wms_rfid_material

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	"gorm.io/gen/field"
)

// AddWmsRfidMaterialDB will add wms record to the DB.
func (a *WmsRfidMaterialService) AddWmsRfidMaterialDB(ctx context.Context, Rfid string, MaterialID int64, Quantity int32) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsRfidMaterial
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	var newRecord = &model_openiiot.WmsRfidMaterial{
		ID:         id,
		Rfid:       Rfid,
		MaterialID: MaterialID,
		Quantity:   Quantity,
	}

	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetWmsRfidMaterialDB will get wms record from the DB in condition
func (a *WmsRfidMaterialService) GetWmsRfidMaterialDB(ctx context.Context, id int64, Rfid string, MaterialID int64, Quantity int32) ([]*model_openiiot.WmsRfidMaterial, error) {
	table := a.db.DBOpeniiotQuery.WmsRfidMaterial
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if Rfid != "" {
		tx = tx.Where(table.Rfid.Eq(Rfid))
	}
	if MaterialID != 0 {
		tx = tx.Where(table.MaterialID.Eq(MaterialID))
	}
	if Quantity != 0 {
		tx = tx.Where(table.Quantity.Eq(Quantity))
	}

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateWmsRfidMaterialDB will update wms record from the DB.
func (a *WmsRfidMaterialService) UpdateWmsRfidMaterialDB(ctx context.Context, id int64, Rfid string, MaterialID int64, Quantity int32) error {
	table := a.db.DBOpeniiotQuery.WmsRfidMaterial
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}
	updates := make(map[string]interface{})
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if Rfid != "" {
		tx = tx.Where(table.Rfid.Eq(Rfid))
	}
	if MaterialID != 0 {
		tx = tx.Where(table.MaterialID.Eq(MaterialID))
	}
	if Quantity != 0 {
		tx = tx.Where(table.Quantity.Eq(Quantity))
	}
	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsRfidMaterialDB will delete wms record from the DB.
func (a *WmsRfidMaterialService) DeleteWmsRfidMaterialDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsRfidMaterial
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

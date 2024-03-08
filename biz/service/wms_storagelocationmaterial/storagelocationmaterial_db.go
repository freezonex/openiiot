package wms_storagelocationmaterial

import (
	"context"
	"errors"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddStorageLocationMaterialDB will add StorageLocationMaterial record to the DB.
func (a *WmsStorageLocationMaterialService) AddStorageLocationMaterialDB(ctx context.Context, storagelocationid int64, materialid int64, quantity int32) (int64, error) {

	if materialid == 0 {
		return -1, errors.New("materialid can not be empty")
	}
	if storagelocationid == 0 {
		return -1, errors.New("storagelocationid can not be empty")
	}
	if quantity == 0 {
		return -1, errors.New("quantity can not be empty")
	}
	table := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.StoreLocationID.Eq(storagelocationid), table.MaterialID.Eq(materialid)).First()
	if existRecord != nil {
		return -1, errors.New("materialid already in this storageloacltionid")
	}
	id := common.GetUUID()

	newRecord := &model_openiiot.WmsStorageLocationMaterial{
		MaterialID:      materialid,
		StoreLocationID: storagelocationid,
		Quantity:        quantity,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetStorageLocationMaterialDB will get storagelocationmaterial record from the DB in condition
func (a *WmsStorageLocationMaterialService) GetStorageLocationMaterialDB(ctx context.Context, id int64, storelocationid int64, materialid int64, quantity int32) ([]*model_openiiot.WmsStorageLocationMaterial, error) {

	table := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial
	tx := table.WithContext(ctx).Select(field.ALL)

	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if storelocationid != 0 {
		tx = tx.Where(table.StoreLocationID.Eq(storelocationid))
	}

	if materialid != 0 {
		tx = tx.Where(table.MaterialID.Eq(materialid))
	}
	if quantity != 0 {
		tx = tx.Where(table.Quantity.Eq(quantity))
	}
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateStorageLocationMaterialDB will update storagelocationmaterial record from the DB.
func (a *WmsStorageLocationMaterialService) UpdateStorageLocationMaterialDB(ctx context.Context, id int64, storagelocationid int64, materialid int64, quantity int32) error {
	table := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("storagelocationmaterial does not exist")
	}

	updates := make(map[string]interface{})

	if storagelocationid != 0 {
		updates[table.StoreLocationID.ColumnName().String()] = storagelocationid
	}
	if materialid != 0 {
		updates[table.MaterialID.ColumnName().String()] = materialid
	}
	if quantity != 0 {
		updates[table.Quantity.ColumnName().String()] = quantity
	}

	if quantity == 0 {
		updates[table.MaterialID.ColumnName().String()] = nil
	}
	_, err := tx.Updates(updates)
	return err
}

// DeleteStorageLocationMaterialDB will delete storagelocationmaterial record from the DB.
func (a *WmsStorageLocationMaterialService) DeleteStorageLocationMaterialDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.WmsStorageLocationMaterial
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("storagelocationmaterial does not exist")
	}
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

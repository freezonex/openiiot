package wms_storage_location

import (
	"context"
	"errors"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddStorageLocationDB will add StorageLocation record to the DB.
func (a *WmsStorageLocationService) AddStorageLocationDB(ctx context.Context, warehouseid int64, name string, occupied bool, materialname string) (int64, error) {

	if name == "" {
		return -1, errors.New("name can not be empty")
	}
	if warehouseid == 0 {
		return -1, errors.New("warehouseid can not be empty")
	}
	table := a.db.DBOpeniiotQuery.WmsStorageLocation
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Name.Eq(name)).First()
	if existRecord != nil {
		return -1, errors.New("storagelocationname exist")
	}
	id := common.GetUUID()

	if occupied {
		occupied = false
	}

	newRecord := &model_openiiot.WmsStorageLocation{
		Name:         name,
		MaterialName: &materialname,
		WarehouseID:  warehouseid,
		Occupied:     &occupied,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetStorageLocationDB will get storagelocation record from the DB in condition
func (a *WmsStorageLocationService) GetStorageLocationDB(ctx context.Context, warehouseid int64, name string, occupied *bool, materialname string) ([]*model_openiiot.WmsStorageLocation, error) {

	table := a.db.DBOpeniiotQuery.WmsStorageLocation
	tx := table.WithContext(ctx).Select(field.ALL)
	if name != "" {
		tx = tx.Where(table.Name.Eq(name))
	}
	if warehouseid != 0 {
		tx = tx.Where(table.WarehouseID.Eq(warehouseid))
	}
	if materialname != "" {
		tx = tx.Where(table.MaterialName.Eq(materialname))
	}

	if occupied != nil {
		tx = tx.Where(table.Occupied.Is(*occupied))
	}

	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateStorageLocationDB will update storagelocation record from the DB.
func (a *WmsStorageLocationService) UpdateStorageLocationDB(ctx context.Context, id int64, warehouseid int64, name string, occupied *bool, materialname string) error {
	table := a.db.DBOpeniiotQuery.WmsStorageLocation
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("storagelocation does not exist")
	}

	updates := make(map[string]interface{})

	if name != "" {
		updates[table.Name.ColumnName().String()] = name
	}
	if warehouseid != 0 {
		updates[table.WarehouseID.ColumnName().String()] = warehouseid
	}
	if materialname != "" {
		updates[table.MaterialName.ColumnName().String()] = materialname
	}
	if occupied != nil {
		updates[table.Occupied.ColumnName().String()] = occupied
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteStorageLocationDB will delete storagelocation record from the DB.
func (a *WmsStorageLocationService) DeleteStorageLocationDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.WmsStorageLocation
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("storagelocation does not exist")
	}
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

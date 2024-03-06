package wms_material

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	"gorm.io/gen/field"
	"strconv"
)

// AddWmsMaterialDB will add wms record to the DB.
func (a *WmsMaterialService) AddWmsMaterialDB(ctx context.Context, ProductCode string, Name string, StorageLocation string, ProductType string, Unit string, Note string) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsMaterial
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	StorageLocationInt, _ := strconv.ParseInt(StorageLocation, 10, 64)

	var newRecord = &model_openiiot.WmsMaterial{
		ID:                id,
		ProductCode:       ProductCode,
		Name:              Name,
		StorageLocationID: &StorageLocationInt,
		ProductType:       &ProductType,
		Unit:              &Unit,
		Note:              &Note,
	}

	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetWmsMaterialDB will get wms record from the DB in condition
func (a *WmsMaterialService) GetWmsMaterialDB(ctx context.Context, id int64, ProductCode string, Name string, Rfid string, WarehouseId string, StorageLocationId int64) ([]*model_openiiot.WmsMaterial, error) {
	table := a.db.DBOpeniiotQuery.WmsMaterial
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if ProductCode != "" {
		tx = tx.Where(table.ProductCode.Eq(ProductCode))
	}
	if Name != "" {
		tx = tx.Where(table.Name.Eq(Name))
	}
	if StorageLocationId != 0 {
		tx = tx.Where(table.StorageLocationID.Eq(StorageLocationId))
	}
	if Rfid != "" {
		tx = tx.Where(table.Name.Eq(Name))
	}
	if WarehouseId != "" {
		tx = tx.Where(table.Name.Eq(Name))
	}
	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateWmsMaterialDB will update wms record from the DB.
func (a *WmsMaterialService) UpdateWmsMaterialDB(ctx context.Context, id int64, ProductCode string, Name string, StorageLocationID int64, ProductType string, Unit string, Note string) error {
	table := a.db.DBOpeniiotQuery.WmsMaterial
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}
	updates := make(map[string]interface{})
	if ProductCode != "" {
		updates[table.ProductCode.ColumnName().String()] = ProductCode
	}
	if Name != "" {
		updates[table.Name.ColumnName().String()] = Name
	}
	if StorageLocationID != 0 {
		updates[table.StorageLocationID.ColumnName().String()] = StorageLocationID
	}
	if ProductType != "" {
		updates[table.ProductType.ColumnName().String()] = ProductType
	}
	if Unit != "" {
		updates[table.Unit.ColumnName().String()] = Unit
	}
	if Note != "" {
		updates[table.Note.ColumnName().String()] = Note
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsMaterialDB will delete wms record from the DB.
func (a *WmsMaterialService) DeleteWmsMaterialDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsMaterial
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

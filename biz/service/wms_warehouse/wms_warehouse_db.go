package wms_warehouse

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	"gorm.io/gen/field"
)

// AddWmsWarehouseDB will add wms record to the DB.
func (a *WmsWarehouseService) AddWmsWarehouseDB(ctx context.Context, name string, WarehouseId string, Type string, Manager string, Department string, Email string, ProjectGroup string, Note string) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsWarehouse
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	var newRecord = &model_openiiot.WmsWarehouse{
		ID:           id,
		Name:         name,
		WarehouseID:  WarehouseId,
		Type:         Type,
		Manager:      &Manager,
		Department:   &Department,
		Email:        &Email,
		ProjectGroup: &ProjectGroup,
		Note:         &Note,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetWmsWarehouseDB will get wms record from the DB in condition
func (a *WmsWarehouseService) GetWmsWarehouseDB(ctx context.Context, id int64, name string, WarehouseId string, Type string, Manager string, Department string, Email string, ProjectGroup string) ([]*model_openiiot.WmsWarehouse, error) {
	table := a.db.DBOpeniiotQuery.WmsWarehouse
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if name != "" {
		tx = tx.Where(table.Name.Eq(name))
	}
	if WarehouseId != "" {
		tx = tx.Where(table.WarehouseID.Eq(WarehouseId))
	}
	if Type != "" {
		tx = tx.Where(table.Type.Eq(Type))
	}
	if Manager != "" {
		tx = tx.Where(table.Manager.Eq(Manager))
	}
	if Email != "" {
		tx = tx.Where(table.Email.Eq(Email))
	}
	if Department != "" {
		tx = tx.Where(table.Department.Eq(Department))
	}
	if ProjectGroup != "" {
		tx = tx.Where(table.ProjectGroup.Eq(ProjectGroup))
	}
	//tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateWmsWarehouseDB will update wms record from the DB.
func (a *WmsWarehouseService) UpdateWmsWarehouseDB(ctx context.Context, id int64, name string, WarehouseId string, Type string, Manager string, Department string, Email string, ProjectGroup string, Note string) error {
	table := a.db.DBOpeniiotQuery.WmsWarehouse
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}
	updates := make(map[string]interface{})
	if name != "" {
		updates[table.Name.ColumnName().String()] = name
	}
	if WarehouseId != "" {
		updates[table.WarehouseID.ColumnName().String()] = WarehouseId
	}
	if Type != "" {
		updates[table.Type.ColumnName().String()] = Type
	}
	if Manager != "" {
		updates[table.Manager.ColumnName().String()] = Manager
	}
	if Department != "" {
		updates[table.Department.ColumnName().String()] = Department
	}
	if Email != "" {
		updates[table.Email.ColumnName().String()] = Email
	}
	if ProjectGroup != "" {
		updates[table.ProjectGroup.ColumnName().String()] = ProjectGroup
	}
	if Note != "" {
		updates[table.Note.ColumnName().String()] = Note
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsWarehouseDB will delete wms record from the DB.
func (a *WmsWarehouseService) DeleteWmsWarehouseDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsWarehouse
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

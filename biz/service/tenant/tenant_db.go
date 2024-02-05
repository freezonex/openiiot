package tenant

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	"gorm.io/gen/field"
)

// AddTenantDB will add tenant record to the DB.
func (a *TenantService) AddTenantDB(ctx context.Context, name string, description string, isDefault bool) (int64, error) {
	if name == "" {
		return -1, errors.New("name can not be empty")
	}
	id := common.GetUUID()
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Name.Eq(name)).First()
	if existRecord != nil {
		return -1, errors.New("tenant exist")
	}

	if description == "" {
		description = "tenant"
	}
	if isDefault {
		isDefault = false
	}

	newRecord := &model_openiiot.Tenant{
		ID:          id,
		Name:        name,
		Description: &description,
		IsDefault:   &isDefault,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetTenantDB will get tenant record from the DB in condition
func (a *TenantService) GetTenantDB(ctx context.Context, id int64, name string) ([]*model_openiiot.Tenant, error) {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if name != "" {
		tx = tx.Where(table.Name.Eq(name))
	}
	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateTenantDB will update tenant record from the DB.
func (a *TenantService) UpdateTenantDB(ctx context.Context, id int64, name string, description string) error {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("tenant does not exist")
	}
	updates := make(map[string]interface{})
	if name != "" {
		updates[table.Name.ColumnName().String()] = name
	}
	if description != "" {
		updates[table.Description.ColumnName().String()] = description
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteTenantDB will delete tenant record from the DB.
func (a *TenantService) DeleteTenantDB(ctx context.Context, id int64) (string, error) {

	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx)

	ax := table.WithContext(ctx).Select(field.ALL)
	ax = ax.Where(table.ID.Eq(id))
	data, err := ax.Find()
	if err != nil {
		return "", err
	}
	name := data[0].Name

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return "", errors.New("tenant does not exist")
	}

	_, err = tx.Where(table.ID.Eq(id)).Delete()
	return name, err
}

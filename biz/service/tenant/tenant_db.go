package tenant

import (
	"context"
	"errors"

	"gorm.io/gen/field"

	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddTenantDB will add tenant record to the DB.
func (a *TenantService) AddTenantDB(ctx context.Context, name string, description string, isDefault bool) (int64, error) {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Name.Eq(name)).First()
	if existRecord != nil {
		return -1, errors.New("tenant exist")
	}
	id := common.GetUUID()
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
func (a *TenantService) DeleteTenantDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx)
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}
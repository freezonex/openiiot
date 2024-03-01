package core

import (
	"context"
	"errors"

	"gorm.io/gen/field"

	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddUserDB will add User record to the DB.
func (a *CoreService) AddCoreDB(ctx context.Context, name string, description string, tenant_id int64, url string, username string, password string, type1 string) (int64, error) {
	//empty value rejection
	if name == "" {
		return -1, errors.New("name can not be empty")
	}
	if username == "" {
		return -1, errors.New("username can not be empty")
	}
	if password == "" {
		return -1, errors.New("password can not be empty")
	}
	if tenant_id == 0 {
		return -1, errors.New("tenant_id can not be empty")
	}
	if url == "" {
		return -1, errors.New("url can not be empty")
	}

	table := a.db.DBOpeniiotQuery.Core
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Name.Eq(name)).First()
	if existRecord != nil {
		return -1, errors.New("corename exist")
	}
	id := common.GetUUID()

	if description == "" {
		description = "core"
	}
	if type1 == "" {
		type1 = "other"
	}

	newRecord := &model_openiiot.Core{
		ID:          id,
		Name:        name,
		Description: &description,
		TenantID:    tenant_id,
		URL:         url,
		Username:    &username,
		Password:    &password,
		Type:        type1,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetUserDB will get user record from the DB in condition
func (a *CoreService) GetCoreDB(ctx context.Context, id int64, name string, tenant_id int64, url string, type1 string) ([]*model_openiiot.Core, error) {
	table := a.db.DBOpeniiotQuery.Core
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if name != "" {
		tx = tx.Where(table.Name.Eq(name))
	}
	if tenant_id != 0 {
		tx = tx.Where(table.TenantID.Eq(tenant_id))
	}
	if url != "" {
		tx = tx.Where(table.URL.Eq(url))
	}
	if type1 != "" {
		tx = tx.Where(table.Type.Eq(type1))
	}

	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateUserDB will update user record from the DB.
func (a *CoreService) UpdateCoreDB(ctx context.Context, id int64, name string, description string, tenant_id int64, url string, username string, password string, type1 string) error {
	table := a.db.DBOpeniiotQuery.Core
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("user does not exist")
	}

	updates := make(map[string]interface{})

	if name != "" {
		updates[table.Name.ColumnName().String()] = name
	}
	if description != "" {
		updates[table.Description.ColumnName().String()] = description
	}
	if tenant_id != 0 {
		updates[table.TenantID.ColumnName().String()] = tenant_id
	}
	if url != "" {
		updates[table.URL.ColumnName().String()] = url
	}
	if username != "" {
		updates[table.Username.ColumnName().String()] = username
	}
	if password != "" {
		updates[table.Password.ColumnName().String()] = password
	}
	if type1 != "" {
		updates[table.Type.ColumnName().String()] = type1
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteUserDB will delete user record from the DB.
func (a *CoreService) DeleteCoreDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.Core
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("core does not exist")
	}
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

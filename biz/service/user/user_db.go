package user

import (
	"context"
	"errors"

	"gorm.io/gen/field"

	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddUserDB will add User record to the DB.
func (a *UserService) AddUserDB(ctx context.Context, username string, password string, description string, tenant_id int64, role string, auth_id string, source string) (int64, error) {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Username.Eq(username)).First()
	if existRecord != nil {
		return -1, errors.New("username exist")
	}
	id := common.GetUUID()
	newRecord := &model_openiiot.User{
		ID:          id,
		Username:    username,
		Password:    &password,
		Description: &description,
		TenantID:    tenant_id,
		Role:        role,
		AuthID:      &auth_id,
		Source:      &source,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetUserDB will get user record from the DB in condition
func (a *UserService) GetUserDB(ctx context.Context, username string, tenant_id int64, role string, auth_id string, source string) ([]*model_openiiot.User, error) {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx).Select(field.ALL)
	if username != "" {
		tx = tx.Where(table.Username.Eq(username))
	}
	if tenant_id != 0 {
		tx = tx.Where(table.TenantID.Eq(tenant_id))
	}
	if role != "" {
		tx = tx.Where(table.Role.Eq(role))
	}
	if auth_id != "" {
		tx = tx.Where(table.AuthID.Eq(auth_id))
	}
	if source != "" {
		tx = tx.Where(table.Source.Eq(source))
	}

	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Username)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateUserDB will update user record from the DB.
func (a *UserService) UpdateUserDB(ctx context.Context, id int64, username string, password string, description string, tenant_id int64, role string, auth_id string, source string) error {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("user does not exist")
	}

	updates := make(map[string]interface{})

	if username != "" {
		updates[table.Username.ColumnName().String()] = username
	}
	if password != "" {
		updates[table.Password.ColumnName().String()] = password
	}
	if description != "" {
		updates[table.Description.ColumnName().String()] = description
	}
	if tenant_id != 0 {
		updates[table.TenantID.ColumnName().String()] = tenant_id
	}
	if role != "" {
		updates[table.Role.ColumnName().String()] = role
	}
	if auth_id != "" {
		updates[table.AuthID.ColumnName().String()] = auth_id
	}
	if source != "" {
		updates[table.Source.ColumnName().String()] = source
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteUserDB will delete user record from the DB.
func (a *UserService) DeleteUserDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("user does not exist")
	}
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

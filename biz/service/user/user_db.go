package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gen/field"
	"strings"

	"time"

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
	if auth_id == "" {
		auth_id = username
	}

	if source == "" {
		source = "useradd"
	}
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

func (a *UserService) GetUserByTokenDB(ctx context.Context, usertoken string) (*time.Time, string, error) {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx)

	if usertoken == "" {
		return nil, "", fmt.Errorf("usertoken is empty")
	}

	var data *model_openiiot.User
	data, err := tx.Where(table.Token.Eq(usertoken)).First()
	if err != nil {
		return nil, "", err
	}

	return data.TokenUpdatetime, data.Username, nil
}

func (a *UserService) AddSuposUserID(ctx context.Context, username string, userid string, userrole string) (int64, error) {
	table1 := a.db.DBOpeniiotQuery.Tenant
	tx1 := table1.WithContext(ctx)

	defaultTenant, _ := tx1.Where(table1.IsDefault.Is(true)).First()
	defaultTenantID := defaultTenant.ID

	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx)
	source := "supOS"
	// Insert new edge IDs
	exist_username, err := tx.Where(table.TenantID.Eq(defaultTenantID), table.Username.Eq(username)).First() //tenantid ? is_default
	if exist_username == nil {
		id := common.GetUUID()
		newUser := model_openiiot.User{
			ID:       id,
			Username: username,
			Role:     userrole,
			AuthID:   &userid,
			TenantID: defaultTenantID,
			Source:   &source,
		}
		err := tx.Create(&newUser)
		if err != nil {
			// Handle the error, possibly return or log it
			return 0, err
		}
		return newUser.ID, nil
	} else {
		return exist_username.ID, err
	}
}

func (a *UserService) UpdateSuposToken(ctx context.Context, id int64, token string) error {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	// Insert new edge IDs
	if existRecord == nil {
		return errors.New("user does not exist")
	}

	updates := make(map[string]interface{})

	if token != "" {
		updates[table.Token.ColumnName().String()] = token
		updates[table.TokenUpdatetime.ColumnName().String()] = time.Now()
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteUserDB will delete user record from the DB.
func (a *UserService) DeleteUserToken(ctx context.Context, token string) error {
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Token.Eq(token)).First()
	// Insert new edge IDs
	if existRecord == nil {
		return errors.New("user does not exist")
	}

	updates := make(map[string]interface{})

	updates[table.Token.ColumnName().String()] = ""

	_, err := tx.Updates(updates)
	return err
}

func (a *UserService) UserLoginDB(ctx context.Context, username string, password string, tenant_name string) (int64, string, string, error) {
	table1 := a.db.DBOpeniiotQuery.Tenant
	tx1 := table1.WithContext(ctx)
	existTenant, err := tx1.Where(table1.Name.Eq(tenant_name)).First()
	if existTenant == nil {
		return 0, "", "", errors.New("tenant does not exist")
	}
	if err != nil {
		return 0, "", "", err
	}

	tenant_id := existTenant.ID
	table := a.db.DBOpeniiotQuery.User
	tx := table.WithContext(ctx)

	existUser, err := tx.Where(table.Username.Eq(username), table.TenantID.Eq(tenant_id)).First()

	if err != nil {
		return 0, "", "", err
	}
	// Insert new edge IDs
	if existUser == nil {
		return 0, "", "", errors.New("user does not exist")
	}
	if *existUser.Password != password {
		return 0, "", "", errors.New("wrong password")
	}

	Accesstoken := strings.Replace(uuid.New().String(), "-", "", -1)
	id := existUser.ID
	role := existUser.Role

	return id, Accesstoken, role, err
}

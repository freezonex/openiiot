package flow

import (
	"context"
	"errors"
	"gorm.io/gen/field"

	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
)

// AddUserDB will add User record to the DB.
func (a *FlowService) AddFlowDB(ctx context.Context, name string, description string, tenantid int64, flowtype string) (int64, error) {
	table := a.db.DBOpeniiotQuery.Flow
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Name.Eq(name)).First()
	if existRecord != nil {
		return -1, errors.New("flow name exist")
	}
	currentusername := ctx.Value("currentusername").(string)

	id := common.GetUUID()
	newRecord := &model_openiiot.Flow{
		ID:             id,
		Name:           name,
		Description:    &description,
		LastModifiedBy: &currentusername,
		TenantID:       tenantid,
		FlowType:       &flowtype,
	}
	err := tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (a *FlowService) AddFlowEdge(ctx context.Context, flowID int64, edgeIDs []int64) error {
	table := a.db.DBOpeniiotQuery.FlowEdge
	tx := table.WithContext(ctx)
	id := common.GetUUID()
	for _, edgeID := range edgeIDs {
		newRecord := &model_openiiot.FlowEdge{
			ID:     id,
			FlowID: flowID,
			EdgeID: edgeID,
		}
		err := tx.Create(newRecord)
		if err != nil {
			return err
		}

	}

	return nil
}

func (a *FlowService) AddFlowCore(ctx context.Context, flowID int64, coreIDs []int64) error {
	table := a.db.DBOpeniiotQuery.FlowCore
	tx := table.WithContext(ctx)
	id := common.GetUUID()
	for _, coreID := range coreIDs {
		newRecord := &model_openiiot.FlowCore{
			ID:     id,
			FlowID: flowID,
			CoreID: coreID,
		}
		err := tx.Create(newRecord)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *FlowService) AddFlowApp(ctx context.Context, flowID int64, appIDs []int64) error {
	table := a.db.DBOpeniiotQuery.FlowApp
	tx := table.WithContext(ctx)
	id := common.GetUUID()
	for _, appID := range appIDs {
		newRecord := &model_openiiot.FlowApp{
			ID:     id,
			FlowID: flowID,
			AppID:  appID,
		}
		err := tx.Create(newRecord)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetUserDB will get user record from the DB in condition
func (a *FlowService) GetFlowDB(ctx context.Context, id int64, name string, tenantid int64, lastmodifiedby string, flowtype string) ([]*model_openiiot.Flow, error) {
	table := a.db.DBOpeniiotQuery.Flow
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if name != "" {
		tx = tx.Where(table.Name.Eq(name))
	}
	if tenantid != 0 {
		tx = tx.Where(table.TenantID.Eq(tenantid))
	}
	if lastmodifiedby != "" {
		tx = tx.Where(table.LastModifiedBy.Eq(lastmodifiedby))
	}
	if flowtype != "" {
		tx = tx.Where(table.FlowType.Eq(flowtype))
	}
	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a *FlowService) GetFlowEdgeIDs(ctx context.Context, flowID int64) ([]int64, error) {
	table := a.db.DBOpeniiotQuery.FlowEdge // Replace with the actual reference to the flow_edge table
	tx := table.WithContext(ctx)

	if flowID != 0 {
		tx = tx.Where(table.FlowID.Eq(flowID))
	}

	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.FlowID)

	// Declare data as a slice of FlowEdge objects
	var data []*model_openiiot.FlowEdge

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	// Extract EdgeID from each FlowEdge object in the data slice
	var edgeIDs []int64
	for _, edgeRecord := range data {
		edgeIDs = append(edgeIDs, edgeRecord.EdgeID)
	}

	return edgeIDs, nil
}

func (a *FlowService) GetFlowCoreIDs(ctx context.Context, flowID int64) ([]int64, error) {
	table := a.db.DBOpeniiotQuery.FlowCore // Replace with the actual reference to the flow_edge table
	tx := table.WithContext(ctx)

	if flowID != 0 {
		tx = tx.Where(table.FlowID.Eq(flowID))
	}

	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.FlowID)

	// Declare data as a slice of FlowEdge objects
	var data []*model_openiiot.FlowCore

	data, err := tx.Find() // Ensure you pass a reference to the data slice
	if err != nil {
		return nil, err
	}

	// Extract EdgeID from each FlowEdge object in the data slice
	var edgeIDs []int64
	for _, edgeRecord := range data {
		edgeIDs = append(edgeIDs, edgeRecord.CoreID)
	}

	return edgeIDs, nil
}

func (a *FlowService) GetFlowAppIDs(ctx context.Context, flowID int64) ([]int64, error) {
	table := a.db.DBOpeniiotQuery.FlowApp // Replace with the actual reference to the flow_edge table
	tx := table.WithContext(ctx)

	if flowID != 0 {
		tx = tx.Where(table.FlowID.Eq(flowID))
	}

	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.FlowID)

	// Declare data as a slice of FlowEdge objects
	var data []*model_openiiot.FlowApp

	data, err := tx.Find() // Ensure you pass a reference to the data slice
	if err != nil {
		return nil, err
	}

	// Extract EdgeID from each FlowEdge object in the data slice
	var edgeIDs []int64
	for _, edgeRecord := range data {
		edgeIDs = append(edgeIDs, edgeRecord.AppID)
	}

	return edgeIDs, nil
}

// UpdateUserDB will update user record from the DB.
func (a *FlowService) UpdateFlowDB(ctx context.Context, id int64, name string, description string, tenantid int64, flowtype string) error {
	table := a.db.DBOpeniiotQuery.Flow
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("flow does not exist")
	}
	currentusername := ctx.Value("currentusername").(string)
	updates := make(map[string]interface{})
	if name != "" {
		updates[table.Name.ColumnName().String()] = name
	}
	if description != "" {
		updates[table.Description.ColumnName().String()] = description
	}
	if tenantid != 0 {
		updates[table.TenantID.ColumnName().String()] = tenantid
	}
	if currentusername != "" {
		updates[table.LastModifiedBy.ColumnName().String()] = currentusername
	}
	if flowtype != "" {
		updates[table.FlowType.ColumnName().String()] = flowtype
	}

	_, err := tx.Updates(updates)
	return err
}

//// delete all and update
//func (a *FlowService) UpdateFlowEdgeDB(ctx context.Context, flowid int64, edgeids []int64) error {
//	table := a.db.DBOpeniiotQuery.FlowEdge
//	tx := table.WithContext(ctx)
//	_, err := tx.Where(table.FlowID.Eq(flowid)).Delete()
//	if err != nil {
//		return err
//	}
//
//	for _, edgeID := range edgeids {
//		id := common.GetUUID()
//		newEdge := model_openiiot.FlowEdge{
//			ID:     id,
//			FlowID: flowid,
//			EdgeID: edgeID,
//		}
//		err := tx.Create(&newEdge)
//		if err != nil {
//			return err
//		}
//	}
//
//	return err
//}

// delete all and update
func (a *FlowService) UpdateFlowEdgeDB(ctx context.Context, flowid int64, edgeids []int64) error {
	table := a.db.DBOpeniiotQuery.FlowEdge
	tx := table.WithContext(ctx)

	// Insert new edge IDs
	for _, edgeID := range edgeids {

		exist_edgeid, _ := tx.Where(table.FlowID.Eq(flowid), table.EdgeID.Eq(edgeID)).First()

		if exist_edgeid == nil {
			newEdge := model_openiiot.FlowEdge{
				FlowID: flowid,
				EdgeID: edgeID,
			}
			err := tx.Create(&newEdge)
			if err != nil {
				return err
			}
		}
	}

	existingEdgeIDs, err := a.GetFlowEdgeIDs(ctx, flowid)
	if err != nil {
		return err // Handle error appropriately
	}
	// Create a map for quick lookup of incoming edgeids

	inputEdgeMap := make(map[int64]bool)
	for _, id := range edgeids {
		inputEdgeMap[id] = true
	}

	for _, existingID := range existingEdgeIDs {
		if !inputEdgeMap[existingID] {
			_, err := tx.Where(table.FlowID.Eq(flowid), table.EdgeID.Eq(existingID)).Delete(&model_openiiot.FlowEdge{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *FlowService) UpdateFlowCoreDB(ctx context.Context, flowid int64, coreids []int64) error {
	table := a.db.DBOpeniiotQuery.FlowCore
	tx := table.WithContext(ctx)

	for _, coreID := range coreids {

		exist_coreid, _ := tx.Where(table.FlowID.Eq(flowid), table.CoreID.Eq(coreID)).First()

		if exist_coreid == nil {
			newCore := model_openiiot.FlowCore{
				FlowID: flowid,
				CoreID: coreID,
			}
			err := tx.Create(&newCore)
			if err != nil {
				return err
			}
		}
	}

	existingCoreIDs, err := a.GetFlowCoreIDs(ctx, flowid)
	if err != nil {
		return err // Handle error appropriately
	}
	// Create a map for quick lookup of incoming edgeids

	inputCoreMap := make(map[int64]bool)
	for _, id := range coreids {
		inputCoreMap[id] = true
	}

	for _, existingID := range existingCoreIDs {
		if !inputCoreMap[existingID] {
			_, err := tx.Where(table.FlowID.Eq(flowid), table.CoreID.Eq(existingID)).Delete(&model_openiiot.FlowCore{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *FlowService) UpdateFlowAppDB(ctx context.Context, flowid int64, appids []int64) error {
	table := a.db.DBOpeniiotQuery.FlowApp
	tx := table.WithContext(ctx)

	for _, appID := range appids {

		exist_edgeid, _ := tx.Where(table.FlowID.Eq(flowid), table.AppID.Eq(appID)).First()

		if exist_edgeid == nil {
			newApp := model_openiiot.FlowApp{
				FlowID: flowid,
				AppID:  appID,
			}
			err := tx.Create(&newApp)
			if err != nil {
				return err
			}
		}
	}

	existingAppIDs, err := a.GetFlowAppIDs(ctx, flowid)
	if err != nil {
		return err // Handle error appropriately
	}
	// Create a map for quick lookup of incoming edgeids

	inputAppMap := make(map[int64]bool)
	for _, id := range appids {
		inputAppMap[id] = true
	}

	for _, existingID := range existingAppIDs {
		if !inputAppMap[existingID] {
			_, err := tx.Where(table.FlowID.Eq(flowid), table.AppID.Eq(existingID)).Delete(&model_openiiot.FlowApp{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// DeleteUserDB will delete user record from the DB.
func (a *FlowService) DeleteFlowDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.Flow
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("flow does not exist")
	}
	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

func (a *FlowService) DeleteFlowEdgeRecords(ctx context.Context, flowid int64) error {
	table := a.db.DBOpeniiotQuery.FlowEdge
	tx := table.WithContext(ctx)

	_, err := tx.Where(table.FlowID.Eq(flowid)).Delete()
	return err
}

func (a *FlowService) DeleteFlowCoreRecords(ctx context.Context, flowid int64) error {
	table := a.db.DBOpeniiotQuery.FlowCore
	tx := table.WithContext(ctx)

	_, err := tx.Where(table.FlowID.Eq(flowid)).Delete()
	return err
}

func (a *FlowService) DeleteFlowAppRecords(ctx context.Context, flowid int64) error {
	table := a.db.DBOpeniiotQuery.FlowApp
	tx := table.WithContext(ctx)

	_, err := tx.Where(table.FlowID.Eq(flowid)).Delete()
	return err
}

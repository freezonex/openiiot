package wms_outbound

import (
	"context"
	"errors"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	wms_outbound_record "freezonex/openiiot/biz/service/wms_outbound_record"
	"gorm.io/gen/field"
	"strings"
	"time"
)

// AddWmsOutboundDB will add wms record to the DB.
func (a *WmsOutboundService) AddWmsOutboundDB(ctx context.Context, Type string, source string, shelfrecords []*freezonex_openiiot_api.ShelfInventory, note string, status string) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx)
	id := common.GetUUID()

	//storageLocationService := wms_material.DefaultWmsMaterialService()
	//storageLocationData, _ := storageLocationService.GetWmsMaterialDB(ctx, 0, "", "", "", "", "")

	if source == "PDA" {
		for _, record := range shelfrecords {

			for _, invotry := range record.Inventory {
				outboundrecordservie := wms_outbound_record.DefaultWmsOutboundRecordService()
				_, err := outboundrecordservie.AddWmsOutboundRecordDB(ctx, id, common.StringToInt64(record.StorageLocationId), common.StringToInt64(invotry.MaterialId), invotry.Quantity)
				if err != nil {
					return -1, err
				}

			}
			operator := "default"
			currentDate := a.generateRefID()

			var newRecord = &model_openiiot.WmsOutbound{
				ID:       id,
				RefID:    currentDate,
				Type:     Type,
				Note:     &note,
				Status:   &status,
				Source:   source,
				Operator: operator,
			}

			err := tx.Create(newRecord)
			if err != nil {
				return -1, err
			}
		}
	}
	if source == "manual" {
		for _, record := range shelfrecords {

			for _, invotry := range record.Inventory {
				outboundrecordservie := wms_outbound_record.DefaultWmsOutboundRecordService()
				_, err := outboundrecordservie.AddWmsOutboundRecordDB(ctx, id, common.StringToInt64(record.StorageLocationId), common.StringToInt64(invotry.MaterialId), invotry.Quantity)
				if err != nil {
					return -1, err
				}

			}
			operator := "default"
			status1 := "done"
			currentDate := a.generateRefID()

			var newRecord = &model_openiiot.WmsOutbound{
				ID:       id,
				RefID:    currentDate,
				Type:     Type,
				Note:     &note,
				Status:   &status1,
				Source:   source,
				Operator: operator,
			}

			err := tx.Create(newRecord)
			if err != nil {
				return -1, err
			}
		}
	}
	return id, nil
}

// GetWmsOutboundDB will get wms record from the DB in condition
func (a *WmsOutboundService) GetWmsOutboundDB(ctx context.Context, id int64, RefId string) ([]*model_openiiot.WmsOutbound, error) {
	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if RefId != "" {
		tx = tx.Where(table.RefID.Eq(RefId))
	}

	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateWmsOutboundDB will update wms record from the DB.
func (a *WmsOutboundService) UpdateWmsOutboundDB(ctx context.Context, id int64, RefID string, Type string, source string, note string, status string) error {
	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))
	existRecord, _ := tx.Where(table.ID.Eq(id)).First()

	if existRecord == nil {
		return errors.New("wms does not exist")
	}
	updates := make(map[string]interface{})
	if RefID != "" {
		updates[table.RefID.ColumnName().String()] = RefID
	}
	if Type != "" {
		updates[table.Type.ColumnName().String()] = Type
	}
	if source != "" {
		updates[table.Source.ColumnName().String()] = source
	}
	if note != "" {
		updates[table.Note.ColumnName().String()] = note
	}
	if status != "" {
		updates[table.Status.ColumnName().String()] = status
	}
	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsOutboundDB will delete wms record from the DB.
func (a *WmsOutboundService) DeleteWmsOutboundDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsOutbound
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

func (a *WmsOutboundService) generateRefID() string {
	// Use the current time to generate the ref_id
	currentTime := time.Now()

	// Format the current time into a string with your desired layout
	// This example includes year, month, day, hour, minute, second, and millisecond
	refID := currentTime.Format("R20060102150405.000")

	// Remove the dot from the milliseconds for a consistent ID format
	refID = strings.Replace(refID, ".", "", -1)

	return refID
}

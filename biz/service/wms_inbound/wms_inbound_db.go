package wms_inbound

import (
	"context"
	"errors"
	"gorm.io/gen/field"
	"strings"
	"time"

	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_inbound_record"
)

// AddWmsInboundDB will add wms record to the DB.
func (a *WmsInboundService) AddWmsInboundDB(ctx context.Context, type1 string, source string, shelfrecords []*freezonex_openiiot_api.ShelfInventory, note string, status string) (int64, error) {

	table := a.db.DBOpeniiotQuery.WmsInbound
	tx := table.WithContext(ctx)
	id := common.GetUUID()
	if source == "PDA" {
		for _, record := range shelfrecords {
			for _, invotry := range record.Inventory {
				inboundrecordservie := wms_inbound_record.DefaultWmsInboundRecordService()
				_, err := inboundrecordservie.AddWmsInboundRecordDB(ctx, id, common.StringToInt64(record.StorageLocationId), common.StringToInt64(invotry.MaterialId), invotry.Quantity)
				if err != nil {
					return -1, err
				}

			}
			operator := "default"
			status1 := "done"
			currentDate := a.generateRefID()

			var newRecord = &model_openiiot.WmsInbound{
				ID:       id,
				RefID:    currentDate,
				Type:     type1,
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
	if source == "manual" {
		for _, record := range shelfrecords {

			for _, invotry := range record.Inventory {
				inboundrecordservie := wms_inbound_record.DefaultWmsInboundRecordService()
				_, err := inboundrecordservie.AddWmsInboundRecordDB(ctx, id, common.StringToInt64(record.StorageLocationId), common.StringToInt64(invotry.MaterialId), invotry.Quantity)
				if err != nil {
					return -1, err
				}

			}
			operator := "default"
			status1 := "done"
			currentDate := a.generateRefID()

			var newRecord = &model_openiiot.WmsInbound{
				ID:       id,
				RefID:    currentDate,
				Type:     type1,
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

	//SELECT s.id
	//FROM wms_storage_location s
	//LEFT JOIN wms_storage_location_material m ON s.id = m.store_location_id
	//WHERE m.id IS NULL;

	return id, nil
}

// GetWmsInboundDB will get wms record from the DB in condition
func (a *WmsInboundService) GetWmsInboundDB(ctx context.Context, id int64, RefId string) ([]*model_openiiot.WmsInbound, error) {
	table := a.db.DBOpeniiotQuery.WmsInbound
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

// UpdateWmsInboundDB will update wms record from the DB.
func (a *WmsInboundService) UpdateWmsInboundDB(ctx context.Context, id int64, RefID string, Type string, Source string, Status string) error {
	table := a.db.DBOpeniiotQuery.WmsInbound
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
	if Source != "" {
		updates[table.Source.ColumnName().String()] = Source
	}
	if Status != "" {
		updates[table.Status.ColumnName().String()] = Status
	}
	_, err := tx.Updates(updates)
	return err
}

// DeleteWmsInboundDB will delete wms record from the DB.
func (a *WmsInboundService) DeleteWmsInboundDB(ctx context.Context, id int64) error {

	table := a.db.DBOpeniiotQuery.WmsInbound
	tx := table.WithContext(ctx)

	existRecord, _ := tx.Where(table.ID.Eq(id)).First()
	if existRecord == nil {
		return errors.New("wms does not exist")
	}

	_, err := tx.Where(table.ID.Eq(id)).Delete()
	return err
}

func (a *WmsInboundService) generateRefID() string {
	// Use the current time to generate the ref_id
	currentTime := time.Now()

	// Format the current time into a string with your desired layout
	// This example includes year, month, day, hour, minute, second, and millisecond
	refID := currentTime.Format("R20060102150405.000")

	// Remove the dot from the milliseconds for a consistent ID format
	refID = strings.Replace(refID, ".", "", -1)

	return refID
}

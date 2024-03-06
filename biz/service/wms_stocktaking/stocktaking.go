package wms_stocktaking

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/service/utils/common"
	"freezonex/openiiot/biz/service/wms_storage_location"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"strings"

	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	"github.com/cloudwego/hertz/pkg/app"
)

func (a *WmsStocktakingService) AddStocktaking(ctx context.Context, req *freezonex_openiiot_api.AddStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddStocktakingResponse, error) {
	var RefIds []string
	var StorageLocationIds []string

	for _, b := range req.Items {
		for _, rfid := range b.Rfids {
			RefIds = append(RefIds, rfid)
		}
		StorageLocationIds = append(StorageLocationIds, b.StorageLocationId)
	}

	RefIdsString := strings.Join(RefIds, ",")
	StorageLocationIdsString := strings.Join(StorageLocationIds, ",")

	stocktakingID, err := a.AddStocktakingDB(ctx, RefIdsString, req.Type, StorageLocationIdsString, "", req.Result)
	if err != nil {
		logs.Error(ctx, "event=AddStocktaking error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddStocktakingResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(stocktakingID)

	return resp, nil
}

// GetStocktaking will get stocktaking record in condition
func (a *WmsStocktakingService) GetStocktaking(ctx context.Context, req *freezonex_openiiot_api.GetStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetStocktakingResponse, error) {
	stocktakings, err := a.GetStocktakingDB(ctx, common.StringToInt64(req.Id), req.RefId, req.Type)

	if err != nil {
		logs.Error(ctx, "event=GetStocktaking error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetStocktakingResponse)
	data := make([]*freezonex_openiiot_api.Stocktaking, 0)

	var Names []string // 初始化一个空的字符串切片来存储id
	for _, b := range strings.Split(stocktakings[0].StorageLocationIds, ",") {

		storageLocationService := wms_storage_location.DefaultStorageLocationService()
		storageLocationData, _ := storageLocationService.GetStorageLocationDB(ctx, common.StringToInt64(b), "", nil, "")
		Names = append(Names, storageLocationData[0].Name) // 将每个Id添加到切片中
	}

	for _, v := range stocktakings {
		data = append(data, &freezonex_openiiot_api.Stocktaking{
			Id:                 common.Int64ToString(v.ID),
			RefId:              v.RefID,
			Type:               v.Type,
			StorageLocationIds: strings.Split(v.StorageLocationIds, ","),
			StorageLocations:   Names,
			Operator:           v.Operator,
			Result:             v.Result,
			CreateTime:         common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:         common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateStocktaking will update stocktaking record
func (a *WmsStocktakingService) UpdateStocktaking(ctx context.Context, req *freezonex_openiiot_api.UpdateStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateStocktakingResponse, error) {
	err := a.UpdateStocktakingDB(ctx, common.StringToInt64(req.Id), req.RefId, req.Type, req.Result)
	if err != nil {
		logs.Error(ctx, "event=UpdateStocktaking error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateStocktakingResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteStocktaking will delete stocktaking record
func (a *WmsStocktakingService) DeleteStocktaking(ctx context.Context, req *freezonex_openiiot_api.DeleteStocktakingRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteStocktakingResponse, error) {
	//Delete stocktaking also should delete stocktaking , edge pool, core pool, application pool, flow
	/*err := a.DeleteStocktakingStocktakingDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteStocktaking stocktaking error=%v", err.Error())
		return nil, err
	}*/

	// Delete stocktaking
	err := a.DeleteStocktakingDB(ctx, common.StringToInt64(req.Id))
	if err != nil {
		logs.Error(ctx, "event=DeleteStocktaking error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.DeleteStocktakingResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

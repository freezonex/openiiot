package grafana

import (
	"encoding/json"
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type GrafanaService struct {
	db     *mysql.MySQL
	c      *config.GrafanaConfig
	client *GrafanaClient
}

type OAuthRequest struct {
	Code        string `form:"code"`
	GrantType   string `form:"grant_type"`
	RedirectURI string `form:"redirect_uri"`
}

var (
	service *GrafanaService
	once    sync.Once
)

func NewGrafanaService(db *mysql.MySQL, c *config.GrafanaConfig) *GrafanaService {
	once.Do(func() {
		service = &GrafanaService{
			db: db,
			c:  c,
		}
	})
	return service
}

type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Slug      string `json:"slug"`
	Folder    int64  `json:"folderId"`
	FolderUID string `json:"folderUid"`
	URL       string `json:"url"`
}

type Dashboard struct {
	Model    map[string]interface{} `json:"dashboard"`
	FolderID int64                  `json:"folderId"`

	// This field is read-only. It is not used when creating a new dashboard.
	Meta DashboardMeta `json:"meta"`

	// These fields are only used when creating a new dashboard, they will always be empty when getting a dashboard.
	Overwrite bool   `json:"overwrite,omitempty"`
	Message   string `json:"message,omitempty"`
	FolderUID string `json:"folderUid,omitempty"`
}

type DashboardSaveResponse struct {
	Slug    string `json:"slug"`
	ID      int64  `json:"id"`
	UID     string `json:"uid"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}

type DataSource struct {
	ID   int64  `json:"id,omitempty"`
	UID  string `json:"uid,omitempty"`
	Name string `json:"name"`

	Type string `json:"type"`
	// This is only returned by the API. It depends on the Type.
	TypeLogoURL string `json:"typeLogoUrl,omitempty"`

	URL    string `json:"url"`
	Access string `json:"access"`

	// This is only returned by the API. It can only be set through the `editable` attribute of provisioned data sources.
	ReadOnly bool `json:"readOnly"`

	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`

	OrgID     int64 `json:"orgId,omitempty"`
	IsDefault bool  `json:"isDefault"`

	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser,omitempty"`

	WithCredentials bool `json:"withCredentials,omitempty"`

	JSONData       map[string]interface{} `json:"jsonData,omitempty"`
	SecureJSONData map[string]interface{} `json:"secureJsonData,omitempty"`

	Version int `json:"version,omitempty"`
}

func convertPbDataSource(ds *freezonex_openiiot_api.DataSource) (*DataSource, error) {

	var jsonDataJSON map[string]interface{}
	err := json.Unmarshal([]byte(ds.JsonData), &jsonDataJSON)
	if err != nil {
		logs.Error("failed to unmarshal JsonData in datasource: %v", err, ds.JsonData)
		return nil, err
	}

	var secureJsonDataJSON map[string]interface{}
	err = json.Unmarshal([]byte(ds.SecureJsonData), &secureJsonDataJSON)
	if err != nil {
		logs.Error("failed to unmarshal SecureJsonData in datasource: %v", err, ds.SecureJsonData)
		return nil, err
	}
	newDataSource := &DataSource{
		ID:              ds.Id,
		UID:             ds.Uid,
		Name:            ds.Name,
		Type:            ds.Type,
		TypeLogoURL:     ds.TypeLogoUrl,
		URL:             ds.Url,
		Access:          ds.Access,
		ReadOnly:        ds.ReadOnly,
		Database:        ds.Database,
		User:            ds.User,
		OrgID:           ds.OrgId,
		IsDefault:       ds.IsDefault,
		BasicAuth:       ds.BasicAuth,
		BasicAuthUser:   ds.BasicAuthUser,
		WithCredentials: ds.WithCredentials,
		JSONData:        jsonDataJSON,
		SecureJSONData:  secureJsonDataJSON,
		Version:         int(ds.Version),
	}
	return newDataSource, nil
}

func convertDataSources(dataSources []*DataSource) ([]*freezonex_openiiot_api.DataSource, error) {
	pbDataSources := make([]*freezonex_openiiot_api.DataSource, len(dataSources))

	for i, ds := range dataSources {
		jsonDataBytes, err := json.Marshal(ds.JSONData)
		if err != nil {
			return nil, err
		}

		secureJsonDataBytes, err := json.Marshal(ds.SecureJSONData)
		if err != nil {
			return nil, err
		}

		pbDataSources[i] = &freezonex_openiiot_api.DataSource{
			Id:              ds.ID,
			Uid:             ds.UID,
			Name:            ds.Name,
			Type:            ds.Type,
			TypeLogoUrl:     ds.TypeLogoURL,
			Url:             ds.URL,
			Access:          ds.Access,
			ReadOnly:        ds.ReadOnly,
			Database:        ds.Database,
			User:            ds.User,
			OrgId:           ds.OrgID,
			IsDefault:       ds.IsDefault,
			BasicAuth:       ds.BasicAuth,
			BasicAuthUser:   ds.BasicAuthUser,
			WithCredentials: ds.WithCredentials,
			JsonData:        string(jsonDataBytes),
			SecureJsonData:  string(secureJsonDataBytes),
			Version:         int32(ds.Version),
		}
	}

	return pbDataSources, nil
}

func convertDashboardModel(dashboard *freezonex_openiiot_api.Dashboard) (*Dashboard, error) {
	var modelJSON map[string]interface{}
	err := json.Unmarshal([]byte(dashboard.Model), &modelJSON)
	if err != nil {
		logs.Error("failed to unmarshal dashboard model: %v", err, dashboard.Model)
		return nil, err
	}
	newDashboard := &Dashboard{
		Model:    modelJSON,
		FolderID: dashboard.FolderId,
		Meta: DashboardMeta{
			IsStarred: dashboard.Meta.IsStarred,
			Slug:      dashboard.Meta.Slug,
			Folder:    dashboard.Meta.FolderId,
			FolderUID: dashboard.Meta.FolderUid,
			URL:       dashboard.Meta.Url,
		},
		Overwrite: dashboard.Overwrite,
		Message:   dashboard.Message,
		FolderUID: dashboard.FolderUid,
	}

	return newDashboard, nil
}

func DefaultGrafanaService() *GrafanaService {
	return service
}

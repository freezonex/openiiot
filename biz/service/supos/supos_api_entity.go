package supos

type SuposApiEntity struct {
	Name   string
	URL    string
	Method string
}

// Define the methods as constants
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

// Define the API endpoints similar to the enum values
var (
	OAUTH2_GET_AUTHORIZE_CODE = SuposApiEntity{
		Name:   "获取授权码",
		URL:    "/inter-api/auth/v1/oauth2/authorize",
		Method: MethodGet,
	}
	OAUTH2_GET_ACCESS_TOKEN = SuposApiEntity{
		Name:   "获取AccessToken",
		URL:    "/open-api/auth/v2/oauth2/token",
		Method: MethodPost,
	}
	// Add other endpoints as needed
	IGNORE = SuposApiEntity{
		Name:   "IGNORE",
		URL:    "IGNORE",
		Method: MethodPut,
	}

	//用户接口
	USER_ADD = SuposApiEntity{
		Name:   "用户新增",
		URL:    "/open-api/auth/v2/users",
		Method: MethodPost,
	}
	USER_LISTS = SuposApiEntity{
		Name:   "分页获取用户列表",
		URL:    "/open-api/auth/v2/users",
		Method: MethodGet,
	}
	USER_BATCH_DELETE = SuposApiEntity{
		Name:   "批量删除用户",
		URL:    "/open-api/supos/auth/v2/delete/users",
		Method: MethodPost,
	}
	USER_DETAIL = SuposApiEntity{
		Name:   "用户详情",
		URL:    "/open-api/auth/v2/users/%s",
		Method: MethodGet,
	}
	USER_UPDATE = SuposApiEntity{
		Name:   "修改用户",
		URL:    "/open-api/auth/v2/users/%s",
		Method: MethodPut,
	}
	USER_BIND_ROLE = SuposApiEntity{
		Name:   "用户绑定角色",
		URL:    "/open-api/auth/v2/users/%s/role",
		Method: MethodPost,
	}
	USER_UNBIND_ROLE = SuposApiEntity{
		Name:   "用户解除绑定角色",
		URL:    "/open-api/auth/v2/users/%s/role",
		Method: MethodPut,
	}
)

// Function to update the URL of an API endpoint if needed
func (e *SuposApiEntity) SetURL(newURL string) {
	e.URL = newURL
}

// Function to get the URL of an API endpoint
func (e SuposApiEntity) GetURL() string {
	return e.URL
}

// Function to get the HTTP method of an API endpoint
func (e SuposApiEntity) GetMethod() string {
	return e.Method
}

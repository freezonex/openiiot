package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hashicorp/go-cleanhttp"
)

// Client is a Grafana API client.
type GrafanaClient struct {
	config  Config
	baseURL url.URL
	client  *http.Client
}

type Config struct {
	// APIKey is an optional API key or service account token.
	APIKey string
	// BasicAuth is optional basic auth credentials.
	BasicAuth *url.Userinfo
	// HTTPHeaders are optional HTTP headers.
	HTTPHeaders map[string]string
	// Client provides an optional HTTP client, otherwise a default will be used.
	Client *http.Client
	// OrgID provides an optional organization ID
	// with BasicAuth, it defaults to last used org
	// with APIKey, it is disallowed because service account tokens are scoped to a single org
	OrgID int64
	// NumRetries contains the number of attempted retries
	NumRetries int
	// RetryTimeout says how long to wait before retrying a request
	RetryTimeout time.Duration
	// RetryStatusCodes contains the list of status codes to retry, use "x" as a wildcard for a single digit (default: [429, 5xx])
	RetryStatusCodes []string
}

// New creates a new Grafana client.
func New(baseURL string, cfg Config) (*GrafanaClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if cfg.BasicAuth != nil {
		u.User = cfg.BasicAuth
	}

	cli := cfg.Client
	if cli == nil {
		cli = cleanhttp.DefaultClient()
	}

	return &GrafanaClient{
		config:  cfg,
		baseURL: *u,
		client:  cli,
	}, nil
}

// WithOrgID returns a new client with the provided organization ID.
func (c GrafanaClient) WithOrgID(orgID int64) *GrafanaClient {
	c.config.OrgID = orgID
	return &c
}

func (c *GrafanaClient) request(method, requestPath string, query url.Values, body []byte, responseStruct interface{}) error {
	var (
		req          *http.Request
		resp         *http.Response
		err          error
		bodyContents []byte
	)
	retryStatusCodes := c.config.RetryStatusCodes
	if len(retryStatusCodes) == 0 {
		retryStatusCodes = []string{"429", "5xx"}
	}

	// retry logic
	for n := 0; n <= c.config.NumRetries; n++ {
		req, err = c.newRequest(method, requestPath, query, bytes.NewReader(body))
		if err != nil {
			return err
		}

		// Wait a bit if that's not the first request
		if n != 0 {
			if c.config.RetryTimeout == 0 {
				c.config.RetryTimeout = time.Second * 5
			}
			time.Sleep(c.config.RetryTimeout)
		}

		resp, err = c.client.Do(req)

		// If err is not nil, retry again
		// That's either caused by client policy, or failure to speak HTTP (such as network connectivity problem). A
		// non-2xx status code doesn't cause an error.
		if err != nil {
			continue
		}

		// read the body (even on non-successful HTTP status codes), as that's what the unit tests expect
		bodyContents, err = io.ReadAll(resp.Body)
		resp.Body.Close() //nolint:errcheck

		// if there was an error reading the body, try again
		if err != nil {
			continue
		}

		shouldRetry, err := matchRetryCode(resp.StatusCode, retryStatusCodes)
		if err != nil {
			return err
		}
		if !shouldRetry {
			break
		}
	}
	if err != nil {
		return err
	}

	if os.Getenv("GF_LOG") != "" {
		log.Printf("response status %d with body %v", resp.StatusCode, string(bodyContents))
	}

	// check status code.
	switch {
	case resp.StatusCode == http.StatusNotFound:
		return ErrNotFound{
			BodyContents: bodyContents,
		}
	case resp.StatusCode >= 400:
		return fmt.Errorf("status: %d, body: %v", resp.StatusCode, string(bodyContents))
	}

	if responseStruct == nil {
		return nil
	}

	err = json.Unmarshal(bodyContents, responseStruct)
	if err != nil {
		return err
	}

	return nil
}

func (c *GrafanaClient) newRequest(method, requestPath string, query url.Values, body io.Reader) (*http.Request, error) {
	url := c.baseURL
	url.Path = path.Join(url.Path, requestPath)
	url.RawQuery = query.Encode()
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return req, err
	}

	// cannot use both API key and org ID. API keys are scoped to single org
	if c.config.APIKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	}
	if c.config.OrgID != 0 {
		req.Header.Add("X-Grafana-Org-Id", strconv.FormatInt(c.config.OrgID, 10))
	}

	if c.config.HTTPHeaders != nil {
		for k, v := range c.config.HTTPHeaders {
			req.Header.Add(k, v)
		}
	}

	if os.Getenv("GF_LOG") != "" {
		if body == nil {
			log.Printf("request (%s) to %s with no body data", method, url.String())
		} else {
			reader := body.(*bytes.Reader)
			if reader.Len() == 0 {
				log.Printf("request (%s) to %s with no body data", method, url.String())
			} else {
				contents := make([]byte, reader.Len())
				if _, err := reader.Read(contents); err != nil {
					return nil, fmt.Errorf("cannot read body contents for logging: %w", err)
				}
				if _, err := reader.Seek(0, io.SeekStart); err != nil {
					return nil, fmt.Errorf("failed to seek body reader to start after logging: %w", err)
				}
				log.Printf("request (%s) to %s with body data: %s", method, url.String(), string(contents))
			}
		}
	}

	req.Header.Add("Content-Type", "application/json")
	return req, err
}

// matchRetryCode checks if the status code matches any of the configured retry status codes.
func matchRetryCode(gottenCode int, retryCodes []string) (bool, error) {
	gottenCodeStr := strconv.Itoa(gottenCode)
	for _, retryCode := range retryCodes {
		if len(retryCode) != 3 {
			return false, fmt.Errorf("invalid retry status code: %s", retryCode)
		}
		matched := true
		for i := range retryCode {
			c := retryCode[i]
			if c == 'x' {
				continue
			}
			if gottenCodeStr[i] != c {
				matched = false
				break
			}
		}
		if matched {
			return true, nil
		}
	}

	return false, nil
}
// // NewDataSource creates a new Grafana data source.
// func (c *GrafanaClient) NewDataSource(s *freezonex_openiiot_api.DataSource) (int64, error) {
// 	data, err := json.Marshal(s)
// 	if err != nil {
// 		return 0, err
// 	}

// 	result := struct {
// 		ID int64 `json:"id"`
// 	}{}

// 	err = c.request("POST", "/api/datasources", nil, data, &result)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.ID, err
// }


func (c *GrafanaClient) DataSources(DSN *freezonex_openiiot_api.GrafanaDSN) ([]*freezonex_openiiot_api.DataSource, error) {
	result := make([]*DataSource, 0)
	config := Config{
		BasicAuth: url.UserPassword(DSN.Username, DSN.Password),
		HTTPHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		Client: http.DefaultClient,
	}

	cli, err := New(DSN.Host, config)
	if err != nil {
		return nil, err
	}
	reqErr := cli.request("GET", "/api/datasources", nil, nil, &result)
	if reqErr != nil {
		return nil, reqErr
	}

	formattedResult, err := convertDataSources(result)
	if err != nil {
        log.Fatalf("Error converting data sources: %v", err)
    }
	
	return formattedResult, nil
}

func (c *GrafanaClient) CreateDashBoard(DSN *freezonex_openiiot_api.GrafanaDSN, dashboard *freezonex_openiiot_api.Dashboard) (*freezonex_openiiot_api.DashboardSaveResponse, error) {
	db, err := convertDashboardModel(dashboard)
	if err != nil {
		logs.Error("Cannot unmarshal model in dashboard", dashboard)
		return nil, err
	}
	data, err := json.Marshal(db)
	if err != nil {
		logs.Error("Cannot marshal dashboard", dashboard)
		return nil, err
	}

	config := Config{
		BasicAuth: url.UserPassword(DSN.Username, DSN.Password),
		HTTPHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		Client: http.DefaultClient,
	}

	cli, err := New(DSN.Host, config)
	if err != nil {
		return nil, err
	}
	result := &freezonex_openiiot_api.DashboardSaveResponse{}
	err = cli.request("POST", "/api/dashboards/db", nil, data, &result)
	if err != nil {
		return nil, err
	}
	logs.Debug("Client level debug", result)
	return result, err
}

func (c *GrafanaClient) dashboard(path string) (*Dashboard, error) {
	result := &Dashboard{}
	err := c.request("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	result.FolderID = result.Meta.Folder

	return result, err
}

func (c *GrafanaClient) DashboardByUID(DSN *freezonex_openiiot_api.GrafanaDSN, uid string, path string) (*Dashboard, error) {
	config := Config{
		BasicAuth: url.UserPassword(DSN.Username, DSN.Password),
		HTTPHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		Client: http.DefaultClient,
	}

	cli, err := New(DSN.Host, config)
	if err != nil {
		return nil, err
	}
	return cli.dashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

func (c *GrafanaClient) NewDataSource(DSN *freezonex_openiiot_api.GrafanaDSN, s *freezonex_openiiot_api.DataSource) (int64, error) {
	convertedData, err:= convertPbDataSource(s)
	if err != nil {
		return 0, err
	}
	data, err := json.Marshal(convertedData)
	if err != nil {
		return 0, err
	}

	config := Config{
		BasicAuth: url.UserPassword(DSN.Username, DSN.Password),
		HTTPHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		Client: http.DefaultClient,
	}

	cli, err := New(DSN.Host, config)
	if err != nil {
		return 0, err
	}
	result := struct {
		ID int64 `json:"id"`
	}{}

	err = cli.request("POST", "/api/datasources", nil, data, &result)
	if err != nil {
		return 0, err
	}

	return result.ID, err
}

func (c *GrafanaClient) DeleteDataSourceByName(DSN *freezonex_openiiot_api.GrafanaDSN, name string) error {
	path := fmt.Sprintf("/api/datasources/name/%s", name)
	config := Config{
		BasicAuth: url.UserPassword(DSN.Username, DSN.Password),
		HTTPHeaders: map[string]string{
			"Content-Type": "application/json",
		},
		Client: http.DefaultClient,
	}

	cli, err := New(DSN.Host, config)
	if err != nil {
		return err
	}
	return cli.request("DELETE", path, nil, nil, nil)
}

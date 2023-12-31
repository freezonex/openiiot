package consts

import "errors"

// general settings
const (
	TOKEN_VALIDATED_CACHE = 1000 //items
	TENANT_RETURN_LIMIT   = 1000 //DB rows
)

// account status
const (
	ACCOUNT_STATUS_LOCK  = "LOCKED"
	ACCOUNT_STATUS_SHARE = "SHARE"
	ACCOUNT_STATUS_FREE  = "FREE"
)

const (
	// flow outcomes
	FAILED  = "failed"
	PASSED  = "passed"
	SKIPPED = "skipped"
	ERROR   = "error"
)

// AdminConfig
const (
	ADMINCONFIG_SYNC_INTERVAL = 1 // minute
	AdminConfigAdministrators = "administrators"
	AdminConfigPublicConfigs  = "system_public"
)

// openiiot flow status
const (
	FLOW_DISABLE = "DISABLED"
	FLOW_ENABLE  = "ENABLE"
)

var ErrKeyNotFound = errors.New("key not found")
var ErrUnexpectedItemCount = errors.New("item count isn't as expected")
var ErrInvalidNodeID = errors.New("invalid node id")
var ErrExternalAPIError = errors.New("external api returned error status code")
var ErrBadRequest = errors.New("request body missing required fields")
var ErrWrongInputFormat = errors.New("wrong input format")
var ErrIndexFileNotFound = errors.New("index.html file not found in allure html report tar file")
var ErrTimeout = errors.New("request timed out")

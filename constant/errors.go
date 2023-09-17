package constant

import "errors"

var (
	ErrFailedJSONMarshal = errors.New("JSON Marshal is failed")
	ErrConfig            = errors.New("Failed to Load env file")
	ErrConfigPath        = errors.New("Failed to get directory of env file")
	ErrServe             = errors.New("Failed to serve")
	ErrRoutes            = errors.New("Failed to run routes")
)

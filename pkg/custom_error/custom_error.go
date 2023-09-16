package custom_error

import (
	"encoding/json"
	"errors"

	"github.com/rizqyfahmi/gin-greetings-clean-architecture/constant"
)

type CustomError struct {
	display error
	plain   error
	path    string
}

func NewCustomError(
	display error,
	plain error,
	path string,
) error {
	return &CustomError{
		display: display,
		plain:   plain,
		path:    path,
	}
}

func (e *CustomError) Error() string {
	message := map[string]interface{}{
		"display": e.display.Error(),
		"plain":   e.plain.Error(),
		"path":    e.path,
	}

	result, err := json.Marshal(message)
	if err != nil {
		return err.Error()
	}

	return string(result)
}

func (e *CustomError) GetDisplay() error {
	return e.display
}

func (e *CustomError) GetPlain() error {
	return e.plain
}

func (e *CustomError) GetPath() string {
	return e.path
}

func (e *CustomError) UnshiftPath(path string) error {
	e.path = path + " > " + e.path
	return e
}

func (e *CustomError) FromListMap(errs []map[string]interface{}) error {
	path := "CustomError:FromListMap"
	result, err := json.Marshal(errs)
	if err != nil {
		return NewCustomError(
			constant.ErrFailedJSONMarshal,
			err,
			path,
		)
	}

	e.plain = errors.New(string(result))
	return e
}

package v2

import (
	"fmt"
	"time"
)

var (
	DefaultStatusChangeDelay = 2 * time.Minute
)

func (m *Params) Validate() error {
	if err := validateStatusChangeDelay(m.StatusChangeDelay); err != nil {
		return err
	}

	return nil
}

func NewParams(statusChangeDelay time.Duration) Params {
	return Params{
		StatusChangeDelay: statusChangeDelay,
	}
}

func DefaultParams() Params {
	return Params{
		StatusChangeDelay: DefaultStatusChangeDelay,
	}
}

func validateStatusChangeDelay(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("status_change_delay cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("status_change_delay cannot be zero")
	}

	return nil
}

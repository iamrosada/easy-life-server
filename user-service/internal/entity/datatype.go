package entity

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

func (o *TeachersIDs) Scan(src interface{}) error {
	if src == nil {
		*o = nil
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("expected string, got %T", src)
	}

	*o = strings.Split(str, ",")
	return nil
}

func (o TeachersIDs) Value() (driver.Value, error) {

	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}

package db

import "strings"

func IsUniqueConstraintError(err error) bool {
	if err != nil && strings.Contains(err.Error(), "duplicate key value") {
		return true
	}

	return false
}

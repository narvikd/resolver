//go:build windows

package resolver

import (
	"errors"
)

// IsHostAlive is not implemented on Windows
func IsHostAlive() error {
	return errors.New("IsHostAlive not implemented on windows")
}

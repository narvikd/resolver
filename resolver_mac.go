//go:build darwin

package resolver

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

func IsHostAlive(host string, timeout time.Duration) bool {
	var out bytes.Buffer
	if !IsHostResolvable(host, timeout) {
		return false
	}
	cmd := exec.Command("timeout", "0.3", "ping", "-c", "1", host)
	cmd.Stdout = &out
	isAlive := cmd.Run() == nil || strings.Contains(out.String(), "bytes from")
	return isAlive
}

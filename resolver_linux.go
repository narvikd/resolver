//go:build linux

package resolver

import (
	pinger "github.com/prometheus-community/pro-bing"
	"time"
)

func IsHostAlive(host string, timeout time.Duration) bool {
	if !IsHostResolvable(host, timeout) {
		return false
	}
	p, err := pinger.NewPinger(host)
	if err != nil {
		return false
	}
	p.Count = 1
	p.Interval = timeout
	p.Timeout = timeout

	if p.Run() != nil {
		return false
	}
	return p.Statistics().PacketsRecv >= 1
}

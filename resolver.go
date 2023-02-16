package resolver

import (
	"net"
	"time"
)

func IsHostResolvable(host string, timeout time.Duration) bool {
	t := time.After(timeout)
	result := make(chan error)
	go func() {
		_, err := net.LookupHost(host)
		result <- err
	}()
	select {
	case <-t:
		return false
	case err := <-result:
		return err == nil
	}
}

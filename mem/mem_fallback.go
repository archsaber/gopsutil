// +build !darwin,!linux,!freebsd,!openbsd,!windows

package mem

import "github.com/archsaber/gopsutil/internal/common"

func VirtualMemory() (*VirtualMemoryStat, error) {
	return nil, common.ErrNotImplementedError
}

func SwapMemory() (*SwapMemoryStat, error) {
	return nil, common.ErrNotImplementedError
}

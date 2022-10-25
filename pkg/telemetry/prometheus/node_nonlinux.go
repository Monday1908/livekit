//go:build !linux

package prometheus

import (
	"runtime"
	"sync"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

var (
	cpuStatsLock              sync.RWMutex
	lastCPUTotal, lastCPUIdle uint64
)

func getCPUStats() (cpuLoad float32, numCPUs uint32, err error) {
	cpuInfo, err := cpu.Get()
	if err != nil {
		return
	}

	cpuStatsLock.Lock()
	if lastCPUTotal > 0 && lastCPUTotal < cpuInfo.Total {
		cpuLoad = 1 - float32(cpuInfo.Idle-lastCPUIdle)/float32(cpuInfo.Total-lastCPUTotal)
	}

	lastCPUTotal = cpuInfo.Total
	lastCPUIdle = cpuInfo.Idle
	cpuStatsLock.Unlock()

	numCPUs = uint32(runtime.NumCPU())

	return
}

func getMemoryStats() (memoryLoad float32, err error) {
	memInfo, err := memory.Get()
	if err != nil {
		return
	}

	memoryLoad = float32(memInfo.Used) / float32(memInfo.Total)
	return
}

func getTCStats() (packets, drops uint32, err error) {
	// linux only
	return
}

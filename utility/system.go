//go:build (darwin && arm64) || linux || windows
// +build darwin,arm64 linux windows

package utility

import (
	"fmt"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type MemoryInfo struct {
	Percent, Used, Free, Total string
}

func GetMemoryInfo() (info MemoryInfo, err error) {
	info = MemoryInfo{}
	memoryInfo, err := memory.Get()
	if err != nil {
		return
	}
	info.Percent = fmt.Sprintf("%.2f%%", float64(memoryInfo.Used)/float64(memoryInfo.Total)*100)
	info.Used = fmt.Sprintf("%.2fGB", float64(memoryInfo.Used)/(1<<10<<10<<10))
	info.Free = fmt.Sprintf("%.2fGB", float64(memoryInfo.Total-memoryInfo.Used)/(1<<10<<10<<10))
	info.Total = fmt.Sprintf("%.2fGB", float64(memoryInfo.Total)/(1<<10<<10<<10))
	return
}

type CpuInfo struct {
	SystemUse, UserUse, Idle, TotalUse string
}

func GetCpuInfo() (info CpuInfo, err error) {
	info = CpuInfo{}

	before, err := cpu.Get()
	if err != nil {
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return
	}
	total := float64(after.Total - before.Total)

	info.UserUse = fmt.Sprintf("%.2f%%", float64(after.User-before.User)/total*100)
	info.SystemUse = fmt.Sprintf("%.2f%%", float64(after.System-before.System)/total*100)
	idle := float64(after.Idle-before.Idle) / total * 100
	info.Idle = fmt.Sprintf("%.2f%%", idle)
	info.TotalUse = fmt.Sprintf("%.2f%%", float64(100-idle))

	return
}

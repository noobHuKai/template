package system

import (
	"github.com/noobHuKai/app/g"
	"github.com/shirou/gopsutil/cpu"
)

type SysBasicInfo struct {
	Cpu    SysCpuInfo    `json:"cpu"`
	Memory SysMemoryInfo `json:"memory"`
	Disk   SysDiskInfo   `json:"disk"`
}

func (s SysBasicInfo) Bytes() ([]byte, error) {
	return g.JsonIter.Marshal(s)
}

type SysCpuInfo struct {
	PhysicalCnt int           `json:"physicalCnt"`
	LogicalCnt  int           `json:"logicalCnt"`
	Percent     float64       `json:"percent"`
	TimeState   cpu.TimesStat `json:"timeState"`
}
type SysDiskInfo struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type SysMemoryInfo struct {
	VirtualMemoryInfo SysMemoryBasicInfo `json:"virtual"`
	SwapMemoryInfo    SysMemoryBasicInfo `json:"swap"`
}
type SysMemoryBasicInfo struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

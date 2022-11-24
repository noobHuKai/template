package system

import (
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/system"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strings"
	"time"
)

// BasicService system basic info
type BasicService struct {
}

func (b *BasicService) GetBasicInfo() system.SysBasicInfo {
	var sysBasicInfo system.SysBasicInfo
	memoryInfo := b.GetMemoryMonitorInfo()
	cpuInfo := b.GetCpuBasicInfo()
	diskInfo := b.GetDiskBasicInfo()
	sysBasicInfo.Cpu = cpuInfo
	sysBasicInfo.Memory = memoryInfo
	sysBasicInfo.Disk = diskInfo
	
	return sysBasicInfo
}

func (b *BasicService) GetMonitorInfo() system.SysBasicInfo {
	var sysBasicInfo system.SysBasicInfo
	memoryInfo := b.GetMemoryMonitorInfo()
	cpuInfo := b.GetCpuMonitorInfo()
	sysBasicInfo.Cpu = cpuInfo
	sysBasicInfo.Memory = memoryInfo

	return sysBasicInfo
}

func (b *BasicService) GetMemoryMonitorInfo() system.SysMemoryInfo {
	var sysMemoryInfo system.SysMemoryInfo
	if virtualMemoryInfo, err := mem.VirtualMemory(); err == nil {
		sysMemoryInfo.VirtualMemoryInfo = system.SysMemoryBasicInfo{
			Total:       virtualMemoryInfo.Total,
			Free:        virtualMemoryInfo.Free,
			Used:        virtualMemoryInfo.Used,
			UsedPercent: virtualMemoryInfo.UsedPercent,
		}
	} else {
		g.Logger.Error(err.Error())
	}

	if swapMemoryInfo, err := mem.SwapMemory(); err == nil {
		sysMemoryInfo.SwapMemoryInfo = system.SysMemoryBasicInfo{
			Total:       swapMemoryInfo.Total,
			Free:        swapMemoryInfo.Free,
			Used:        swapMemoryInfo.Used,
			UsedPercent: swapMemoryInfo.UsedPercent,
		}
	} else {
		g.Logger.Error(err.Error())
	}
	return sysMemoryInfo
}

func (b *BasicService) GetCpuMonitorInfo() system.SysCpuInfo {
	var cpuInfo system.SysCpuInfo
	if percent, err := cpu.Percent(0, false); err == nil {
		cpuInfo.Percent = percent[0]
	} else {
		g.Logger.Error(err.Error())
	}
	return cpuInfo
}
func (b *BasicService) GetCpuBasicInfo() system.SysCpuInfo {
	var cpuInfo system.SysCpuInfo
	if percent, err := cpu.Percent(time.Second, false); err == nil {
		cpuInfo.Percent = percent[0]
	} else {
		g.Logger.Error(err.Error())
	}
	if physicalCnt, err := cpu.Counts(false); err == nil { // cpu物理核数
		cpuInfo.PhysicalCnt = physicalCnt
	} else {
		g.Logger.Error(err.Error())
	}
	if logicalCnt, err := cpu.Counts(true); err == nil { // cpu逻辑核数
		cpuInfo.LogicalCnt = logicalCnt
	} else {
		g.Logger.Error(err.Error())
	}
	if timesStats, err := cpu.Times(false); err == nil {
		cpuInfo.TimeState = timesStats[0]
	} else {
		g.Logger.Error(err.Error())
	}

	return cpuInfo
}

func (b *BasicService) GetDiskBasicInfo() system.SysDiskInfo {
	var diskInfo system.SysDiskInfo
	infos, _ := disk.Partitions(false)
	for _, info := range infos {
		if strings.Contains(info.Opts, "nodev") {
			continue
		}
		usageStat, err := disk.Usage(info.Mountpoint)
		if err != nil {
			continue
		}
		diskInfo.Total += usageStat.Total
		diskInfo.Used += usageStat.Used
		diskInfo.Free += usageStat.Free
	}
	diskInfo.UsedPercent = float64(diskInfo.Used) / float64(diskInfo.Total) * 100
	return diskInfo
}

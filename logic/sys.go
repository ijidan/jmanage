package logic

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"time"
)

//系统信息
type Sys struct {
}

//获取CPU信息
func (l *Sys) GetCPUInfo() []cpu.InfoStat {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil
	}
	return cpuInfo
}

//获取CPU使用率
func (l *Sys) GetCPUPercent() interface{} {
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil
	}
	return percent
}

//获取CPU个数
func (l *Sys) GetCPUNum() int {
	num, err := cpu.Counts(true)
	if err != nil {
		return 0
	}
	return num
}

//获取HOST INFO
func (l *Sys) GetHostInfo() *host.InfoStat {
	hostInfo, err := host.Info()
	if err != nil {
		return nil
	}
	return hostInfo
}

//获取docker状态
func (l *Sys) GetDockerStat() []docker.CgroupDockerStat {
	dockerStat, err := docker.GetDockerStat()
	if err != nil {
		return nil
	}
	return dockerStat
}

//获取docker ID 列表
func (l *Sys) GetDockerIDList() []string {
	dockerIdList, err := docker.GetDockerIDList()
	if err != nil {
		return nil
	}
	return dockerIdList
}

//获取系统信息
func (l *Sys) GetSysInfo() map[string]interface{} {
	cpuInfo := l.GetCPUInfo()
	dockerStat := l.GetDockerStat()
	data := map[string]interface{}{
		"cpuInfo":      gconv.SliceAny(cpuInfo),
		"cpuNum":       l.GetCPUNum(),
		"cpuPercent":   l.GetCPUPercent(),
		"hostInfo":     l.GetHostInfo(),
		"dockerStat":   gconv.SliceAny(dockerStat),
		"dockerIDList": l.GetDockerIDList(),
	}
	return data
}

//获取实例
func NewSys() *Sys {
	return &Sys{}
}

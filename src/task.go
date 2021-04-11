package main

import (
	"net"
)

// 数据结构，部分
// 业务逻辑代码，注释多，更好一些，如果是平台代码，则无需太多注释
// status machine // 增删改查，查是遍历，不需要
// 销毁任务，只需要将标志位清空，具体字段可以尝试不处理

// 先写一个文件，然后进行拆分

type Task interface {
	Create()  // 注册任务，增，由receive msgbox进行收发消息
	Destroy() // 销毁任务，删，处理任务后，不直接销毁，统一发送销毁消息给执行侧，在下一个时间片进行销毁，直接在handle里面处理，
	Handle()  // 处理任务，改，，调度到后，统一处理
	Report()  // 上报任务，这里可以尝试放到处理任务的一个子函数里面，理由是无需对结构体进行读取操作
}

// some are regInfo, some are runInfo, some are resultInfo
type PingTask struct {
	DstIP           net.IP
	VpnName         string
	PktNum          uint32
	PktLen          uint32
	SendInterval    uint32
	TimeoutInterval uint32

	currentSeq uint32
	rcvRes     uint32
	timeoutRes uint32
}

type TraceRouteTask struct {
	// reg
	DstIP    net.IP
	VpnName  string
	FirstTTL uint32
	MaxTTL   uint32

	// run and res
	currentTTL   uint32
	currentDstIP net.IP
}

func mai3n() {

}

type resource struct {
}
type dispather struct {
	ID   uint32
	name string
	res  resource
	grid
}

// 通过channel互传
type grid interface {
	New()
	Receive()
	Stop()
	Start()
}

// trans
func executor() {

}

// ctrl
func controller() {

}

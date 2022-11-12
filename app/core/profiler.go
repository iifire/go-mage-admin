package core

import "time"

// Profiler 分析器
type Profiler struct {
	//计数器
	timers []interface{}
}

// Start 开始断点
func (p *Profiler) Start(code string) {
	p.add("start", code)
	return
}

// Stop 结束断点
func (p *Profiler) Stop(code string) {
	p.add("stop", code)
	return
}

// add 添加记录器
func (p *Profiler) add(t string, code string) {
	curTime := time.Now().UnixNano()
	timer := map[string]interface{}{}
	timer["code"] = code
	timer["time"] = curTime
	timer["type"] = t
	p.timers = append(p.timers, timer)
	return
}

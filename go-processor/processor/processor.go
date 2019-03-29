// Package processor processor define
// Created by chenguolin 2019-01-13
package processor

import (
	"runtime/debug"

	"github.com/remeh/sizedwaitgroup"
	golog "gitlab.local.com/golang/go-log"
)

// processor state
const (
	running = iota
	stopping
	stopped
)

// Processor define
type Processor struct {
	state         int                           //processor state
	scanner       Scanner                       //processor scanner
	handle        Handle                        //processor handle
	concurrentCnt int                           //processor concurrent count
	waitGroup     sizedwaitgroup.SizedWaitGroup //waitGroup
}

// NewProcessor new processor
func NewProcessor(scanner Scanner, handle Handle, concurrentCnt int) *Processor {
	return &Processor{
		state:         stopped,
		scanner:       scanner,
		handle:        handle,
		concurrentCnt: concurrentCnt,
		waitGroup:     sizedwaitgroup.New(concurrentCnt),
	}
}

// Start processor
func (p *Processor) Start() {
	// already running
	if p.state == running {
		return
	}

	// set running status
	p.state = running
	// start scanner
	p.scanner.Start()

	// start concurrent process
	// use sized wait group
	for i := 0; i < p.concurrentCnt; i++ {
		go p.process()
	}

	// wait all goroutine done
	go func() {
		p.waitGroup.Wait()
		// set 2 stopped
		p.state = stopped
	}()
}

// do process
func (p *Processor) process() {
	// add goroutine 2 waitGroup
	p.waitGroup.Add()

	defer func() {
		if err := recover(); err != nil {
			golog.Error("Processor process panic",
				golog.Object("Error", err))
			debug.PrintStack()
		}
		// goroutine done
		p.waitGroup.Done()
	}()

	for {
		// get next record
		record, ok := p.scanner.Next()

		// 需要安全退出
		// 如果ok为false 并且 scanner状态为关闭 则直接退出
		// 如果ok为true 说明channel还有数据 需要继续处理
		if !ok && p.scanner.IsStopped() {
			break
		}

		// start goroutine call handle process
		go func(record Record) {
			defer func() {
				if err := recover(); err != nil {
					golog.Error("Handle process panic",
						golog.Object("Error", err))
					debug.PrintStack()
				}
			}()
			p.handle.Process(record)
		}(record)
	}
}

// Stop processor
func (p *Processor) Stop() {
	if p.state != running {
		return
	}

	// set running false
	p.state = stopping
	// stop scanner
	p.scanner.Stop()
}

// isRunning check processor is running
func (p *Processor) isRunning() bool {
	return p.state == running
}

// isStopping check processor is stopping
func (p *Processor) isStopping() bool {
	return p.state == stopping
}

// isStopped check processor is stopped
func (p *Processor) isStopped() bool {
	return p.state == stopped
}

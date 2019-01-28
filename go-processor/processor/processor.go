// Package processor processor define
// Created by chenguolin 2019-01-13
package processor

import (
	"runtime/debug"

	"github.com/remeh/sizedwaitgroup"
	golog "gitlab.local.com/golang/go-log"
)

// Processor define
type Processor struct {
	isRunning     bool    //processor is running
	scanner       Scanner //processor scanner
	handle        Handle  //processor handle
	concurrentCnt int     //processor concurrent count
}

// NewProcessor new processor
func NewProcessor(scanner Scanner, handle Handle, concurrentCnt int) *Processor {
	return &Processor{
		isRunning:     false,
		scanner:       scanner,
		handle:        handle,
		concurrentCnt: concurrentCnt,
	}
}

// Start processor
func (p *Processor) Start() {
	// already running
	if p.isRunning {
		return
	}

	// set running status
	p.isRunning = true
	// start scanner
	p.scanner.Start()

	// start process
	go func() {
		defer func() {
			if err := recover(); err != nil {
				golog.Error("Processor start panic",
					golog.Object("Error", err))
				debug.PrintStack()
				// stop processor
				p.Stop()
			}
		}()

		// sized wait group
		swg := sizedwaitgroup.New(p.concurrentCnt)
		for {
			// if processor or scanner stop return straight
			if !p.isRunning || !p.scanner.IsStopped() {
				break
			}
			// get next record
			record, ok := p.scanner.Next()
			if !ok {
				// not found record continue
				continue
			}

			// go call handle process
			go func(record Record) {
				defer func() {
					if err := recover(); err != nil {
						golog.Error("Handle process panic",
							golog.Object("Error", err))
						debug.PrintStack()
					}
				}()

				swg.Add()
				defer swg.Done()
				p.handle.Process(record)
			}(record)
		}
		swg.Wait()
	}()
}

// Stop processor
func (p *Processor) Stop() {
	// set running false
	p.isRunning = false

	// stop scanner
	p.scanner.Stop()
}

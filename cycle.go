package cycle

import (
	"time"
)

type Cycle struct {
	// The ticker
	ticker *time.Ticker
	// Stop cycling channel
	stop chan int
	// When cycling ends
	endTime time.Time
	// Automatically stops cycling afyter x milliseconds.
	// 0 means no autamatically stop.
	stopAfter time.Duration
	// Interval between each cycle
	interval time.Duration
	// Number of cycles concluded
	nCycles int
	// The Handler called every cycle.
	Handler func(t time.Time)
	// Handler called on stop
	Stopped chan int
	// Execute handler asyncronously
	Async bool
}

// Creates a new cyclic handler
func New(handler func(t time.Time), interval time.Duration, stopAfter time.Duration) Cycle {
	return Cycle{
		ticker:    time.NewTicker(interval),
		stop:      make(chan int, 1),
		stopAfter: stopAfter,
		interval:  interval,
		Handler:   handler,
		Stopped:   make(chan int, 1),
	}
}

// Stop cycling
func (c *Cycle) Stop() {
	c.ticker.Stop()
	c.stop <- 1
	c.Stopped <- 1
}

// Check if cycle's life has expired, if true stop cycling
func (c *Cycle) CheckLifeExpired() {
	if c.stopAfter != 0 && time.Since(c.endTime) > 0 {
		c.Stop()
	}
}

// Set lifetime
func (c *Cycle) SetLifetime() {
	c.endTime = time.Now().Add(c.stopAfter)
}

// Start cyclic handler
func (c *Cycle) Run() {
	go func() {
		c.SetLifetime()
		for {
			c.CheckLifeExpired()
			select {
			case <-c.stop:
				return
			// every cycle
			case t := <-c.ticker.C:
				c.Handler(t)

				// Incr number of cycles
				c.nCycles++
			}
		}
	}()
}

// Resets interval between each cycle
func (c *Cycle) ResetInterval(t time.Duration) {
	c.ticker.Reset(t)
}

// Resets the time to stop cycling
func (c *Cycle) ResetStopTime(t time.Duration) {
	c.stopAfter = t
	c.SetLifetime()
}

// Number of cycles executed
func (c *Cycle) Count() int {
	return c.nCycles
}

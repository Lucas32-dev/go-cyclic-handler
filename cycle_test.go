package cycle

import (
	"testing"
	"time"
)

func TestCycleCount(t *testing.T) {
	e := 0
	c := New(func(tt time.Time) { e++ }, time.Second, time.Second*3)
	c.Run()
	<-c.Stopped
	if c.Count() != 3 || e != 3 {
		t.Errorf(`number of cycles expected: 3, received: %d, 
		executions expected: 3, executions made: %d`, c.Count(), e)
	}

}

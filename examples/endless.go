package examples

import (
	"log"
	"time"

	cycle "github.com/Lucas32-dev/go-cyclic-handler"
)

func Endless() {
	// will not stop cycling util c.Stop() is called
	c := cycle.New(func(t time.Time) {
		log.Println(t)
	}, time.Second, 0)
	c.Run()
}

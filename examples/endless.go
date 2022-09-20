package examples

import (
	"log"
	"time"

	cycle "github.com/Lucas32-dev/go-cyclic-handler"
)

func Endless() {
	handler := func(t time.Time) {
		log.Println(t)
	}

	// will not stop cycling util c.Stop() is called
	c := cycle.New(
		handler,
		time.Second, // interval between each cycle
		0,           // don't stop cycling
	)
	c.Run()
}

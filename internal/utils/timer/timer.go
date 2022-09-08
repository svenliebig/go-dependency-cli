package timer

import (
	"fmt"
	"time"
)

var timers = map[string]time.Time{}
var ellapsed = map[string]time.Duration{}

type Timer interface {
	Start(name string)
	Stop(name string)
	Print(name string)
}

func Start(name string) {
	timers[name] = time.Now()
}

func Stop(name string) {
	before := timers[name]
	now := time.Now()
	// fmt.Printf("Stopping time for [%s], before it was %d, duration was %d.\n", name, before.UnixNano(), now.UnixNano())
	ellapsed[name] = time.Duration(now.UnixNano() - before.UnixNano())
}

func Print(name string) {
	fmt.Printf("Timer [%s] took %dms\n", name, ellapsed[name])
}

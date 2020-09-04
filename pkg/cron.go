package pkg

import (
	"time"
)

func runEveryTime(period int, f func(time2 time.Time)) {
	ticker := time.NewTicker(time.Duration(period) * time.Millisecond)

	for {
		select {
		case t := <-ticker.C:
			f(t)
		}
	}
}
func Cron() {
	go runEveryTime(1000, func(t time.Time) {
		ReplicationCron()
	})
}

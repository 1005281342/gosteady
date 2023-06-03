package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	log.Println(late("22 9 * * *", time.Hour))
	log.Println(late("22 9 * * *", 12*time.Hour))
	log.Println(late("22 9 * * *", 24*time.Hour))
}

func late(spec string, d time.Duration) bool {
	var s, err = cron.ParseStandard(spec)
	if err != nil {
		return true
	}

	var now = time.Now()
	var current = s.Next(now.Add(-d))
	log.Printf("now:%s, can:%s", now, current)
	if current.Before(now) {
		return true
	}
	return false
}

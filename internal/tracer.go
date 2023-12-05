package internal

import (
	"log"
	"time"
)

func Trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func Un(s string, start time.Time) {
	end := time.Now()
	log.Println("  END:", s, "ElapsedTime:", end.Sub(start))
}

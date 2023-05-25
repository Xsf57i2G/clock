package main

import (
	"fmt"
	"time"
)

// clock prints the current
// time in the given format.
func clock(format string) {
	t := time.Now()
	for {
		t = t.Add(time.Second)
		fmt.Printf("\r%v", t.Format(format))

		time.Sleep(time.Until(t))
	}
}

func main() {
	date, time := "Jan 2", "15:04:05"
	clock(date + " " + time)
}
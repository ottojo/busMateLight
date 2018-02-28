package main

import (
	"github.com/ottojo/ulmAbfahrtenMonitor/swu"
	"time"
	"fmt"
	"strconv"
	"github.com/ottojo/lights"
)

var s *swu.Session

func main() {
	s = swu.NewSession("1255")
	refresh()
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			go refresh()
		}
	}()
	time.Sleep(2 * time.Hour)
	ticker.Stop()
	lights.SetMateLightAll(lights.Color{0, 0, 0})
	fmt.Println("Ticker stopped")
}

func refresh() {

	d := s.GetDepartures()
	var countdowns []int

	for _, d := range d {
		if d.ItdServingLine.Direction == "Wissenschaftsstadt" {
			c, _ := strconv.Atoi(d.Countdown)
			if c > 9 {
				c = 9
			}
			countdowns = append(countdowns, c)

			if len(countdowns) >= 4 {
				break
			}
		}
	}
	displayBarGraph(countdowns, lights.Color{1, 1, 1}, lights.Color{0, 0, 0})
}

func displayBarGraph(values []int, color, background lights.Color) {
	lights.SetMateLightAll(background)
	for x, c := range values {
		for i := 8; i > 8-c; i-- {
			time.Sleep(10 * time.Millisecond)
			lights.SetMateLightPixel(x, i, color)
		}
	}
}

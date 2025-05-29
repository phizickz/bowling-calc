package main

import (
	"fmt"
	"math/rand"
)

type frame struct {
	seq      int
	throwOne int
	throwTwo int
}

type frames struct {
	frame []*frame
}

var (
	frameCounter  int
	t1            int
	t2            int
	remainingPins int
)

func increment() {
	frameCounter++
}

func makeFrame(one, two int) *frame {
	frame := &frame{
		seq:      frameCounter,
		throwOne: one,
		throwTwo: two,
	}
	increment()
	return frame
}

func throw(seed int) int {
	return rand.Intn(seed + 1)
}

func main() {
	frameCounter = 1
	frames := frames{
		frame: make([]*frame, 11),
	}
	for i := range frames.frame {
		t1 = 0
		t2 = 0
		remainingPins = 10
		t1 = throw(remainingPins)
		remainingPins -= t1
		if t1 < 10 {
			t2 = throw(remainingPins)
		}
		frames.frame[i] = makeFrame(t1, t2)
		fmt.Println(frames.frame[i])
	}
}

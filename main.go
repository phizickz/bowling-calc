package main

import (
	"fmt"
	"math/rand"
)

type frame struct {
	seq      int
	throwOne int
	throwTwo int
	score    int
}

type frames struct {
	frame []*frame
}

func makeFrame(seq, one, two int) *frame {
	frame := &frame{
		seq:      seq,
		throwOne: one,
		throwTwo: two,
	}
	return frame
}

func throw(seed int) int {
	return rand.Intn(seed + 1)
}

func isSpare(frame *frame) bool {
	if frame.throwOne+frame.throwTwo == 10 {
		return true
	}
	return false
}

func isStrike(frame *frame) bool {
	if frame.throwOne == 10 {
		return true
	}
	return false
}

func calculateFrameScore(frames *frames, i int) {
	if isStrike {

	}
	else if isSpare {

	}
	else {
	score := frames.frame[i].throwOne + frames.frame[i].throwTwo
	}

	frames.frame[i].score = score
}

func main() {
	frames := &frames{
		frame: make([]*frame, 11),
	}

	totalPins := 10
	for i := 0; i <= len(frames.frame); i++ {
		currFrame := i + 1
		t1 := 0
		t2 := 0
		t1 = throw(totalPins)
		if currFrame == 10 && t1 == 10 {
			frames.frame[i] = makeFrame(currFrame, t1, t2)
			calculateFrameScore(frames, i)
			continue
		}

		if t1 < 10 {
			t2 = throw(totalPins - t1)
		}
		frames.frame[i] = makeFrame(currFrame, t1, t2)
		calculateFrameScore(frames, i)
		if currFrame >= 10 {
			break
		}
	}
	for i := range frames.frame {
		fmt.Println(frames.frame[i])
	}
}

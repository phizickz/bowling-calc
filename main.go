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

type bonus struct {
	originFrame    *frame
	eligibleThrows int
}

func calculateFrameScore(frames *frames) {
	accumulatedBonuses := []bonus{}
	for _, frame := range frames.frame {
		if frame == nil {
			continue
		}
		score := 0
		if len(accumulatedBonuses) > 0 {
			for _, b := range accumulatedBonuses {
				if b.eligibleThrows > 0 {
					b.originFrame.score += frame.throwOne
					b.eligibleThrows -= 1
				}
				if b.eligibleThrows > 0 {
					b.originFrame.score += frame.throwTwo
					b.eligibleThrows -= 1
				}
			}
		}
		if isStrike(frame) {
			score += 10
			accumulatedBonuses = append(accumulatedBonuses, bonus{frame, 2})
		} else if isSpare(frame) {
			score += 10
			accumulatedBonuses = append(accumulatedBonuses, bonus{frame, 1})
		} else {
			score += frame.throwOne + frame.throwTwo
		}

		frame.score = score
	}
}

func frameToString(frame *frame) {
	fmt.Printf("Frame: %v - ", frame.seq)
	fmt.Printf("Throws: %v, %v. ", frame.throwOne, frame.throwTwo)
	fmt.Printf("Score: %v.", frame.score)
	fmt.Printf("\n")
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
			continue
		}

		if t1 < 10 {
			t2 = throw(totalPins - t1)
		}
		frames.frame[i] = makeFrame(currFrame, t1, t2)
		if currFrame >= 10 {
			break
		}
	}
	var total int
	calculateFrameScore(frames)
	for _, f := range frames.frame {
		if f == nil {
			continue
		}
		frameToString(f)
		total += f.score
	}
	fmt.Printf("Total score: %v", total)
}

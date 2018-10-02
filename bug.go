// ticker / clear cgo bug test

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var BlinkMSec = 200

const NBlinkers = 20

var Blinkers [NBlinkers]*time.Ticker

// BlinkFunc
func BlinkFunc(blinkn int) {
	for {
		if Blinkers[blinkn] == nil {
			return // shutdown..
		}
		<-Blinkers[blinkn].C
		// fmt.Printf("blink %v\n", blinkn)
	}
}

func main() {
	var wait sync.WaitGroup
	for i := 0; i < NBlinkers; i++ {
		Blinkers[i] = time.NewTicker(time.Duration(rand.Intn(BlinkMSec)) * time.Millisecond)
		go BlinkFunc(i)
		wait.Add(1)
	}

	go runClear()
	go runClear()
	go runClear()
	go runClear()
	go runClear()

	wait.Wait()
}

var mu sync.Mutex

func runClear() {
	for j := 0; j < 10000; j++ {
		txt := fmt.Sprintf("some random text %v and some more %v and some more %v\n", j, j, j)
		mu.Lock()
		WriteText([]byte(txt))
		mu.Unlock()
		// fmt.Print(txt)
		time.Sleep(time.Duration(rand.Intn(BlinkMSec)) * time.Millisecond)
	}
}

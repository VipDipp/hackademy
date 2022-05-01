package letter

import (
	"sync"
)

type FreqMap map[rune]int

var wg = sync.WaitGroup{}
var m = make(map[rune]int)
var mtx = &sync.Mutex{}

func Frequency(str string) FreqMap {
	var m = make(FreqMap)
	for _, v := range str {
		m[v]++
	}
	return m
}

func concurrentFrequencyPartial(str string) {
	for _, v := range str {
		mtx.Lock()
		m[v]++
		mtx.Unlock()
	}
	wg.Done()
}

func ConcurrentFrequency(strs []string) FreqMap {
	m = make(FreqMap)
	wg.Add(len(strs))
	for _, v := range strs {
		go concurrentFrequencyPartial(v)
	}
	wg.Wait()
	return m
}

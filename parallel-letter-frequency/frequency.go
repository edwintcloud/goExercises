package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency gets the frequence of each sentence in parallel
func ConcurrentFrequency(sentences []string) FreqMap {
	m := FreqMap{}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	wg.Add(len(sentences))

	for _, sentence := range sentences {
		go func(s string) {
			defer wg.Done()
			for _, r := range s {
				mutex.Lock()
				m[r]++
				mutex.Unlock()
			}
		}(sentence)
	}
	wg.Wait()
	return m
}

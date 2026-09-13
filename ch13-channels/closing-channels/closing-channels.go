package main

func countReports(numSentCh chan int) int {
	count := 0
	for {
		chVal, ok := <-numSentCh
		if !ok {
			break
		}
		count += chVal
	}
	return count
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

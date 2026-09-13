package main

func countReports(numSentCh chan int) int {
	count := 0
	for {
		if v, ok := <-numSentCh; ok { // ok! {
			count += v
			// break
			continue
		}
		// count += <- numSentCh
		break
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

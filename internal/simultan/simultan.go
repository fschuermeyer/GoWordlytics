package simultan

import "sync"

func Run(dos []func()) {
	var wg sync.WaitGroup

	for _, do := range dos {
		RunUnit(&wg, do)
	}

	wg.Wait()
}

func RunUnit(wg *sync.WaitGroup, do func()) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		do()
	}()
}

package simultan_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/fschuermeyer/GoWordlytics/internal/simultan"
)

func TestRun(t *testing.T) {
	var counter int32
	dos := make([]func(), 10)

	for i := range dos {
		dos[i] = func() {
			atomic.AddInt32(&counter, 1)
		}
	}

	simultan.Run(dos)

	if counter != int32(len(dos)) {
		t.Errorf("Expected counter to be %d, but got %d", len(dos), counter)
	}
}

func TestRunUnit(t *testing.T) {
	var counter int32
	var wg sync.WaitGroup

	do := func() {
		atomic.AddInt32(&counter, 1)
		time.Sleep(100 * time.Millisecond)
	}

	simultan.RunUnit(&wg, do)
	wg.Wait()

	if counter != 1 {
		t.Errorf("Expected counter to be 1, but got %d", counter)
	}
}

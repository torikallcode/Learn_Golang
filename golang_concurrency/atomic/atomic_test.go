package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// func TestAtomic(t *testing.T) {
// 	var counter int32
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < 1000; j++ {
// 				atomic.AddInt32(&counter, 1) //menambahkan 1 ke variable counter secara atomic agar aman dari race condition
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final counter:", counter)
// }

func TestAtomic(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}

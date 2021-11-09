// package main

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func main() {

// 	rand.Seed(time.Now().UnixNano())

// 	const N = 5
// 	var values [N]int32

// 	var wg sync.WaitGroup
// 	var o sync.Once
// 	wg.Add(N)

// 	a := func() {
// 		fmt.Println("hello")
// 	}
// 	for i := 0; i < N; i++ {
// 		// fmt.Println(i)
// 		i := i
// 		o.Do(a)
// 		// fmt.Println(i
// 		go func() {
// 			values[i] = 50 + rand.Int31n(10)
// 			log.Println("Done:", i)
// 			wg.Done() // <=> wg.Add(-1)
// 		}()
// 	}
// 	// a := func() {
// 	// 	fmt.Println("hello")
// 	// }
// 	wg.Wait()
// 	// All elements are guaranteed to be
// 	// initialized now.
// 	log.Println("values:", values)
// }
package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	m sync.Mutex
	n uint64
}

func (c *Counter) Value() uint64 {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}

func (c *Counter) Increase(delta uint64) {
	c.m.Lock()
	c.n += delta
	c.m.Unlock()
}

func main() {
	var c Counter
	for i := 0; i < 100; i++ {
		go func() {
			for k := 0; k < 100; k++ {
				c.Increase(1)
			}
		}()
	}

	// The loop is just for demo purpose.
	for c.Value() < 10000 {
		runtime.Gosched()
	}
	fmt.Println(c.Value()) // 10000
}

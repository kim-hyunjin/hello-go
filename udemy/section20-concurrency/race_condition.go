package section20concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func raceConditionExam() {
	counter := 0

	const gs = 1000
	var wg sync.WaitGroup
	wg.Add(gs)
	
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go func() {
			v := counter // reading shared variable
			runtime.Gosched()
			v++
			counter = v // write to shared variable
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("all go routines done!")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}
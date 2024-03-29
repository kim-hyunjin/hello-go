package section20concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func exam() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("GOARCH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar: ", i)
	}
}

package routines

import (
	"fmt"
	"sync"
	"time"
)

/**
GO Routine Definition
*********************************************************************************
 A goroutine is a lightweight thread managed by the Go runtime.
 Goroutines run in the same address space, so access to shared memory must be synchronized.
 The sync package provides useful primitives, although you won't need them much in Go as there are other primitives


 Unlike java which uses OS thread as threads which is memory heavy, slow and hardware dependent
 GO routines are green threads. An abstraction of actual thread and lightweight. we can use millions of goroutines
**/

/*
*
Wait grop is go routine synchronization.
so that the thread will wait until go routine completion
*
*/
var wg = sync.WaitGroup{}

func sampleRoutine() {
	time.Sleep(10 * time.Second)
	fmt.Println("Received Data from the satellite")

	//Mark the wait group as done
	wg.Done()
}

func UseRoutine() {

	fmt.Println("Communication initialized with Satellite")
	// Here 1 routine is there so 1 group is added
	wg.Add(1)

	/**
	Go routine is invoked here
	This function will be pushed to separate thread and won't block this thread execution
	**/
	go sampleRoutine()

	fmt.Println("Monitoring other parameters")

	// Wait till the go wait group is done
	wg.Wait()
}

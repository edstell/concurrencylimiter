package concurrencylimiter

import "sync"

// A ConcurrencyLimiter executes submitted tasks in parallel until the
// concurrency limit is reached, after which submission of new tasks blocks
// until the number 'in-flight' falls below the limit.
type ConcurrencyLimiter struct {
	wg    sync.WaitGroup
	doers chan func(func())
}

// New which will allow 'limit' concurrent tasks.
func New(limit int) *ConcurrencyLimiter {
	cl := &ConcurrencyLimiter{
		doers: make(chan func(func()), limit),
	}
	for i := 0; i < limit; i++ {
		cl.doers <- cl.do
	}
	return cl
}

// Do the task, blocking until there are fewer than the limit in-flight.
func (cl *ConcurrencyLimiter) Do(task func()) {
	do := <-cl.doers
	cl.wg.Add(1)
	go do(task)
}

// Wait for all in-flight tasks to finish.
func (cl *ConcurrencyLimiter) Wait() {
	cl.wg.Wait()
}

// worker executes the task, then adds itself back to the pool of available
// workers.
func (cl *ConcurrencyLimiter) do(task func()) {
	defer cl.wg.Done()
	task()
	cl.doers <- cl.do
}

package concurrencylimiter

import "sync"

// A Limiter executes submitted tasks in parallel until the
// concurrency limit is reached, after which submission of new tasks
// blocks until the number 'in-flight' falls below the limit.
type Limiter struct {
	wg      sync.WaitGroup
	workers chan func(func())
}

// NewLimiter which will allow 'limit' concurrent tasks.
func NewLimiter(limit int) *Limiter {
	l := &Limiter{
		workers: make(chan func(func()), limit),
	}
	for i := 0; i < limit; i++ {
		l.workers <- l.worker
	}
	return l
}

// Submit the task for execution, blocking until there are fewer than the limit
// tasks in-flight.
func (l *Limiter) Submit(task func()) {
	worker := <-l.workers
	l.wg.Add(1)
	go worker(task)
}

// Wait for all in-flight tasks to finish.
func (l *Limiter) Wait() {
	l.wg.Wait()
}

// worker executes the task, then adds itself back to the pool of available
// workers.
func (l *Limiter) worker(task func()) {
	defer l.wg.Done()
	task()
	l.workers <- l.worker
}

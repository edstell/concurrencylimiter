# concurrencylimiter
A simple package for executing work concurrently - up to a limit.

The intended usecase looks something like:
```
func concurrentlyDo(tasks []func(), withLimit int) {
	cl := concurrencylimiter.New(withLimit)
	for _, task := range tasks {
		cl.Do(task)
	}
	cl.Wait()
}
```

Where tasks will be executed in parallel until the limit is reached, after which
submitting new tasks will block until there are fewer than the limit in-flight.

# concurrencylimiter
A simple package for executing work in concurrently - up to a limit.

The intended usecase looks something like:
```
func do(tasks []func()) {
	limiter := concurrencylimiter.NewLimiter(5)
	for _, task := range tasks {
		limiter.Submit(task)
	}
	limiter.Wait()
}
```

Where tasks will be executed in parallel until the limit is reached, after which
submitting new tasks will block until there are fewer than the limit in-flight.

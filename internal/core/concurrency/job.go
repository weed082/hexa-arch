package concurrency

type Job struct {
	Callback func()
}

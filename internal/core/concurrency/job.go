package concurrency

type Job struct {
	execute func(params interface{}) error
}

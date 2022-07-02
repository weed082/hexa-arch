package port

type WorkerPool interface {
	Generate(count int)
	Stop()
	RegisterJob(callback func())
}

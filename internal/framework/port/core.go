package port

type WorkerPool interface {
	RegisterJob(callback func())
}

package port

type WorkerPool interface {
	Generate(count int)
}

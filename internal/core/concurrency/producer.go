package concurrency

// Producer simulates an external library that invokes the registered
// callback when it has new data for us once per 100ms
type Producer struct {
}

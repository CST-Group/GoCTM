package memory

type DistributedMemory struct {
	Memory

	ID       int64
	Name     string
	Memories []Memory
}

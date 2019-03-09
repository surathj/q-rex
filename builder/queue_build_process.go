package builder

type QueueBuildProcess interface {
	SetSize(int) QueueBuildProcess
	SetType(QueueType) QueueBuildProcess
}
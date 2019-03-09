package builder

type BuilderDirector struct{
	builder QueueBuildProcess
}

func (md *BuilderDirector) SetQueueBuilder(q QueueBuildProcess){
	md.builder = q
}

func (md *BuilderDirector) Construct(){
	md.builder.SetSize(10).SetType(Simple)
}
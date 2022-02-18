package utils

func GetProcess() *Process {
	t := &thread{registers: registers{Pc: pc{Priority: 1.5}}}
	return &Process{
		Private: t,
		thread:  t,
	}
}

type Process struct {
	Private interface{}
	thread  *thread
}

type thread struct {
	registers
}

type registers struct {
	Pc pc
}

type pc struct {
	Priority float32
}

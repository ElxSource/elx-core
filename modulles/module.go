package modules



type varEvaluable string

type Task struct {
	Name string
	body Executable
}

type Executable interface {
	Execute()
}

type Debug struct {
	msg varEvaluable
}

func (task *Debug) Execute() {
	log.Println(task.msg)
}


type Echo struct{
  msg varEvaluable
}

func (task *Echo) Execute() {
	log.Println(task.msg)
}

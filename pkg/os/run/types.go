package run

type Args struct {
	Cmd string `name:"cmd"`
}

type Proc struct {
	Pid      int64  `name:"pid"`
	Result   string `name:"result"`
	ExitCode int64  `name:"exit_code"`
}

type Result struct {
	Error string `name:"error"`
	Proc  Proc   `name:"proc"`
}

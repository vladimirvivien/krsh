package run

import (
	"fmt"

	"github.com/vladimirvivien/gexe"
	"github.com/vladimirvivien/krsh/pkg/typekit"
	"go.starlark.net/starlark"
)

var (
	Namespace = "os"
	Name      = "run"
	Func      = run
	Builtin   = starlark.NewBuiltin(string(Name), Func)
)

// Register command
func init() {
	typekit.Register(Namespace, Name, Builtin)
}

// run executes the provided OS CLI command on the local machine.
// Starlark format: result = os.run(cmd="script-command")
func run(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var args Args
	if err := typekit.KwargsToGo(kwargs, &args); err != nil {
		return typekit.BuiltinError(Name, fmt.Errorf("%s: %s", Name, err))
	}

	proc := gexe.New().RunProc(args.Cmd)
	if proc.Err() != nil {
		return typekit.ScriptError(Name, proc.Err())
	}
	return typekit.BuiltinResult(
		Name,
		Result{
			Proc: Proc{
				Pid:      int64(proc.ID()),
				Result:   proc.Result(),
				ExitCode: int64(proc.ExitCode()),
			},
		},
	)
}

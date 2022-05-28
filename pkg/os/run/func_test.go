package run

import (
	"strings"
	"testing"

	"github.com/vladimirvivien/krsh/pkg/typekit"
	"go.starlark.net/starlark"
)

func TestRunFunc(t *testing.T) {
	tests := []struct {
		name       string
		kwargs     []starlark.Tuple
		expected   string
		shouldFail bool
	}{
		{
			name: "simple exec",
			kwargs: []starlark.Tuple{
				{starlark.String("cmd"), starlark.String("echo 'Hello World!'")},
			},
			expected: "Hello World!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Func(&starlark.Thread{}, nil, nil, test.kwargs)
			if err != nil {
				t.Fatal(err)
			}
			var result Result
			if err := typekit.Starlark(res).Go(&result); err != nil {
				t.Fatal(err)
			}
			if !test.shouldFail && result.Error != "" {
				t.Fatal(result.Error)
			}

			if result.Proc.Result != test.expected {
				t.Errorf("command returned unexpected result: %s", result.Proc.Result)
			}
			if result.Proc.Pid == 0 {
				t.Errorf("successful command returned 0 pid")
			}
			if result.Proc.ExitCode != 0 {
				t.Errorf("successful command returned non-zero exit code")
			}
		})
	}
}

func TestRunFuncScript(t *testing.T) {
	tests := []struct {
		name     string
		script   string
		expected string
	}{
		{
			name:     "simple script",
			script:   `result=os.run(cmd="""echo 'Salut Monde!'""")`,
			expected: "Salut Monde!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := starlark.ExecFile(&starlark.Thread{}, "test.star", strings.NewReader(test.script), typekit.Registry())
			if err != nil {
				t.Fatal(err)
			}

			resultVal := output["result"]
			if resultVal == nil {
				t.Fatal("run() should be assigned to a variable for test")
			}
			var result Result
			if err := typekit.Starlark(resultVal).Go(&result); err != nil {
				t.Fatal(err)
			}

			if result.Proc.Result != test.expected {
				t.Errorf("command returned unexpected result: %s", result.Proc.Result)
			}
			if result.Proc.Pid == 0 {
				t.Errorf("successful command returned 0 pid")
			}
			if result.Proc.ExitCode != 0 {
				t.Errorf("successful command returned non-zero exit code")
			}

		})
	}
}

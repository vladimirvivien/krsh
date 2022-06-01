package run

import (
	"context"
	"testing"

	"github.com/vladimirvivien/krsh/pkg/typekit"
	"go.starlark.net/starlark"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

func TestRun(t *testing.T) {
	testTable := features.Table{
		{
			Name: "simple run",
			Assessment: func(ctx context.Context, t *testing.T, config *envconf.Config) context.Context {
				kwargs := []starlark.Tuple{
					{starlark.String("name"), starlark.String("hello-world")},
					{starlark.String("image"), starlark.String("busybox")},
				}
				val, err := run(&starlark.Thread{}, nil, nil, kwargs)
				if err != nil {
					t.Fatal(err)
				}

				var result Result
				if err := typekit.Starlark(val).Go(&result); err != nil {
					t.Fatal(err)
				}

				if result.Error != "" {
					t.Fatal(result.Error)
				}

				return ctx
			},
		},
	}.Build("run deployment")

	testEnv.Test(t, testTable.Feature())
}

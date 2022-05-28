package run

import (
	"context"
	"fmt"

	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/container"
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/deployment"
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/meta"
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/pod"
	"github.com/vladimirvivien/krsh/pkg/typekit"
	"go.starlark.net/starlark"
	appsV1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/e2e-framework/klient/conf"
	"sigs.k8s.io/e2e-framework/klient/k8s/resources"
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

// run deploys a single pod using the specified image.
func run(_ *starlark.Thread, b *starlark.Builtin, _ starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var args Args
	if err := typekit.KwargsToGo(kwargs, &args); err != nil {
		return typekit.BuiltinError(Name, fmt.Errorf("%s: %s", Name, err))
	}

	cfg, err := conf.New(conf.ResolveKubeConfigFile())
	if err != nil {
		return typekit.ScriptError(Name, err)
	}

	res, err := resources.New(cfg)
	if err != nil {
		return typekit.ScriptError(Name, err)
	}

	dep := makeDeployment(args)
	if err := res.Create(context.TODO(), &dep); err != nil {
		return typekit.ScriptError(Name, err)
	}

	return typekit.BuiltinResult(Name, Result{Deployment: dep})
}

func makeDeployment(args Args) appsV1.Deployment {
	return deployment.Deployment(
		meta.Object(args.Name).Namespace(args.Namespace),
		deployment.Replicas(1),
		meta.MatchLabels(args.Labels),
		deployment.StrategyDefault,
		pod.Template(meta.ObjectMetaNone, pod.Spec(container.Name(args.Name).Image(args.Image))),
	).Build()
}

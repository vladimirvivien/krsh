package run

import (
	appsV1 "k8s.io/api/apps/v1"
)

type Args struct {
	Annotations   map[string]string `name:"annotations" optional:"true"`
	Image         string            `name:"image"`
	Labels        map[string]string `name:"labels" optional:"true"`
	Name          string            `name:"name"`
	Namespace     string            `name:"namespace" optional:"true"`
	Privileged    bool              `name:"privileged" optional:"true"`
	PullPolicy    string            `name:"pull_policy" optional:"true"`
	RestartPolicy string            `name:"restart_policy" optional:"true"`
}

type Result struct {
	Error      string `name:"error"`
	Deployment appsV1.Deployment
}

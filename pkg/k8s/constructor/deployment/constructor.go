// Package deployment provides constructor types for appV1.Deployment
package deployment

import (
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/meta"
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/pod"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
)

type Constructor struct {
	deployment appsV1.Deployment
}

// Deployment is the initializer function for type deployment.Constructor
func Deployment(
	deploymentMeta meta.ObjectMetaConstructor,
	replicas *int32,
	selector meta.LabelSelectorConstructor,
	strategy StrategyConstructor,
	template pod.TemplateSpecConstructor,
) Constructor {
	tempSpec := template.Build()
	depSel := selector.Build()
	return Constructor{
		deployment: appsV1.Deployment{
			ObjectMeta: deploymentMeta.Build(),
			Spec: appsV1.DeploymentSpec{
				Replicas: replicas,
				Selector: &depSel,
				Template: v1.PodTemplateSpec{
					ObjectMeta: tempSpec.ObjectMeta,
					Spec:       tempSpec.Spec,
				},
				Strategy: strategy.Build(),
			},
		},
	}
}

// Build is the finalizer method that constructs and return a value of type appsV1.Deployment
func (c Constructor) Build() appsV1.Deployment {
	return c.deployment
}

// Replicas is an initializer func that converts int -> *int32
func Replicas(r int) *int32 {
	rep := int32(r)
	return &rep
}

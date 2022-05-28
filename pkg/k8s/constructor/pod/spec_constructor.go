package pod

import (
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/container"
	coreV1 "k8s.io/api/core/v1"
)

type SpecConstructor struct {
	spec coreV1.PodSpec
}

// Spec initializer method for type PodSpecConstructor
func Spec(containerConstructors ...container.Constructor) SpecConstructor {
	spec := SpecConstructor{spec: coreV1.PodSpec{}}
	for _, constructor := range containerConstructors {
		spec.spec.Containers = append(spec.spec.Containers, constructor.Build())
	}
	return spec
}

// AddVolume is a setter method to store coreV1.Volume definition
func (c SpecConstructor) AddVolume(vol coreV1.Volume) SpecConstructor {
	c.spec.Volumes = append(c.spec.Volumes, vol)
	return c
}

// Build is the finalizer method that returns a value of type coreV1.PodSpec
func (c SpecConstructor) Build() coreV1.PodSpec {
	return c.spec
}

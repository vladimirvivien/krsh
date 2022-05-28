// Package pod package contains constructor types to build *coreV1.Pod values
package pod

import (
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/meta"
	coreV1 "k8s.io/api/core/v1"
)

type Constructor struct {
	pod coreV1.Pod
}

// Pod initializer function for type pod.Constructor
func Pod(metaConstructor meta.ObjectMetaConstructor, specConstructor SpecConstructor) Constructor {
	return Constructor{pod: coreV1.Pod{ObjectMeta: metaConstructor.Build(), Spec: specConstructor.Build()}}
}

func (c Constructor) Build() coreV1.Pod {
	return c.pod
}

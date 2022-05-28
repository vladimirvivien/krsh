package pod

import (
	"github.com/vladimirvivien/krsh/pkg/k8s/constructor/meta"
	coreV1 "k8s.io/api/core/v1"
)

type TemplateSpecConstructor struct {
	spec coreV1.PodTemplateSpec
}

// Template initializer method for type TemplateSpecConstructor
func Template(metaConstructor meta.ObjectMetaConstructor, specConstructor SpecConstructor) TemplateSpecConstructor {
	return TemplateSpecConstructor{
		spec: coreV1.PodTemplateSpec{
			ObjectMeta: metaConstructor.Build(),
			Spec:       specConstructor.Build(),
		},
	}
}

func (c TemplateSpecConstructor) Build() coreV1.PodTemplateSpec {
	return c.spec
}

// Package meta contains constructor type to build values of type *coreV1.ObjectMeta
package meta

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	EmptyName        = ""
	EmptyLabels      = map[string]string{}
	DefaultNamespace = "default"
	ObjectMetaNone   = ObjectMetaConstructor{obj: metaV1.ObjectMeta{}}
)

// ObjectMetaConstructor constructor for type coreV1.ObjectMeta
type ObjectMetaConstructor struct {
	obj metaV1.ObjectMeta
}

// Object is the initializer function for ObjectMetaConstructor
func Object(name string) ObjectMetaConstructor {
	return ObjectMetaConstructor{obj: metaV1.ObjectMeta{Name: name}}
}

// Namespace setter for namespace value
func (c ObjectMetaConstructor) Namespace(ns string) ObjectMetaConstructor {
	c.obj.Namespace = ns
	return c
}

// Labels setter for labels
func (c ObjectMetaConstructor) Labels(labels map[string]string) ObjectMetaConstructor {
	c.obj.Labels = labels
	return c
}

// Annotations setter for annotations
func (c ObjectMetaConstructor) Annotations(labels map[string]string) ObjectMetaConstructor {
	c.obj.Annotations = labels
	return c
}

// Build is the finalizer that builds and returns metaV1.ObjectMeta
func (c ObjectMetaConstructor) Build() metaV1.ObjectMeta {
	if c.obj.Namespace == "" {
		c.obj.Namespace = DefaultNamespace
	}
	return c.obj
}
